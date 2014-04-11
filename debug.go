package jsongo

import (
	"encoding/json"
	"fmt"
	"os"
)

//DebugPrint Print a JSONNode as json withindent
func (that *JSONNode) DebugPrint(prefix string) {
	asJson, err := json.MarshalIndent(that, "", "  ")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(-1)
	}
	fmt.Printf("%s%s\n", prefix, asJson)
}
