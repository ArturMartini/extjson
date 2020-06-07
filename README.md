# Goland Extract Library (GEL)
This is a library to simplify extract json values in golang

## Features 
* Extract json value by library  
* Casting values to int and float
* Support to reference multiple files


File:
```json
{
  "key": "value",
  "key1": {
    "key2": "value2"
  },
  "keyInt": 1,
  "keyFloat": 1.00
}
```

Usage:
```go
package main 
import (
    "github.com/arturmartini/jel"
    "fmt"
)

//Load file in memory 
err := gel.LoadFile("test/json_file.json", "alias_file")

//Reference of file name and key property to return the string value
value := gel.GetStrValue("alias_file", "key")

//Suport to get value in complex structure json using by '.' between keys
value2 := gel.GetStrValue("alias_file", "key1.key2")

//Suport to get int and float values 
intValue := gel.GetIntValue("alias_file", "keyInt")
floatValue := gel.GetFloatValue("alias_file", "keyFloat")

fmt.Println(value) // output "value"
fmt.Println(value2) // output "value2"
fmt.Println(intValue) //output 1
fmt.Println(floatValue) //output 1.00
```