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
	root := jsongo.JSONNode{}
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
	root := jsongo.JSONNode{}
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
	root := jsongo.JSONNode{}
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
	root := jsongo.JSONNode{}

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
	root := jsongo.JSONNode{}

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

Unmarshal can Generate map as needed and will overwrite any JSONNode of type Value if needed
Except if you use UnmarshalDontGenerate to avoid Auto Generation while Unmarshaling, if you do so:
-If needed a JSONNode without type will be set to Value and set with the value of Unmarshal of the data
-JSONNode of type Array or Map wont generate new keys...


### Example
#### Unmarshal Without UnmarshalDontGenerate
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)

type Test struct {
	Static string `json:"static"`
	Over   int    `json:"over"`
}

func main() {
	root := jsongo.JSONNode{}
	root2 := jsongo.JSONNode{}
	
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
	fmt.Printf("root=>%s\n", tojson)
	root2.At("1", "1.1")
	err = json.Unmarshal(tojson, &root2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	tojson2, err := json.Marshal(&root2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("root2=>%s\n", tojson2)
}
```

#### Unmarshal With UnmarshalDontGenerate
```go
package main
import (
	"github.com/benny-deluxe/jsongo"
	"fmt"
)
type Test struct {
	Static string `json:"static"`
	Over   int    `json:"over"`
}

func main() {
	root := jsongo.JSONNode{}
	root2 := jsongo.JSONNode{}
	
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
	fmt.Printf("root=>%s\n", tojson)
	
	root2.At("1", "Tricks to set 1 as a map")
	root2.At("1").UnmarshalDontGenerate(true, true)
	err = json.Unmarshal(tojson, &root2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	tojson2, err := json.Marshal(&root2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("root2=>%s\n", tojson2)
}

**TODO**

-get keys and or iterate in JsonMaps

