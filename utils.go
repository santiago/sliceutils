package sliceutils

import (
	"fmt"
	"reflect"
	// Safer random number generation
	"crypto/rand"
	"encoding/base64"
)

const RndStrSize = 3

// Flatten flattens a slice of type interface{} which might contain
// subslices. The result is a plain slice of type interface{}
func Flatten(arr []interface{}) []interface{} {

	// Prepare the result
	result := []interface{}{}

	// Iterate over the slice
	for _, el := range arr {

		// Check if the Kind of the element is not slice
		if reflect.TypeOf(el).Kind() != reflect.Slice {

			// Append the element to the resulting array
			result = append(result, el)

		} else {
			// If the Kind of the element is a slice

			// Convert Interface into a slice of interface{}
			iel := interfaceSlice(el)

			// Flatten the element-slice
			sub := Flatten(iel)

			// Iterate over the elements of the flattened element-slice
			// append them to the resulting slice
			for _, v := range sub {
				result = append(result, v)
			}

		}

	}

	return result
}

// interfaceSlice converts an interface type slice into a
// slice of interface{} type
func interfaceSlice(slice interface{}) []interface{} {

	// Get the value of the input interface type
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {

		// Panic if the Kind is not slice, wont deal with
		// elaborated data types such as maps or channels
		panic("interfaceSlice() given a non-slice type")
	}

	// Create a slice of interface{} with the same size
	// as the original slice
	ret := make([]interface{}, s.Len())

	// Iterate over the slice elements
	// and populate the resulting slice
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	bytesRead := make([]byte, n)
	_, err := rand.Read(bytesRead)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return bytesRead, nil
}

func generateRandomInt() (int, error) {
	bytesRead := make([]byte, 1)
	_, err := rand.Read(bytesRead)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return 0, err
	}

	return int(bytesRead[0]), nil
}

func generateRandomFloat() (float32, error) {
	bytesRead := make([]byte, RndStrSize)
	_, err := rand.Read(bytesRead)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return 0, err
	}

	return float32((bytesRead[0] + bytesRead[1] + bytesRead[2])) / 3.0, nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func generateRandomString(bytesLong int) (string, error) {
	randomBytes, err := generateRandomBytes(bytesLong)
	return base64.URLEncoding.EncodeToString(randomBytes), err
}

func generateRandomSlice(maxItems int, kind reflect.Kind) ([]interface{}, error) {

	ret := []interface{}{}

	switch kind {
	case reflect.String:

		for i := 0; i < maxItems; i++ {
			s, _ := generateRandomString(RndStrSize)
			ret = append(ret, s)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		for i := 0; i < maxItems; i++ {
			s, _ := generateRandomInt()
			ret = append(ret, s)
		}

	case reflect.Float32, reflect.Float64:

		for i := 0; i < maxItems; i++ {
			s, _ := generateRandomFloat()
			ret = append(ret, s)
		}
	default:

		return ret, fmt.Errorf("Unsupported kind '%s'", kind)
	}

	return ret, nil
}

// GenerateSliceSlices generates a slice of elements and slices of kind kind
func GenerateSliceSlices(maxItems int, kind reflect.Kind) ([]interface{}, error) {
	ret := []interface{}{}

	switch kind {
	case reflect.String:

		for i := 0; i < maxItems; i++ {
			ri, _ := generateRandomInt()

			if ri > 150 {
				s, _ := generateRandomString(RndStrSize)
				ret = append(ret, s)
			} else if ri > 80 && ri < 150 {
				s1, _ := generateRandomSlice(int(ri/8), kind)
				s2, _ := generateRandomSlice(int(ri/80), kind)
				sub := []interface{}{s1, s2}
				ret = append(ret, sub)

			} else {
				s, _ := generateRandomSlice(int(ri/10), kind)
				ret = append(ret, s)
			}

		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		for i := 0; i < maxItems; i++ {
			ri, _ := generateRandomInt()

			if ri > 150 {
				s, _ := generateRandomInt()
				ret = append(ret, s)
			} else if ri > 80 && ri < 150 {
				s1, _ := generateRandomSlice(int(ri/8), kind)
				s2, _ := generateRandomSlice(int(ri/80), kind)
				sub := []interface{}{s1, s2}
				ret = append(ret, sub)

			} else {
				s, _ := generateRandomSlice(int(ri/10), kind)
				ret = append(ret, s)
			}

		}

	case reflect.Float32, reflect.Float64:

		for i := 0; i < maxItems; i++ {
			ri, _ := generateRandomInt()

			if ri > 150 {
				s, _ := generateRandomFloat()
				ret = append(ret, s)
			} else if ri > 80 && ri < 150 {
				s1, _ := generateRandomSlice(int(ri/8), kind)
				s2, _ := generateRandomSlice(int(ri/80), kind)
				sub := []interface{}{s1, s2}
				ret = append(ret, sub)

			} else {
				s, _ := generateRandomSlice(int(ri/10), kind)
				ret = append(ret, s)
			}

		}
	default:

		return ret, fmt.Errorf("Unsupported kind '%s'", kind)
	}

	return ret, nil
}
