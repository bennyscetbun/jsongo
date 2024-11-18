jsongo
======

**Jsongo is a simple library for golang to help you build Json without static struct or map[string]interface**

[json.Marshal](http://golang.org/pkg/encoding/json/#Marshal) and [json.Unmarshal](http://golang.org/pkg/encoding/json/#Unmarshal) have never been that easy

**If you had only one function to look at, look at the "[At](#at)" function**

***If you want an easy way to turn your json into a structure you should use the "[Print](#print)" function after unmarshalling json in a Node***

***[2024/11/18] Project is not dead. We are using it quite often. We haven't found any necessary update***

You can find the doc on godoc.org [![GoDoc](https://godoc.org/github.com/bennyscetbun/jsongo?status.png)](https://godoc.org/github.com/bennyscetbun/jsongo)


## Node

Node is the basic Structure that you must use when using jsongo. It can either be a:
- Map (jsongo.NodeTypeMap)
- Array (jsongo.NodeTypeArray)
- Value (jsongo.NodeTypeValue) *Precisely a pointer store in an interface{}*
- Undefined (jsongo.NodeTypeUndefined) *default type*

*When a Node Type is set you cant change it without using Unset() first*
____
### Val
#### Synopsis:
turn this Node to NodeTypeValue and set that value
```go
func (that *Node) Val(val interface{}) 
```

#### Examples
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
	root.Val(42)
	root.DebugPrint("")
}
```
##### output:
```
42
```
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)
type MyStruct struct {
	Member1 string
	Member2 int
}

func main() {
	root := jsongo.Node{}
	root.Val(MyStruct{"The answer", 42})
	root.DebugPrint("")
}
```
##### output:
```
{
  "Member1": "The answer",
  "Member2": 42
}
```
_____
### Array
#### Synopsis:
 Turn this Node to a NodeTypeArray and/or set the array size (reducing size will make you loose data)
```go
func (that *Node) Array(size int) *[]Node
```

#### Examples
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
	a := root.Array(4)
    for i := 0; i < 4; i++ {
        (*a)[i].Val(i)
    }
	root.DebugPrint("")
}
```
##### output:
```
[
  0,
  1,
  2,
  3
]
```
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
    a := root.Array(4)
    for i := 0; i < 4; i++ {
        (*a)[i].Val(i)
    }
    root.Array(2) //Here we reduce the size and we loose some data
	root.DebugPrint("")
}
```
##### output:
```
[
  0,
  1
]
```
____
### Map
#### Synopsis:
Turn this Node to a NodeTypeMap and/or Create a new element for key if necessary and return it
```go
func (that *Node) Map(key string) *Node
```

#### Examples
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
    root.Map("there").Val("you are")
    root.Map("here").Val("you should be")
	root.DebugPrint("")
}
```
##### output:
```
{
  "here": "you should be",
  "there": "you are"
}
```
____
### At
#### Synopsis:
Helps you move through your node by building them on the fly

*val can be string or int only*

*strings are keys for NodeTypeMap*

*ints are index in NodeTypeArray (it will make array grow on the fly, so you should start to populate with the biggest index first)*
```go
func (that *Node) At(val ...interface{}) *Node
```

#### Examples
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
    root.At(4, "Who").Val("Let the dog out") //is equivalent to (*root.Array(5))[4].Map("Who").Val("Let the dog out")
    root.DebugPrint("")
}
```
##### output:
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
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func main() {
    root := jsongo.Node{}
    root.At(4, "Who").Val("Let the dog out")
    //to win some time you can even even save a certain Node
	node := root.At(2, "What")
	node.At("Can", "You").Val("do with that?")
	node.At("Do", "You", "Think").Val("Of that")
    root.DebugPrint("")
}
```
##### output:
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
____
### Print
#### Synopsis:
Helps you build your code by printing a go structure from the json you ve just unmarshaled

```go
func (that *Node) Print()
```

____
### Other Function
There is plenty of other function, you should check the complete doc [![GoDoc](https://godoc.org/github.com/bennyscetbun/jsongo?status.png)](https://godoc.org/github.com/bennyscetbun/jsongo)

#### A last Example for fun
##### code:
```go
package main

import (
    "github.com/bennyscetbun/jsongo"
)

func ShowOnlyValue(current *jsongo.Node) {
    switch current.GetType() {
    	case jsongo.NodeTypeValue:
			println(current.Get().(string))
		case jsongo.NodeTypeMap:
			for _, key := range current.GetKeys() {
				ShowOnlyValue(current.At(key))
			}
		case jsongo.NodeTypeArray:
			for _, key := range current.GetKeys() {
				ShowOnlyValue(current.At(key))
			}
	}
}

func main() {
    root := jsongo.Node{}
    root.At(4, "Who").Val("Let the dog out")
	node := root.At(2, "What")
	node.At("Can", "You").Val("do with that?")
	node.At("Do", "You", "Think").Val("Of that")
	ShowOnlyValue(&root)
}
```
##### output:
```
Of that
do with that?
Let the dog out
```
_____
_____
## Json Marshal/Unmarshal

One of the main purpose of jsongo was to create Json from data without using static structure or map[string]interface.

You can use the full power of the [encoding/json](http://golang.org/pkg/encoding/json/) package with jsongo.

### Marshal
#### Example
##### code:
```go
package main

