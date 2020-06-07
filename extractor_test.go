package jel

import (
	"fmt"
	"testing"
)

func init() {
	LoadFile("test/json_file.json", "json_file")
	LoadFile("test/json_file2.json", "json_file2")
}

func TestLoadFiles(t *testing.T) {
	err := LoadFile("test/json_file.json", "json_file")
	validate(t, nil, err)

	err = LoadFile("test/json_file2.json", "json_file2")
	validate(t, nil, err)
}

func TestGetValue(t *testing.T) {
	expected := "value"
	value := GetStrValue("json_file", "key")
	validate(t, expected, value)
}

func TestGetStrValueInnerObject(t *testing.T) {
	expected := "value1"
	value := GetStrValue("json_file", "key_obj.key1")
	validate(t, expected, value)
}

func TestGetStrValueInnerObjectLayerTwo(t *testing.T) {
	expected := "value3-2"
	value := GetStrValue("json_file", "key_obj2.key_obj3.key3-2")
	validate(t, expected, value)
}

func TestGetIntValue(t *testing.T) {
	expected := 1
	value := GetIntValue("json_file", "intKey")
	validate(t, expected, value)
}

func TestGetIntValueInnerObject(t *testing.T) {
	expected := 2
	value := GetIntValue("json_file", "intKey_obj.key2")
	validate(t, expected, value)
}

func TestGetIntValueInnerObjectLayerTwo(t *testing.T) {
	expected := 2
	value := GetIntValue("json_file", "intKey_obj2.intKey_obj3.key2")
	validate(t, expected, value)
}

func TestGetStrComplexObject(t *testing.T) {
	expected := "complex"
	value := GetStrValue("json_file", "complexObj.complexObj2.complexObj3.complexObj4.key1")
	validate(t, expected, value)
}

func TestMutipleFiles(t *testing.T) {
	expected := "value"
	value := GetStrValue("json_file2", "key2")
	validate(t, expected, value)

	expected = "value1"
	value = GetStrValue("json_file2", "key2_obj.key1")
	validate(t, expected, value)

	expected = "value2"
	value = GetStrValue("json_file2", "key2_obj.key2")
	validate(t, expected, value)
}

func validate(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		reportError(t, expected, actual)
	}
}

func reportError(t *testing.T, expected, actual interface{}) {
	t.Error(fmt.Sprintf("Test: %s, Expected: %s, Actual: %s", t.Name(), expected, actual))
}