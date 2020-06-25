package extjson

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

func TestGetList(t *testing.T) {
	expected := 2
	values := GetList("list")
	validate(t, expected, len(values))
	validate(t, "1", values[0])
	validate(t, "2", values[1])
}

func TestGetMap(t *testing.T) {
	expected := 2
	maps := GetMap("map")
	validate(t, expected, len(maps))
	validate(t, "value1", maps["key1"])
	validate(t, "value2", maps["key2"])
}

func TestGetComplexList(t *testing.T) {
	expected := 2
	values := GetList("complexList.list")
	validate(t, expected, len(values))
	validate(t, "3", values[0])
	validate(t, "4", values[1])
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

func TestCleanup(t *testing.T) {
	err := LoadFile("test/json_file.json", "json_file")
	validate(t, nil, err)

	v := GetStr("key")
	validate(t, "value", v)

	Cleanup()

	v2 := GetStr("key")
	validate(t, "", v2)
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

func TestListError(t *testing.T) {
	values := GetList("not.exists")
	validate(t, 0, len(values))
}

func TestMapError(t *testing.T) {
	values := GetMap("not.exists")
	validate(t, 0, len(values))
}

func TestFoundKey(t *testing.T) {
	SetContext("json_file")

	foundNotFound := FoundKey("keyNotfound")
	validate(t, false, foundNotFound)

	found := FoundKey("key")
	validate(t, true, found)

	foundComplex := FoundKey("complexObj.complexObj2")
	validate(t, true, foundComplex)

	foundInt := FoundKey("intKey_obj.key1")
	validate(t, true, foundInt)

	notFound := FoundKey("not.exists")
	validate(t, false, notFound)
}

func TestAdd(t *testing.T) {
	value := map[string]interface{}{
		"keyadd1": "valueadd1",
		"keyadd2": map[string]interface{}{
			"keyadd22": "valueadd2",
			"keyadd23": 1.0,
			"keyadd3": 2.01,
		},
	}

	Add(value)
	v1 := GetStr("keyadd1")
	validate(t, "valueadd1", v1)

	v2 := GetInt("keyadd2.keyadd23")
	validate(t, 1, v2)

	v3 := GetFloat("keyadd2.keyadd3")
	validate(t, 2.01, v3)
}

func validate(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		reportError(t, expected, actual)
	}
}

func reportError(t *testing.T, expected, actual interface{}) {
	t.Fatal(fmt.Sprintf("Test: %s, Expected: %s, Actual: %s", t.Name(), expected, actual))
}