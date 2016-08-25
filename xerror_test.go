package xerror_test

import (
	"fmt"
	"testing"

	errors "github.com/bearhang/xerror"
)

func func1() error {
	return errors.New("func1")
}

func func2() error {
	err := func1()
	return errors.New(err, "func2")
}

func func3() error {
	err := func2()
	return errors.New(err, "func3")
}

func TestDefault(t *testing.T) {
	err := func3()
	fmt.Println(err)
}
