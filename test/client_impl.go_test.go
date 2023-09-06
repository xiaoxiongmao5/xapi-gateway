package test

import "testing"

func TestRunClientImpl(t *testing.T) {
	res, err := RunClientImpl()
	if err != nil {
		t.Error("err=", err)
	}
	if !res {
		t.Error("TestRunClientImpl 返回false")
	}
}
