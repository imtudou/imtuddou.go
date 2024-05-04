package demo02

import "testing"

func TestSub(t *testing.T) {
	res := sub(10, 1)
	if res != 9 {
		t.Errorf("TestSub is error")
	}
	t.Logf("TestSub result is :%d", res)
}
