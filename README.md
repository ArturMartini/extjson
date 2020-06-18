# Golang Extract Library (GEL)
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
  "keyFloat": 1.00,
  "list": [
    "1",
    "2"
  ]
}
```

File2:
```json
{
  "key": "value99"
}
```

Usage:
```go
package main 
import (
    "github.com/arturmartini/gel"
    "fmt"
)

func main() {
    //Load file in memory 
    err := gel.LoadFile("test/json_file.json", "file")
    
    //Reference of key property to return the string value
    value := gel.GetStr("key")
    
    //Suport to get value in complex structure json using by '.' between keys
    value2 := gel.GetStr("key1.key2")
    
    //Suport to get int and float values 
    intValue := gel.GetInt("keyInt")
    floatValue := gel.GetFloat("keyFloat")
    
    //Suport to get list of string values
    list := gel.GetList("list")

    
    //When load file automatically change context file
    err = gel.LoadFile("test/File2", "file2")
    otherValue := gel.GetStr("key")
    
    //Setting context to previous file
    gel.SetContext("file")
        
    fmt.Println(value)      // "value"
    fmt.Println(value2)     // "value2"
    fmt.Println(intValue)   // 1
    fmt.Println(floatValue) // 1.00
    fmt.Println(list)       // [1 2]
    fmt.Println(otherValue) // "value99"
}
```