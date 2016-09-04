package xerror_test

import (
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

type testcase struct {
	in  int    //flag
	out string //output error
}

var testcases = []testcase{
	{
		in:  errors.ELongFile,
		out: "",
	},
	{
		in:  errors.EShortFile,
		out: "",
	},
	{
		in:  errors.ELongFunc,
		out: "",
	},
	{
		in:  errors.EShortFunc,
		out: "",
	},
	{
		in:  errors.ELongFile | errors.ELongFunc,
		out: "",
	},
	{
		in:  errors.ELongFile | errors.EShortFunc,
		out: "",
	},
	{
		in:  errors.EShortFile | errors.ELongFunc,
		out: "",
	},
	{
		in:  errors.EShortFile | errors.EShortFunc,
		out: "",
	},
}

func TestDefault(t *testing.T) {
	err := func3()
	t.Log(err)
}

func TestFlag(t *testing.T) {
	for _, c := range testcases {
		err := errors.SetFlags(c.in)
		if err != nil {
			t.Error(err)
		}
		err = func3()
		t.Log("flag is ", c.in)
		t.Log(err)
	}
}
