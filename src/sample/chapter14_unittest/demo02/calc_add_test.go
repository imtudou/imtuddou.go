package demo02

import (
	"testing"
)

// 编写测试用例 测试add 函数是否正确

func TestAdd(t *testing.T) {

	// 调用
	res := add(10)
	if res != 55 {
		t.Errorf("TestAdd fail")
	}
	t.Logf("TestAdd result is :%d", res)
}

func Test_Add(t *testing.T) {
	// 调用
	res := add(3)
	if res != 55 {
		t.Errorf("Test_Add fail")
	}
	t.Logf("Test_Add result is :%d", res)
}
