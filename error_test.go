package utils

import "testing"

func TestError_GetErrStr(t *testing.T) {
	e := Error{
		ErrStr: " TEST_ERROR ",
	}
	if e.getErrStr() != "TEST_ERROR" {
		t.Errorf("Expected 'TEST_ERROR' but got '%s'", e.ErrStr)
	}
}

func TestError_GetErrMsg(t *testing.T) {
	e := Error{
		ErrMsg: "there was an error",
	}
	if e.getErrMsg() != e.ErrMsg {
		t.Errorf("Expected '%s', bug got '%s'", e.ErrMsg, e.getErrMsg())
	}
}

func TestError_NewError(t *testing.T) {
	e := Error{
		ErrStr: "",
		ErrMsg: "there was an error",
	}
	err := e.NewError()
	if err.Error() != "there was an error" {
		t.Errorf("Expected 'there was an error', but got '%s'", err.Error())
	}

	e = Error{
		ErrStr: " TEST_ERR ",
		ErrMsg: "there was an error",
	}
	expected := "TEST_ERR : there was an error"
	err = e.NewError()
	if err.Error() != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, err.Error())
	}
}
