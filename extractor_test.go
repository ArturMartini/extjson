package gel

import (
	"errors"
	"fmt"
	"testing"
)

func init() {
	LoadFile("test/json_file2.json", "json_file2")
	LoadFile("test/json_file.json", "json_file")
}

func TestLoadFiles(t *testing.T) {
	err := LoadFile("test/json_file2.json", "json_file2")
	validate(t, nil, err)

	err = LoadFile("test/json_file.json", "json_file")
	validate(t, nil, err)
}

func TestGetValue(t *testing.T) {
	expected := "value"
	value := GetStr("key")
	validate(t, expected, value)
}

func TestGetStrInnerObject(t *testing.T) {
	expected := "value1"
	value := GetStr("key_obj.key1")
	validate(t, expected, value)
}

func TestGetStrInnerObjectLayerTwo(t *testing.T) {
	expected := "value3-2"
	value := GetStr("key_obj2.key_obj3.key3-2")
	validate(t, expected, value)
}

func TestGetInt(t *testing.T) {
	expected := 1
	value := GetInt("intKey")
	validate(t, expected, value)
}

func TestGetIntInnerObject(t *testing.T) {
	expected := 2
	value := GetInt("intKey_obj.key2")
	validate(t, expected, value)
}

func TestGetIntInnerObjectLayerTwo(t *testing.T) {
	expected := 2
	value := GetInt("intKey_obj2.intKey_obj3.key2")
	validate(t, expected, value)
}

func TestGetFloat(t *testing.T) {
	expected := 1.01
	value := GetFloat("floatKey")
	validate(t, expected, value)
}

func TestGetFloatInnerObject(t *testing.T) {
	expected := 2.01
	value := GetFloat("floatKey_obj.key2")
	validate(t, expected, value)
}

func TestGetFloatInnerObjectLayerTwo(t *testing.T) {
	expected := 2.02
	value  := GetFloat("floatKey_obj2.intKey_obj3.key2")
	validate(t, expected, value)
}

func TestGetStrComplexObject(t *testing.T) {
	expected := "complex"
	value := GetStr("complexObj.complexObj2.complexObj3.complexObj4.key1")
	validate(t, expected, value)
}

func TestMutipleFiles(t *testing.T) {
	SetContext("json_file2")
	expected := "value"
	value := GetStr("key2")
	validate(t, expected, value)

	expected = "value1"
	value = GetStr("key2_obj.key1")
	validate(t, expected, value)

	expected = "value2"
	value = GetStr("key2_obj.key2")
	validate(t, expected, value)
}

func TestLoadError(t *testing.T) {
	expected := errors.New("open not/exists: no such file or directory")
	err := LoadFile("not/exists", "notExists")
	validate(t, expected.Error(), err.Error())
}

func TestStrError(t *testing.T) {
	expected := ""
	value := GetStr("not.exists")
	validate(t, expected, value)
}

func TestIntError(t *testing.T) {
	expected := -1
	value := GetInt("not.exists")
	validate(t, expected, value)
}

func TestFloatError(t *testing.T) {
	expected := -1.00
	value := GetFloat("not.exists")
	validate(t, expected, value)
}

func validate(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		reportError(t, expected, actual)
	}
}

func reportError(t *testing.T, expected, actual interface{}) {
	t.Fatal(fmt.Sprintf("Test: %s, Expected: %s, Actual: %s", t.Name(), expected, actual))
}