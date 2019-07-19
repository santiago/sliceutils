package sliceutils

import (
	"reflect"
	"testing"
)

func TestGenerateRandomInt(t *testing.T) {
	_, err := generateRandomInt()

	if err != nil {
		t.Error(err)
	}

}

func TestGenerateRandomFloat(t *testing.T) {
	_, err := generateRandomFloat()

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateRandomString(t *testing.T) {
	_, err := generateRandomString(RndStrSize)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateRandomSliceInt(t *testing.T) {
	_, err := generateRandomSlice(5, reflect.Int)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateRandomSliceFloat32(t *testing.T) {
	_, err := generateRandomSlice(5, reflect.Float32)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateRandomSliceFloat64(t *testing.T) {
	_, err := generateRandomSlice(5, reflect.Float64)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateRandomSliceString(t *testing.T) {
	_, err := generateRandomSlice(5, reflect.String)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateSliceSlicesInt(t *testing.T) {
	_, err := GenerateSliceSlices(5, reflect.Int)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateSliceSlicesFloat(t *testing.T) {
	_, err := GenerateSliceSlices(5, reflect.Float32)

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateSliceSlicesString(t *testing.T) {
	_, err := GenerateSliceSlices(5, reflect.String)

	if err != nil {
		t.Error(err)
	}
}

func TestInterfaceSlice(t *testing.T) {
	a := []interface{}{1,
		2,
		3,
		4}

	flat := interfaceSlice(a)

	if reflect.TypeOf(flat).Kind() != reflect.Slice {
		t.Errorf("Invalid return Kind for InterfaceSlice, got '%s' expected '%s'", reflect.TypeOf(flat).Kind(), reflect.Slice)
	}
}

func TestFlattenInt(t *testing.T) {

	a, err := GenerateSliceSlices(5, reflect.Int)

	if err != nil {
		t.Error(err)
	}

	flat := Flatten(a)

	if reflect.TypeOf(flat).Kind() != reflect.Slice {
		t.Errorf("Invalid return Kind for Flatten, got '%s' expected '%s'", reflect.TypeOf(flat).Kind(), reflect.Slice)
	}

	if len(flat) == 0 {
		t.Errorf("Empty result for Flatten, got length 0")
	}
}

func TestFlattenFloat(t *testing.T) {

	a, err := GenerateSliceSlices(5, reflect.Float32)

	if err != nil {
		t.Error(err)
	}

	flat := Flatten(a)

	if reflect.TypeOf(flat).Kind() != reflect.Slice {
		t.Errorf("Invalid return Kind for Flatten, got '%s' expected '%s'", reflect.TypeOf(flat).Kind(), reflect.Slice)
	}

	if len(flat) == 0 {
		t.Errorf("Empty result for Flatten, got length 0")
	}
}

func TestFlattenString(t *testing.T) {

	a, err := GenerateSliceSlices(5, reflect.String)

	if err != nil {
		t.Error(err)
	}

	flat := Flatten(a)

	if reflect.TypeOf(flat).Kind() != reflect.Slice {
		t.Errorf("Invalid return Kind for Flatten, got '%s' expected '%s'", reflect.TypeOf(flat).Kind(), reflect.Slice)
	}

	if len(flat) == 0 {
		t.Errorf("Empty result for Flatten, got length 0")
	}
}
