package validate

import (
	"fmt"
	"strings"

	"github.com/litesoft-go/utils/paths"
)

func Children(path *paths.Path, validatableFunctions ...func(*paths.Path) error) (err error) {
	for _, vf := range validatableFunctions {
		err = vf(path)
		if err != nil {
			return
		}
	}
	return
}

//noinspection GoUnusedExportedFunction
func OptionalInt32NotNegative(path *paths.Path, pathLeg string, currentValue *int32) error {
	if currentValue != nil {
		value := *currentValue
		if value < 0 {
			return path.Negative(pathLeg, value)
		}
	}
	return nil
}

func OptionalInt64NotNegative(path *paths.Path, pathLeg string, currentValue *int64) error {
	if currentValue != nil {
		value := *currentValue
		if value < 0 {
			return path.Negative(pathLeg, value)
		}
	}
	return nil
}

func OptionalNoLTWS(path *paths.Path, pathLeg, currentValue string) error {
	if (currentValue != "") && (strings.TrimSpace(currentValue) == "") {
		return path.LeadingTrailingWhitespace(pathLeg, currentValue)
	}
	return nil
}

func RequiredNoLTWS(path *paths.Path, pathLeg, currentValue string) error {
	if currentValue == "" {
		return path.Missing(pathLeg, currentValue)
	}
	if strings.TrimSpace(currentValue) == "" {
		return path.LeadingTrailingWhitespace(pathLeg, currentValue)
	}
	return nil
}

func RequiredOneOf(path *paths.Path, pathLeg string, currentValue interface{}, acceptableValues ...interface{}) error {
	why := "nothing is acceptable"
	if len(acceptableValues) != 0 {
		for _, acceptable := range acceptableValues {
			if currentValue == acceptable {
				return nil
			}
		}
		why = fmt.Sprintf("must be one of %v", acceptableValues)
	}
	return path.Plus(pathLeg).
		ErrorOf(why, currentValue)
}

func StringKeys(path *paths.Path, stringsMap map[string]string) (err error) {
	for k := range stringsMap {
		err = checkMapKey(path, k)
		if err != nil {
			return
		}
	}
	return
}

func MapStrings(path *paths.Path, stringsMap map[string]string) (err error) {
	for k, v := range stringsMap {
		err = checkMapKey(path, k)
		if err != nil {
			return
		}
		err = checkMapValue(path, k, v)
		if err != nil {
			return
		}
	}
	return
}

func checkMapKey(path *paths.Path, key string) error {
	if msg, currentValue := checkMapKorV("key", key); msg != "" {
		return path.
			ErrorOf(msg, currentValue)
	}
	return nil
}

func checkMapValue(path *paths.Path, key, value string) error {
	if msg, currentValue := checkMapKorV("value", value); msg != "" {
		return path.Plus("["+key+"]").
			ErrorOf(msg, currentValue)
	}
	return nil
}

func checkMapKorV(what, toCheck string) (msg, currentValue string) {
	if toCheck == "" {
		return "empty " + what, toCheck
	}
	if strings.TrimSpace(toCheck) == "" {
		return what + " with leading or trailing whitespace", "'" + toCheck + "'"
	}
	return
}
