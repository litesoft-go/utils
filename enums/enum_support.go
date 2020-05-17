package enums

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ComparisonTransformer transforms keys, Enum Name(s) and Aliases, to be able to match on
// strict or transformed (e.g. LowerCase, UpperCase, or Folded)
type ComparisonTransformer func(string) string

type Enum struct {
	name              string
	typeName          string
	unmarshalledValue string
	isDefault         bool
}

func New(name string) Enum {
	return Enum{name: name}
}

func (in *Enum) Name() string {
	if in == nil {
		return ""
	}
	return in.name
}

func (in *Enum) GetUnmarshalledValue() string {
	if in == nil {
		return ""
	}
	return in.unmarshalledValue
}

func (in *Enum) IsDefault() bool {
	return in.isDefault
}

func (in *Enum) String() (str string) {
	str = in.getTypeName() + ": " + in.Name()
	value := in.GetUnmarshalledValue()
	if value != "" {
		str += "(" + value + ")"
	}
	return
}

func (in *Enum) MarshalJSON() ([]byte, error) {
	name := in.GetUnmarshalledValue()
	if name == "" {
		name = in.Name()
	}
	return json.Marshal(&name)
}

func (in *Enum) setName(name string) {
	in.name = name
}

func (in *Enum) setUnmarshalledValue(value string) {
	in.unmarshalledValue = value
}

func (in *Enum) getTypeName() (typeName string) {
	if in != nil {
		typeName = in.typeName
	}
	return
}

func (in *Enum) setTypeName(typeName string) {
	in.typeName = typeName
}

func (in *Enum) setDefault() {
	in.isDefault = true
}

func (in *Enum) isEmpty() bool {
	return (in.getTypeName() == "") && (in.Name() == "") && (in.GetUnmarshalledValue() == "")
}

type IEnum interface {
	Name() string
	GetUnmarshalledValue() string
	IsDefault() bool
	String() string

	UpdateFrom(found IEnum)

	setName(string)
	setUnmarshalledValue(string)
	getTypeName() string
	setTypeName(string)
	setDefault()
	isEmpty() bool
}

func UnmarshalJSON(in IEnum, data []byte) error {
	name := ""
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	found, err := Populate(in, name)
	if err == nil {
		if !found {
			err = fmt.Errorf("")
		}
	}
	return err
}

func Populate(emptyEntry IEnum, key string) (found bool, err error) {
	if !emptyEntry.isEmpty() {
		return false, fmt.Errorf(
			"passed Enum to Populate, not Empty: %v", emptyEntry)
	}
	typeName := reflect.TypeOf(emptyEntry).String()
	set := getSet(typeName)
	errField := "type"
	if set != nil {
		foundEntry, exists := set.getRegisteredEntryForKey(key)
		if exists {
			emptyEntry.setTypeName(typeName)
			name := foundEntry.Name()
			emptyEntry.setName(name)
			if key != name {
				emptyEntry.setUnmarshalledValue(key)
			}
			emptyEntry.UpdateFrom(foundEntry)
			return true, nil
		}
		errField = "key"
	}
	err = fmt.Errorf("can't Populate (%s: %s) - %s not registered", typeName, key, errField)
	return
}

//noinspection GoUnusedExportedFunction
func GetRegisteredMembers(enumEntryEmptyOrNot IEnum) (members []IEnum) {
	typeName := reflect.TypeOf(enumEntryEmptyOrNot).String()
	set := getSet(typeName)
	if set != nil {
		members = append(members, set.entries...)
	}
	return
}

//noinspection GoUnusedExportedFunction
func AddEmptyNameWithTransformer(enumEntry IEnum, ct ComparisonTransformer) {
	typeName := reflect.TypeOf(enumEntry).String()
	addTransformer(typeName, ct)
	addEmptyName(typeName, enumEntry)
}

//noinspection GoUnusedExportedFunction
func AddDefaultWithTransformer(enumEntry IEnum, ct ComparisonTransformer) {
	typeName := reflect.TypeOf(enumEntry).String()
	addTransformer(typeName, ct)
	addDefault(typeName, enumEntry)
}

