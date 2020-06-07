# Json Extract Library (go-jel)
This is a library to simplify extract json values in golang

## Features 
* Extract json value by library  
* Casting values to int and float
* Support to reference multiple files

Usage:

```go
package main 
import (
    "github.com/arturmartini/jel"
    "fmt"
)

//Load file in memory 
err := jel.LoadFile("test/json_file.json", "json_file")

//Reference of file name and key property to return the string value
value := jel.GetStrValue("json_file", "key")

//Suport to get value in complex structure json using by '.' between keys
value := jel.GetStrValue("json_file", "key.key2")

//Suport to get int and float values 
intValue := jel.GetIntValue("json_file", "keyInt")
floatValue := jel.GetFloatValue("json_file", "keyFloat")

fmt.Println(value) // output "value"
fmt.Println(intValue) //output 1
fmt.Println(floatValue) //output 1.00
```