import (
    "encoding/json"
	"fmt"
    "github.com/bennyscetbun/jsongo"
)

type Test struct {
	Static string `json:"static"`
	Over int `json:"over"`
}

func main() {
	root := jsongo.Node{}
	root.At("A", "AA", "AAA").Val(42)

	node := root.At("A", "AB")
	node.At(1).Val("Peace")
	node.At(0).Val(Test{"struct suck when you build json", 9000})
	root.At("B").Val("Oh Yeah")

	tojson, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
}
```
##### output:
```
{
  "A": {
    "AA": {
      "AAA": 42
    },
    "AB": [
      {
        "static": "struct suck when you build json",
        "over": 9000
      },
      "Peace"
    ]
  },
  "B": "Oh Yeah"
}
```
____
### Unmarshal
Unmarshal using Node follow some simple rules:
- Any NodeTypeUndefined Node will be set to the right type, any other type wont be changed
- Array will grow if necessary
- New keys will be added to Map
- Values set to nil "*.Val(nil)*" will be turn into the type decide by Json
- It will respect any current mapping and will return errors if needed

You can set a node as "DontExpand" with the UnmarshalDontExpand function and thoose rules will apply:
- The type wont be change for any type
- Array wont grow
- New keys wont be added to Map
- Values set to nil "*.Val(nil)*" will be turn into the type decide by Json
- It will respect any current mapping and will return errors if needed

#### Example of full expand
##### code:
```go
package main

import (
    "encoding/json"
    "github.com/bennyscetbun/jsongo"
    "fmt"
)

func main() {
    root := jsongo.Node{}
    fromjson := `{
	  "A": {
		"AA": {
		  "AAA": 42
		},
		"AB": [
		  {
			"static": "struct suck when you build json",
			"over": 9000
		  },
		  "Peace"
		]
	  },
	  "B": "Oh Yeah"
	}`
    err := json.Unmarshal([]byte(fromjson), &root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	root.DebugProspect(0, "\t")
}
```
##### output:
```
Is of Type: NodeTypeMap
A:
        Is of Type: NodeTypeMap
        AA:
                Is of Type: NodeTypeMap
                AAA:
                        Is of Type: NodeTypeValue
                        Value of type: float64
                        42
        AB:
                Is of Type: NodeTypeArray
                [0]:
                        Is of Type: NodeTypeMap
                        static:
                                Is of Type: NodeTypeValue
                                Value of type: string
                                struct suck when you build json
                        over:
                                Is of Type: NodeTypeValue
                                Value of type: float64
                                9000
                [1]:
                        Is of Type: NodeTypeValue
                        Value of type: string
                        Peace
B:
        Is of Type: NodeTypeValue
        Value of type: string
        Oh Yeah
```
#### Example expand with mapping
##### code:
```go
package main

import (
    "encoding/json"
    "github.com/bennyscetbun/jsongo"
    "fmt"
)
type Test struct {
    Static string `json:"static"`
    Over int `json:"over"`
}

func main() {
	root := jsongo.Node{}
    fromjson := `{
      "A": {
		"AA": {
		  "AAA": 42
		},
		"AB": [
		  {
			"static": "struct suck when you build json",
			"over": 9000
		  },
		  "Peace"
		]
	  },
	  "B": "Oh Yeah"
	}`
	root.At("A", "AB", 0).Val(Test{})
    err := json.Unmarshal([]byte(fromjson), &root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	root.DebugProspect(0, "\t")
}
```
##### output:
```
Is of Type: NodeTypeMap
A:
        Is of Type: NodeTypeMap
        AB:
                Is of Type: NodeTypeArray
                [0]:
                        Is of Type: NodeTypeValue
                        Value of type: main.Test
                        {Static:struct suck when you build json Over:9000}
                [1]:
                        Is of Type: NodeTypeValue
                        Value of type: string
                        Peace
        AA:
                Is of Type: NodeTypeMap
                AAA:
                        Is of Type: NodeTypeValue
                        Value of type: float64
                        42
B:
        Is of Type: NodeTypeValue
        Value of type: string
        Oh Yeah
```
#### Example expand with some UnmarshalDontExpand
##### code:
```go
package main

import (
    "encoding/json"
    "github.com/bennyscetbun/jsongo"
    "fmt"
)
type Test struct {
	Static string `json:"static"`
	Over int `json:"over"`
}

func main() {
    root := jsongo.Node{}
    fromjson := `{
      "A": {
		"AA": {
		  "AAA": 42
		},
		"AB": [
		  {
			"static": "struct suck when you build json",
			"over": 9000
		  },
		  "Peace"
		]
	  },
	  "B": "Oh Yeah"
	}`
	root.At("A", "AB").UnmarshalDontExpand(true, false).At(0).Val(Test{})
    err := json.Unmarshal([]byte(fromjson), &root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	root.DebugProspect(0, "\t")
}
```
##### output:
```
Is of Type: NodeTypeMap
A:
        Is of Type: NodeTypeMap
        AB:
                Is of Type: NodeTypeArray
                [0]:
                        Is of Type: NodeTypeValue
                        Value of type: main.Test
                        {Static:struct suck when you build json Over:9000}
        AA:
                Is of Type: NodeTypeMap
                AAA:
                        Is of Type: NodeTypeValue
                        Value of type: float64
                        42
B:
        Is of Type: NodeTypeValue
        Value of type: string
        Oh Yeah
```
