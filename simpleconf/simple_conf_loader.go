package simpleconf

import (
	"bufio"
	"fmt"
	"github.com/litesoft-go/utils/interfaces"
	"github.com/litesoft-go/utils/ioutils"
	"os"
	"strings"
	"unicode"

	"io"
)

type ConfigMap struct {
	keyValues map[string][]string
}

func (in *ConfigMap) ValuesFor(key string) []string {
	return in.keyValues[key]
}

func (in *ConfigMap) ValueOf(key string) (value string, err error) {
	return in.ExtractValueOf(key, in.ValuesFor(key))
}

func (in *ConfigMap) ExtractValueOf(key string, from []string) (value string, err error) {
	switch len(from) {
	case 0:
		return
	case 1:
		value = from[0]
	default:
		err = fmt.Errorf("multiple values for '%s', use ValuesFor", key)
	}
	return
}

//noinspection GoUnusedExportedFunction
func LoadFile(filePath string) (config Config, err error) {
	var readerCloser io.ReadCloser
	readerCloser, err = os.Open(filePath)
	if err == nil {
		config, err = Load(readerCloser)
	}
	return
}

func Load(readerCloser io.ReadCloser) (config Config, err error) {
	if !interfaces.IsNil(readerCloser) {
		defer ioutils.Close(readerCloser)
		config, err = LoadReader(readerCloser)
	}
	return
}

func LoadReader(reader io.Reader) (config *ConfigMap, err error) {
	collector := Collector{keyValues: make(map[string][]string)}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		innerErr := collector.addLine(scanner.Text())
		if innerErr != nil {
			return nil, innerErr
		}
	}
	err = scanner.Err()
	if err == nil {
		err = collector.addList()
		if err == nil {
			config = &ConfigMap{keyValues: collector.keyValues}
		}
	}
	return
}

type Collector struct {
	keyValues map[string][]string
	inListKey string
	list      []string
}

func (in *Collector) addList() error {
	if in.inListKey != "" {
		if len(in.list) == 0 {
			return fmt.Errorf("list '%s' indicated, but no entries added", in.inListKey)
		}
		in.keyValues[in.inListKey] = in.list
		in.inListKey, in.list = "", nil
	}
	return nil
}

func (in *Collector) addLine(origLine string) (err error) {
	if len(origLine) == 0 {
		return
	}
	startsWithWhiteSpace := unicode.IsSpace([]rune(origLine)[0])
	line := strings.TrimSpace(origLine)
	if len(line) == 0 {
		return
	}
	firstRune := []rune(line)[0]
	//  ('#' & '!' from .properties, and ';' from JetBrains' .conf)
	if (firstRune == '#') || (firstRune == '!') || (firstRune == ';') { // comment
		return
	}
	if startsWithWhiteSpace { // Array element
		if in.inListKey == "" {
			return fmt.Errorf("no list indicated, but indented line encountered: %s", origLine)
		}
		in.list = append(in.list, line)
		return
	}
	err = in.addList()
	if err != nil {
		return
	}
	sepAt := strings.IndexAny(line, ":=")
	if sepAt == -1 {
		return fmt.Errorf("no separator ':' or '=' found in line: %s", origLine)
	}
	key := strings.TrimSpace(line[:sepAt])
	value := strings.TrimSpace(line[sepAt+1:])
	if key == "" {
		return fmt.Errorf("'key' empty in line: %s", origLine)
	}
	if value == "" {
		in.inListKey = key
	} else {
		in.keyValues[key] = []string{value}
	}
	return
}
