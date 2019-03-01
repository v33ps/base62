package base62

import "testing"

func TestEncodeGreaterThanZero(t *testing.T) {
	expected := "19F"
	actual := Encode(4443)
	if actual != expected {
		t.Errorf("TestEncodeGreaterThanZero failed, actual value: %v", actual)
	}
}

func TestEncodeZero(t *testing.T) {
	expected := "0"
	actual := Encode(0)
	if actual != expected {
		t.Errorf("TestEncodeZero failed, actual value: %v", actual)
	}
}

func TestDecode(t *testing.T) {
	expected := 4443
	actual := Decode("19F")
	if actual != expected {
		t.Errorf("TestDecode failed, actual value: %v", actual)
	}
}
