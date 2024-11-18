package jsongo

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//Thanks https://github.com/chuckpreslar/inflect for the UpperCamelCase

// Split's a string so that it can be converted to a different casing.
// Splits on underscores, hyphens, spaces and camel casing.
func split(str string) []string {
	// FIXME: This isn't a perfect solution.
	// ex. WEiRD CaSINg (Support for 13 year old developers)
	return strings.Split(regexp.MustCompile(`-|_|([a-z])([A-Z])`).ReplaceAllString(strings.Trim(str, `-|_| `), `$1 $2`), ` `)
}

// UpperCamelCase converts a string to it's upper camel case version.
func UpperCamelCase(str string) string {
	pieces := split(str)

	for index, s := range pieces {
		pieces[index] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(s[0])), strings.ToLower(s[1:]))
	}

	return strings.Join(pieces, ``)
}

func (that *Node) printValue(indentlevel int, indentchar string) {
	fmt.Printf(" %T ", that.Get())
}

func (that *Node) printMap(indentlevel int, indentchar string) {
	fmt.Printf(" struct {\n")
	//order keys to make print consistent
	keys := make([]string, 0, len(that.m))

	for key := range that.m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		printfindent(indentlevel+1, indentchar, "%s", UpperCamelCase(key))
		that.m[key].print(indentlevel+1, indentchar)
		fmt.Printf(" `json:\"%s\"`\n", key)
	}
	printfindent(indentlevel, indentchar, "}")
}

func (that *Node) printArray(indentlevel int, indentchar string) {
	if len(that.a) == 0 {
		fmt.Printf(" []interface{} ")
		return
	}
	fmt.Printf(" [] ")
	for key := range that.a {
		that.a[key].print(indentlevel+1, indentchar)
		break
	}
}

// DebugProspect Print all the data the we ve got on a node and all it s children
func (that *Node) print(indentlevel int, indentchar string) {
	switch that.t {
	case NodeTypeValue:
		that.printValue(indentlevel, indentchar)
	case NodeTypeMap:
		that.printMap(indentlevel, indentchar)
	case NodeTypeArray:
		that.printArray(indentlevel, indentchar)
	case NodeTypeUndefined:
		printfindent(indentlevel, indentchar, "Is of Type: NodeTypeUndefined\n")
	}
}

// Print Print all the data the we ve got on a node and all it s children as a go struct :) (FOR DEV PURPOSE)
func (that *Node) Print() {
	that.print(0, "\t")
	fmt.Printf("\n")
}
