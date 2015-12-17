package utils

import "testing"

func TestMod(t *testing.T) {
	if Mod(0, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(1, 15) != 1 {
		t.Errorf("Error")
	}
	if Mod(15, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(-1, 15) != 14 {
		t.Errorf("Error")
	}
	if Mod(-15, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(-16, 15) != 14 {
		t.Errorf("Error")
	}
}

func TestStringInSlice(t *testing.T) {
	if StringInSlice("a", []string{"a", "b"}) == false {
		t.Errorf("Error")
	}
	if StringInSlice("a", []string{"b", "c"}) == true {
		t.Errorf("Error")
	}
	if StringInSlice("", []string{"a", "b"}) == true {
		t.Errorf("Error")
	}
	if StringInSlice("a bigger test", []string{"a bigger test", "hehe"}) == false {
		t.Errorf("Error")
	}
}
