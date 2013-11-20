package assert

import "testing"

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("Expected no error but got %s", err.Error())
	}
}

func NotNil(t *testing.T, val interface{}) {
	if val == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func Equal(t *testing.T, val1, val2 interface{}) {
	if val1 != val2 {
		t.Errorf("Expected %#v to equal %#v", val1, val2)
	}
}
