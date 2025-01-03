package jsongo

import (
	"encoding/json"
	"fmt"
	"os"
)

// DebugPrint Print a Node as json withindent
func (that *Node) DebugPrint(prefix string) {
	asJSON, err := json.MarshalIndent(that, "", "  ")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(-1)
	}
	fmt.Printf("%s%s\n", prefix, asJSON)
}

func printfindent(indentlevel int, indentchar string, format string, args ...interface{}) {
	for i := 0; i < indentlevel; i++ {
		fmt.Printf("%s", indentchar)
	}
	fmt.Printf(format, args...)
}

func (that *Node) debugProspectValue(indentlevel int, indentchar string) {
	printfindent(indentlevel, indentchar, "Is of Type: NodeTypeValue\n")
	printfindent(indentlevel, indentchar, "Value of type: %T\n", that.Get())
	printfindent(indentlevel, indentchar, "%+v\n", that.Get())
}

func (that *Node) debugProspectMap(indentlevel int, indentchar string) {
	printfindent(indentlevel, indentchar, "Is of Type: NodeTypeMap\n")
	for key := range that.m {
		printfindent(indentlevel, indentchar, "%s:\n", key)
		that.m[key].DebugProspect(indentlevel+1, indentchar)
	}
}

func (that *Node) debugProspectArray(indentlevel int, indentchar string) {
	printfindent(indentlevel, indentchar, "Is of Type: NodeTypeArray\n")
	for key := range that.a {
		printfindent(indentlevel, indentchar, "[%d]:\n", key)
		that.a[key].DebugProspect(indentlevel+1, indentchar)
	}
}

// DebugProspect Print all the data the we ve got on a node and all it s children
func (that *Node) DebugProspect(indentlevel int, indentchar string) {
	switch that.t {
	case NodeTypeValue:
		that.debugProspectValue(indentlevel, indentchar)
	case NodeTypeMap:
		that.debugProspectMap(indentlevel, indentchar)
	case NodeTypeArray:
		that.debugProspectArray(indentlevel, indentchar)
	case NodeTypeUndefined:
		printfindent(indentlevel, indentchar, "Is of Type: NodeTypeUndefined\n")
	}
}
