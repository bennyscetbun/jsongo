jsongo
======

Jsongo is a simple library to help you build Json without static struct

You can find the doc on godoc.org [Here](http://godoc.org/github.com/Benny-Deluxe/jsongo) 

The first thing you should use is the At function and the Val function
### Example
#### simple

```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

type Test struct {
	Static string `json:"static"`
	Over int `json:"over"`
}

func main() {
	root := jsongo.JsonMap{}
	root.At("1", "1.1", "1.1.1").Val(42)
	root.At("1", "1.2", 0).Val(Test{"struct suck when you build json", 9000})
	root.At("1", "1.2", 1).Val("Peace")
	root.At("2").Val("Oh Yeah")
	tojson, err := json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
}
```

You can avoid to move with the At function all the time (and save some time)
### Example
#### simple Faster
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

type Test struct {
	Static string `json:"static"`
	Over int `json:"over"`
}

func main() {
	root := jsongo.JsonMap{}
	root.At("1", "1.1", "1.1.1").Val(42)

	node := root.At("1", "1.2")
	node.At(0).Val(Test{"struct suck when you build json", 9000})
	node.At(1).Val("Peace")

	root.At("2").Val("Oh Yeah")

	tojson, err := json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
}
```

You can use the Array function to set transform the current node into an array and set its size
### Example
#### simple even Faster
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

type Test struct {
	Static string `json:"static"`
	Over int `json:"over"`
}

func main() {
	root := jsongo.JsonMap{}
	root.At("1", "1.1", "1.1.1").Val(42)

	node := root.At("1", "1.2")
	nodeArray := node.Array(4)
	(*nodeArray)[0].Val(Test{"struct suck when you build json", 9000})
	(*nodeArray)[1].Val("Peace")

	root.At("2").Val("Oh Yeah")

	tojson, err := json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
}
```

You Beware that when you use At on a node you gonna set if it s a Map, a Array or a Value. You can use Unset to undo that.

### Examples
#### Common error
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

func main() {
	root := jsongo.JsonMap{}

	root.At("1", "2").Val(42)
	tojson, err := json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)

	root.At("1", 0).Val(42)
	tojson, err = json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)

	/*
	//You can even try that one :)
	root.At("1").Val(42)
	tojson, err = json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
	*/
}
```

#### Use of Unset
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

func main() {
	root := jsongo.JsonMap{}

	root.At("1", "2").Val(42)
	tojson, err := json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
	root.At("1").Unset()
	root.At("1", 0).Val(42)
	tojson, err = json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)

	//You can even try that one :)
	root.At("1").Unset()
	root.At("1").Val(42)
	tojson, err = json.Marshal(&root)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", tojson)
}
```

TODO
____

Unmarshal to be Unmarshaler compliant
AutoUnmarshal will help you build a Jsongo schema from a []byte json

