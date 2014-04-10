// Copyright 2014 Benny Scetbun. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package Jsongo is a simple library to help you build Json without static struct
//
// Source code and project home:
// https://github.com/benny-deluxe/jsongo
//

package jsongo

import (
	"encoding/json"
	"errors"
)

//ErrorKeyAlreadyExist error if a key already exist in current JsonMap
var ErrorKeyAlreadyExist = errors.New("jsongo key already exist")

//ErrorMultipleType error if a JsonMap already got a different type of value
var ErrorMultipleType = errors.New("jsongo this node is already set to a different type (Map, Array, Value)")

//ErrorArrayNegativeValue error if you ask for a negative index in an array
var ErrorArrayNegativeValue = errors.New("jsongo negative index for array")

//ErrorArrayNegativeValue error if you ask for a negative index in an array
var ErrorAtUnsupportedType = errors.New("jsongo Unsupported Type as At argument")

//JsonMap Datastructure to build and maintain Nodes
type JsonMap struct {
	m map[string]*JsonMap
	a []JsonMap
	v interface{}
	t Type //Type of that jsonMap 0: Not defined, 1: map, 2: array, 3: value
}

type Type int
const (
	TypeUndefined Type = iota
	TypeMap
	TypeArray
	TypeValue
)

//At At helps you move through your node by building them on the fly
//val can be string or int values
//string are keys for map in json
//int are index in array in json
func (that *JsonMap) At(val ...interface{}) *JsonMap {
	if len(val) == 0 {
		return that
	}
	switch vv := val[0].(type) {
	case string:
		return that.atMap(vv, val[1:]...)
	case int:
		return that.atArray(vv, val[1:]...)
	}
	panic(ErrorAtUnsupportedType)
}

//atMap return the JsonMap in current map
func (that *JsonMap) atMap(key string, val ...interface{}) *JsonMap {
	if that.t != TypeUndefined && that.t != TypeMap {
		panic(ErrorMultipleType)
	}
	if that.m == nil {
		that.m = make(map[string]*JsonMap)
		that.t = TypeMap
	}
	if next, ok := that.m[key]; ok {
		return next.At(val...)
	}
	that.m[key] = new(JsonMap)
	return that.m[key].At(val...)
}

//atArray return the JsonMap in current TypeArray (and make it grow if necessary)
func (that *JsonMap) atArray(key int, val ...interface{}) *JsonMap {
	if that.t == TypeUndefined {
		that.t = TypeArray
	} else if that.t != TypeArray {
		panic(ErrorMultipleType)
	}
	if key < 0 {
		panic(ErrorArrayNegativeValue)
	}
	if key >= len(that.a) {
		newa := make([]JsonMap, key+1)
		for i := 0; i < len(that.a); i++ {
			newa[i] = that.a[i]
		}
		that.a = newa
	}
	/*	if that.a[key] == nil {
		that.a[key] = new(JsonMap)
	}*/
	return that.a[key].At(val...)
}

//TypeMap Turn this node to a map and Create a new element for key
func (that *JsonMap) TypeMap(key string) *JsonMap {
	if that.t != TypeUndefined && that.t != TypeMap {
		panic(ErrorMultipleType)
	}
	if that.m == nil {
		that.m = make(map[string]*JsonMap)
		that.t = TypeMap
	}
	if _, ok := that.m[key]; ok {
		panic(ErrorKeyAlreadyExist)
	}
	that.m[key] = &JsonMap{}
	return that.m[key]
}

//TypeArray Turn this node to an array and/or set array size (reducing size will make you loose data)
func (that *JsonMap) TypeArray(size int) *[]JsonMap {
	if that.t == TypeUndefined {
		that.t = TypeArray
	} else if that.t != TypeArray {
		panic(ErrorMultipleType)
	}
	if size < 0 {
		panic(ErrorArrayNegativeValue)
	}
	var min int
	if size < len(that.m) {
		min = size
	} else {
		min = len(that.m)
	}
	newa := make([]JsonMap, size)
	for i := 0; i < min; i++ {
		newa[i] = that.a[i]
	}
	that.a = newa
	return &(that.a)
}

//Val Turn this node to user value and set that user value
func (that *JsonMap) Val(val interface{}) {
	if that.t == TypeUndefined {
		that.t = TypeValue
	} else if that.t != TypeValue {
		panic(ErrorMultipleType)
	}
	that.v = val
}

//Unset Will unset the node. All the children data will be lost
func (that *JsonMap) Unset() {
	*that = JsonMap{}
}

//MarshalJSON Make JsonMap a Marshaler Interface compatible
func (that *JsonMap) MarshalJSON() ([]byte, error) {
	var ret []byte
	var err error
	switch that.t {
	case TypeMap:
		ret, err = json.Marshal(that.m)
	case TypeArray:
		ret, err = json.Marshal(that.a)
	case TypeValue:
		ret, err = json.Marshal(that.v)
	default:
		ret, err = json.Marshal(nil)
	}
	if err != nil {
		return nil, err
	}
	return ret, err
}

/*func (that *JsonMap) UnmarshalJSON(data []byte) error {
	println("YOUHOU")
	return nil
}*/