//noinspection GoUnusedExportedFunction
func AddTransformer(enumEntry IEnum, ct ComparisonTransformer) {
	addTransformer(reflect.TypeOf(enumEntry).String(), ct)
}

//noinspection GoUnusedExportedFunction
func AddEmptyName(enumEntry IEnum) {
	addEmptyName(reflect.TypeOf(enumEntry).String(), enumEntry)
}

//noinspection GoUnusedExportedFunction
func AddDefault(enumEntry IEnum) {
	addDefault(reflect.TypeOf(enumEntry).String(), enumEntry)
}

//noinspection GoUnusedExportedFunction
func AddWithAliases(enumEntry IEnum, aliases ...string) {
	addRegular(reflect.TypeOf(enumEntry).String(), enumEntry, aliases...)
}

//noinspection GoUnusedExportedFunction
func Add(enumEntries ...IEnum) {
	for _, enumEntry := range enumEntries {
		addRegular(reflect.TypeOf(enumEntry).String(), enumEntry)
	}
}

func addTransformer(typeName string, ct ComparisonTransformer) {
	getSet(typeName).setTransformer(ct)
}

func addEmptyName(typeName string, entry IEnum) {
	name := entry.Name()
	getSet(typeName).addEntry(typeName, entry, name == "", "addEmptyName, but expected empty name was actually: "+name)
}

func addDefault(typeName string, entry IEnum) {
	name := entry.Name()
	if name != "" {
		panic("the Default entry may not have a Name, but given: " + name)
	}
	getSet(typeName).setDefault(typeName, entry)
}

func addRegular(typeName string, entry IEnum, aliases ...string) {
	getSet(typeName).addEntry(typeName, entry, entry.Name() != "", "add (regular), but name was empty", aliases...)
}

func getSet(typeName string) *enumSet {
	set := members[typeName]
	if set == nil {
		set = newSet(typeName)
		members[typeName] = set
	}
	return set
}

var members = map[string]*enumSet{}

type enumSet struct {
	typeName     string
	transformer  ComparisonTransformer
	defaultEntry IEnum
	entryByName  map[string]IEnum
	entries      []IEnum
}

func newSet(typeName string) *enumSet {
	return &enumSet{typeName: typeName,
		entryByName: map[string]IEnum{}}
}

func (in *enumSet) getRegisteredEntryForKey(key string) (foundEntry IEnum, exists bool) {
	if in.transformer != nil {
		key = in.transformer(key)
	}
	foundEntry, exists = in.entryByName[key]
	if !exists && (in.defaultEntry != nil) {
		foundEntry, exists = in.defaultEntry, true
	}
	return
}

func (in *enumSet) setTransformer(ct ComparisonTransformer) {
	if (in.transformer != nil) || (in.defaultEntry != nil) || (len(in.entryByName) != 0) {
		panic("transformers must be added first")
	}
	in.transformer = ct
}

func (in *enumSet) setDefault(typeName string, entry IEnum) {
	checkAddable(typeName, entry)
	if (in.defaultEntry != nil) || (len(in.entryByName) != 0) {
		panic("default entries must be added before any other entries")
	}
	in.defaultEntry = entry
	entry.setDefault()
}

func (in *enumSet) addEntry(typeName string, entry IEnum, okName bool, whyNot string, aliases ...string) {
	checkAddable(typeName, entry)
	if !okName {
		panic(whyNot)
	}
	in.entries = append(in.entries, entry)
	in.add2Map(entry.Name(), entry)
	for _, alias := range aliases {
		in.add2Map(alias, entry)
	}
}

func checkAddable(typeName string, entry IEnum) {
	if (entry.getTypeName() != "") || (entry.GetUnmarshalledValue() != "") {
		panic("Can't Add, appears to have already been initialized: " + entry.String())
	}
	entry.setTypeName(typeName)
}

func (in *enumSet) add2Map(key string, entry IEnum) {
	if in.transformer != nil {
		key = in.transformer(key)
	}
	_, exists := in.entryByName[key]
	if exists {
		panic("Duplicate entry: " + entry.String() + " key='" + key + "'")
	}
	in.entryByName[key] = entry
}
