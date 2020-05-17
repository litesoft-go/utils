package enums

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Enum struct {
	name     string
	typeName string
}

func New(name string) Enum {
	return Enum{name: name}
}

func (in *Enum) GetName() string {
	if in == nil {
		return ""
	}
	return in.name
}

func (in *Enum) String() string {
	return in.getTypeName() + ": " + in.GetName()
}

func (in *Enum) MarshalJSON() ([]byte, error) {
	name := in.GetName()
	return json.Marshal(&name)
}

func (in *Enum) setName(name string) {
	in.name = name
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

type IEnum interface {
	GetName() string
	String() string

	UpdateFrom(found IEnum)

	setName(string)
	getTypeName() string
	setTypeName(string)
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

func Populate(emptyEntry IEnum, name string) (found bool, err error) {
	typeName := reflect.TypeOf(emptyEntry).String()
	emptyEntry.setTypeName(typeName)
	emptyEntry.setName(name)
	errMsg := "name empty"
	if name != "" {
		set := members[typeName]
		errMsg = "type not registered"
		if set != nil {
			entry, exists := set.entryByName[name]
			if exists {
				emptyEntry.UpdateFrom(entry)
				return true, nil
			}
		}
	}
	err = fmt.Errorf("can't Populate (%s: %s) - %s", typeName, name, errMsg)
	return
}

func Add(enumEntries ...IEnum) {
	for _, enumEntry := range enumEntries {
		typeName := reflect.TypeOf(enumEntry).String()
		enumEntry.setTypeName(typeName)
		name := enumEntry.GetName()
		if name == "" {
			panic("Can't Add: " + enumEntry.String())
		}
		set := members[typeName]
		if set == nil {
			set = newSet(enumEntry)
			members[typeName] = set
		}
		set.add(enumEntry)
	}
}

var members = map[string]*enumSet{}

type enumSet struct {
	typeName    string
	entryByName map[string]IEnum
	//entries     []IEnum
}

func newSet(entry IEnum) *enumSet {
	typeName := entry.getTypeName()
	return &enumSet{typeName: typeName,
		entryByName: map[string]IEnum{}}
}

func (in *enumSet) add(entry IEnum) {
	_, exists := in.entryByName[entry.GetName()]
	if exists {
		panic("Duplicate entry: " + entry.String())
	}
	in.entryByName[entry.GetName()] = entry
}
