jsongo
======

**Jsongo is a simple library to help you build Json without static struct**

[json.Marshal](http://golang.org/pkg/encoding/json/#Marshal) and [json.Unmarshal](http://golang.org/pkg/encoding/json/#Unmarshal) have never been that easy

**If you had only one function to look at, look at the "[At](#at)" function**

You can find the doc on godoc.org [Here](http://godoc.org/github.com/Benny-Deluxe/jsongo)


##JsonNode

JsonNode is the basic Structure that you must use when using jsongo. It can either be a:
- Map (jsongo.TypeMap)
- Array (jsongo.TypeArray)
- Value (jsongo.TypeValue) *Precisely a pointer store in an interface{}*
- Undefined (jsongo.TypeUndefined) *default type*

*When a JSONNode Type is set you cant change it without using Unset() first*
____
###Val
####Synopsis:
turn this JSONNode to TypeValue and set that value (val must be a pointer)
```go
func (that *JSONNode) Val(val interface{}) 
```

####Examples
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
	root := jsongo.JSONNode{}
	root.Val(42)
	root.DebugPrint("")
}
```
#####output:
```
42
```
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)
type MyStruct struct {
	Member1 string
	Member2 int
}

func main() {
	root := jsongo.JSONNode{}
	root.Val(MyStruct{"The answer", 42})
	root.DebugPrint("")
}
```
#####output:
```
{
  "Member1": "The answer",
  "Member2": 42
}
```
_____
###Array
####Synopsis:
 Turn this JSONNode to a TypeArray and/or set the array size (reducing size will make you loose data)
```go
func (that *JSONNode) Array(size int) *[]JSONNode
```

####Examples
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
	a := root.Array(4)
    for i := 0; i < 4; i++ {
        (*a)[i].Val(i)
    }
	root.DebugPrint("")
}
```
#####output:
```
[
  0,
  1,
  2,
  3
]
```
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
    a := root.Array(4)
    for i := 0; i < 4; i++ {
        (*a)[i].Val(i)
    }
    root.Array(2) //Here we reduce the size and we loose some data
	root.DebugPrint("")
}
```
#####output:
```
[
  0,
  1
]
```
____
###Map
####Synopsis:
Turn this JSONNode to a TypeMap and/or Create a new element for key if necessary and return it
```go
func (that *JSONNode) Map(key string) *JSONNode
```

####Examples
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
    root.Map("there").Val("you are")
    root.Map("here").Val("you should be")
	root.DebugPrint("")
}
```
#####output:
```
{
  "here": "you should be",
  "there": "you are"
}
```
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
    root.Map("here").Val("you are")
    root.Map("here").Val("you gonna PANIC")
	root.DebugPrint("")
}
```
#####output:
```
panic: jsongo key already exist
```
____
###At
####Synopsis:
Helps you move through your node by building them on the fly

*val can be string or int only*

*strings are keys for TypeMap*

*ints are index in TypeArray (it will make array grow on the fly, so you should start to populate with the biggest index first)*
```go
func (that *JSONNode) At(val ...interface{}) *JSONNode
```

####Examples
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
    root.At(4, "Who").Val("Let the dog out") //is equivalent to (*root.Array(5))[4].Map("Who").Val("Let the dog out")
    root.DebugPrint("")
}
```
#####output:
```
[
  null,
  null,
  null,
  null,
  {
    "Who": "Let the dog out"
  }
]
```
#####code:
```go
package main

import (
    "github.com/benny-deluxe/jsongo"
)

func main() {
    root := jsongo.JSONNode{}
    root.At(4, "Who").Val("Let the dog out")
    //to win some time you can even even save a certain JSONNode
	node := root.At(2, "What")
	node.At("Can", "You").Val("do with that?")
	node.At("Do", "You", "Think").Val("Of that")
    root.DebugPrint("")
}
```
#####output:
```
[
  null,
  null,
  {
    "What": {
      "Can": {
        "You": "do with that?"
      },
      "Do": {
        "You": {
          "Think": "Of that"
        }
      }
    }
  },
  null,
  {
    "Who": "Let the dog out"
  }
]
```