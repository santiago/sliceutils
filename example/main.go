package main

import (
	"fmt"
	"reflect"

	"github.com/santiago/sliceutils"
)

func main() {

	types := []interface{}{reflect.String,
		reflect.Int,
		reflect.Float32}

	for _, t := range types {

		tp := t.(reflect.Kind)

		a, err := sliceutils.GenerateSliceSlices(5, tp)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Generated random slice of '%s':\n", tp)
		fmt.Println(a)

		flat := sliceutils.Flatten(a)

		fmt.Printf("Flattened slice of '%s':\n", tp)
		fmt.Println(flat)
		fmt.Println("")

	}

}
