package main

import (
	"fmt"
	"maps"
)

func startMaps() {
	var mV0 map[string]string
	mV1 := map[string]string{}
	mV2 := map[string]string{"one": "first number", "two": "second number"}
	mV3 := make(map[string]string, 8)
	mV4 := mV2                 // reference copied
	mV5 := map[string]string{} // truthy copy
	for k, v := range mV2 {
		mV5[k] = v
	}
	// var mV6 map[string]string
	// maps.Copy(mV6, mV2) // panic: assignment to entry in nil map
	mV6 := make(map[string]string, len(mV2))
	maps.Copy(mV6, mV2)

	fmt.Printf("\nvar mV0 map[string]string\nmV0 = %#v\n", mV0)
	fmt.Printf("\nmV1 := map[string]string{}\nmV1 = %#v\n", mV1)
	fmt.Printf(
		"\nmV2 := map[string]string{\"one\": \"first number\", \"two\": \"second number\"}\nmV2 = %#v\n",
		mV2,
	)
	fmt.Printf("\nmV3 := make(map[string]string, 8)\nmV3 = %#v\n", mV3)
	fmt.Printf("\nmV4 := mV2\nmV4 = %#v\n", mV4)
	fmt.Printf(
		"\nmV5 := map[string]string{}\nfor k, v := range mV2 {\n\tmV5[k] = v\n}\nmV5 = %#v\n",
		mV5,
	)
	fmt.Printf("\nmV6 := make(map[string]string, len(mV2))\nmaps.Copy(mV6, mV2)\nmV6 = %#v\n", mV6)

	fmt.Printf(
		"\nlen(mV0) = %d\nlen(mV1) = %d\nlen(mV2) = %d\nlen(mV3) = %d\n",
		len(mV0), len(mV1), len(mV2), len(mV3),
	)

	// fmt.Printf("\nCOMPARE\nmV1 == mV0: %t\n", mV1 == mV0) // compiler error: invalid operation: mV1 == mV0 (map can only be compared to nil)
	fmt.Printf("\nCOMPARE\nmV0 == nil: %t\nmV1 == nil: %t\nmV3 == nil: %t\n", mV0 == nil, mV1 == nil, mV3 == nil)

	// mV0["one"] = "first element"
	// fmt.Printf("\nmV0[\"one\"] = \"first element\"\nmV0 = %#v\n", mV0) // panic: assignment to entry in nil map

	mV1["one"] = "first element"
	fmt.Printf("\nAFTER\nmV1[\"one\"] = \"first element\"\nmV1 = %#v\n", mV1)

	mV2["three"] = "third number"
	fmt.Printf("\nAFTER\nmV2[\"three\"] = \"third number\"\nmV2 = %#v\nmV4 = %#v\nmV5 = %#v\nmV6 = %#v\n", mV2, mV4, mV5, mV6)

	mV3["one"] = "first"
	fmt.Printf("\nAFTER\nmV3[\"one\"] = \"first\"\nmV3 = %#v\n", mV3)

	mV4["one"] = "change first number"
	fmt.Printf("\nAFTER\nmV4[\"one\"] = \"change first number\"\nmV2 = %#v\nmV4 = %#v\nmV5 = %#v\nmV6 = %#v\n", mV2, mV4, mV5, mV6)

	mV5["two"] = "change seconde number"
	fmt.Printf("\nAFTER\nmV5[\"two\"] = \"change seconde number\"\nmV2 = %#v\nmV4 = %#v\nmV5 = %#v\nmV6 = %#v\n", mV2, mV4, mV5, mV6)

	delete(mV2, "two")
	fmt.Printf("\nAFTER\ndelete(mV2, \"two\")\nmV2 = %#v\nmV4 = %#v\nmV5 = %#v\nmV6 = %#v\n", mV2, mV4, mV5, mV6)

	delete(mV0, "empty")
	fmt.Printf("\nAFTER\ndelete(mV0, \"empty\")\nmV0 = %#v\n", mV0)

	fmt.Printf("\nfor k,v := range mV2 {\n")
	for k, v := range mV2 {
		fmt.Printf("\tmV2[%v] = %v\n", k, v)
	}
	fmt.Print("}\n")

	mV1 = nil
	fmt.Printf("\nAFTER\nmV1 = nil\nmV1 = %#v\n", mV1)
	fmt.Printf("\nCOMPARE\nmV1 == nil: %t\n", mV1 == nil)

	fmt.Printf("\nBEFORE\nfor k := range mV5 {\tdelete(mV5, k)\t}\nmV5 = %#v\n", mV5)
	for k := range mV5 {
		delete(mV5, k)
	}
	fmt.Printf("\nAFTER\nfor k := range mV5 {\tdelete(mV5, k)\t}\nmV5 = %#v\n", mV5)
	fmt.Printf("\nCOMPARE\nmV5 == nil: %t\n", mV5 == nil)
}
