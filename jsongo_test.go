package jsongo

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMarshaling(t *testing.T) {
	tests := map[string]func() Node{
		`189`: func() (ret Node) {
			ret.Val(189)
			return
		},
		`178.2`: func() (ret Node) {
			ret.Val(178.2)
			return
		},
		`0`: func() (ret Node) {
			ret.Val(0)
			return
		},
		`""`: func() (ret Node) {
			ret.Val("")
			return
		},
		`"bla"`: func() (ret Node) {
			ret.Val("bla")
			return
		},
		`null`: func() (ret Node) {
			ret.Val((*string)(nil))
			return
		},
		`{}`: func() (ret Node) {
			ret.SetType(NodeTypeMap)
			return
		},
		`{"1":0}`: func() (ret Node) {
			ret.At("1").Val(0)
			return
		},
		`{"1":null}`: func() (ret Node) {
			ret.At("1").Val((*int)(nil))
			return
		},
		`[]`: func() (ret Node) {
			ret.SetType(NodeTypeArray)
			return
		},
		`[0,null,[],{},"youpi"]`: func() (ret Node) {
			ret.At(0).Val(0)
			ret.At(1).Val((*int)(nil))
			ret.At(2).SetType(NodeTypeArray)
			ret.At(3).SetType(NodeTypeMap)
			ret.At(4).Val("youpi")
			return
		},
		`[null,null,0]`: func() (ret Node) {
			ret.At(2).Val(0)
			return
		},
	}

	for expected, f := range tests {
		t.Run(expected, func(t *testing.T) {
			node := f()
			data, err := json.Marshal(&node)
			if err != nil {
				t.Error(err)
				return
			}
			if strings.Compare(string(data), expected) != 0 {
				t.Errorf("%s is not equal to expected %s", string(data), expected)
			}
		})
	}
}
