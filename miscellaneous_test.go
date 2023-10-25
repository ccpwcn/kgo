package kgo

import (
	"fmt"
	"testing"
)

func TestGetRunFuncName(t *testing.T) {
	name := RunFuncName()
	fmt.Println(name)
}
