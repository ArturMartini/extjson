package extjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const (
	jsonType Type = 1
)

type ext struct {
	file file
}

type file struct {
	values map[string]interface{}
}

type Type int

var instance = ext{}
var files = map[string]file{}

func GetStr(key string) string {
	var value string
	var values interface{}
	values = instance.file.values
	breadPath := strings.Split(key, ".")

	for idx, p := range breadPath {
		v, ok := values.(map[string]interface{})
		if ok {
			values = v[p]
		}

		str, ok := values.(string)
		if !ok {
			continue
		}

		if idx >= len(breadPath) {
			continue
		}

		value = str
	}

	return value
}

func GetInt(key string) int {
	var value float64
	var values interface{}
	var hasValue bool
	values = instance.file.values
	breadPath := strings.Split(key, ".")

	for idx, p := range breadPath {
		v, ok := values.(map[string]interface{})
		if ok {
			values = v[p]
		}

		integer, ok := values.(float64)
		if !ok {
			continue
		}

		if idx >= len(breadPath) {
			continue
		}

		hasValue = true
		value = integer
	}

	if !hasValue {
		return -1
	}

	return int(math.Round(value))
}

func GetFloat(key string) float64 {
	var value float64
	var values interface{}
	var hasValue bool
	values = instance.file.values
	breadPath := strings.Split(key, ".")

	for idx, p := range breadPath {
		v, ok := values.(map[string]interface{})
		if ok {
			values = v[p]
		}

		float, ok := values.(float64)
		if !ok {
			continue
		}

		if idx >= len(breadPath) {
			continue
		}

		hasValue = true
		value = float
	}

	if !hasValue {
		return -1.00
	}
	return value
}

func GetList(key string) []string {
	var valueList []string
	var values interface{}
	var hasValue bool
	values = instance.file.values
	breadPath := strings.Split(key, ".")

	for idx, p := range breadPath {
		v, ok := values.(map[string]interface{})
		if ok {
			values = v[p]
		}

		list, ok := values.([]interface{})
		if !ok {
			continue
		}

		if idx >= len(breadPath) {
			continue
		}

		hasValue = true
		for _, vList := range list {
			valueList = append(valueList, vList.(string))
		}
	}

	if !hasValue {
		return nil
	}
	return valueList
}

func GetMap(key string) map[string]string {
	maps := map[string]string{}
	var values interface{}
	var hasValue bool
	values = instance.file.values
	breadPath := strings.Split(key, ".")

	for idx, p := range breadPath {
		vMap, ok := values.(map[string]interface{})
		if ok {
			values = vMap[p]
		} else {
			continue
		}

		notStr := false
		if len(vMap) > 0 && values != nil {
			for k, v := range values.(map[string]interface{}) {
				vStr, ok := v.(string)
				if !ok {
					notStr = true
				} else {
					maps[k] = vStr
					hasValue = true
				}
			}
		}

		if notStr {
			continue
		}

		if idx >= len(breadPath) {
			continue
		}
	}

	if !hasValue {
		return nil
	}
	return maps
}

func LoadFile(path, name string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error when try load file: " + path)
		return err
	}

	var values map[string]interface{}
	err = json.Unmarshal(bytes, &values)
	if err != nil {
		log.Println("Error when try unmarshal json: " + path + "\nError: " + err.Error())
		return err
	}
	file := file{values: values}
	files[name] = file
	instance.file = file
	return nil
}

func FoundKey(key string) bool {
	var values interface{}
	values = instance.file.values
	breadPath := strings.Split(key, ".")
	found := false

	for idx, p := range breadPath {
		hasValue := false
		v, ok := values.(map[string]interface{})
		if ok {
			hasValue = true
			values = v[p]
		}


		if idx + 1 >= len(breadPath) && hasValue {
			found = true
		}
	}

	return found
}

func Add(value map[string]interface{}) {
	for k, v := range value {
		instance.file.values[k] = v
	}
}

func SetContext(name string) {
	if f, ok := files[name]; ok {
		instance.file = f
	}
}

func Cleanup() {
	instance.file = file{map[string]interface{}{}}
}
