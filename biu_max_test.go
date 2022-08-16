package biu

import (
	"fmt"
	"testing"
)

func TestGetInt8Max(t *testing.T) {
	max := GetInt8Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetUint8Max(t *testing.T) {
	max := GetUint8Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetInt16Max(t *testing.T) {
	max := GetInt16Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetUint16Max(t *testing.T) {
	max := GetUint16Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetInt32Max(t *testing.T) {
	max := GetInt32Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetUint32Max(t *testing.T) {
	max := GetUint32Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetInt64Max(t *testing.T) {
	max := GetInt64Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}

func TestGetUint64Max(t *testing.T) {
	max := GetUint64Max()
	sprint := fmt.Sprint(max)
	fmt.Println(fmt.Sprintf("the val: %s and len: %d.\n", sprint, len(sprint)))
}
