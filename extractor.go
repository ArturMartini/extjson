package gel

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
	files map[string]file
}

type file struct {
	values map[string]interface{}
}

type Type int

var instance = ext{files: map[string]file{}}

func GetStrValue(file, key string) string {
	var value string
	var values interface{}
	values = instance.files[file].values
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

		if idx >= len(breadPath){
			continue
		}

		value = str
	}
	return value
}

func GetIntValue(file, key string) int {
	var value float64
	var values interface{}
	values = instance.files[file].values
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

		if idx >= len(breadPath){
			continue
		}

		value = integer
	}
	return int(math.Round(value))
}

func GetFloatValue(file, key string) float64 {
	var value float64
	var values interface{}
	values = instance.files[file].values
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

		if idx >= len(breadPath){
			continue
		}

		value = float
	}
	return value
}

func LoadFile(path, name string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when try load file: " + path)
		return err
	}

	var values map[string]interface{}
	err = json.Unmarshal(bytes, &values)
	if err != nil {
		log.Fatal("Error when try unmarshal json: " + path + "\nError: " + err.Error())
		return err
	}

	instance.files[name] = file{values: values}
	return nil
}