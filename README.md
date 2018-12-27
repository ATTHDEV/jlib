# JSON parser for Go

Easy to use library for using json in golang

**Getting Started**
---
---

## Install

```shell
go get github.com/ATTHDEV/jlib/
```

## Import

```go
import (
  "github.com/ATTHDEV/jlib"
)
```

## Examples

### Create json object 

* create with string style

```go 
json := jlib.NewObject(`{ 
      "page": 1, 
      "fruits": ["apple", "peach"] ,
      "drink" : 
          {
            "name" : "coffee",
            "price" : 50
          }
      }`)
```

* create with object style

```go 
json := jlib.NewObject(jlib.Object{
      "page": 1,
      "fruits": jlib.Array{"apple", "peach"},
      "drink": jlib.Object{
        "name":  "coffee",
        "price": 50,
      },
})

```

### Read value
```go 
page := json.String("page")

fruits := json.Array("fruits")
apple := fruits.String(0)
peach := fruits.String(1)

drink := json.Object("drink")
price := drink.Int("price")

name := json.Object("drink").String("name")
```

### Add and update value
```go 
drink.Put("name", "cocoa") // update drink name
json.Put("Hobby" , jlib.Array{"play game", "watch movie"}) // add hobby
```

### Delete value
```go 
json.Delete("drink") //delete drink
json.Put("fruits", json.Array("fruits").DeleteAt(0)) // delete apple
```

### Loop through object

* loop with json object

```go
for k, v := range drink.ToMap() {
  ...
}
```

* loop with json array

```go
for _, f := range fruits.ToArray() {
   ...
}
```

### Convert to json String
```go 
json.ToString() // return json string
json.ToPrettyString() //  return Pretty json string
```

## Test
To run the project tests:

```shell
go test
```

**License**
---
---
MIT
