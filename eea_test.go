package gocod

import "testing"

func TestEEA(t *testing.T) {
	ggt, s, tvar, err := EEA(124, 13)
	if err != nil {
		t.Error(err)
	}
	t.Log(ggt, s, tvar)
	if ggt != 1 || s != 2 || tvar != -19 {
		t.Error("wrong results")
	}
}

func TestEeaOpt(t *testing.T) {
	ggt, s, tvar, err := EeaOpt(124, 13)
	if err != nil {
		t.Error(err)
	}
	t.Log(ggt, s, tvar)
	if ggt != 1 || s != 2 || tvar != -19 {
		t.Error("wrong results")
	}
}
