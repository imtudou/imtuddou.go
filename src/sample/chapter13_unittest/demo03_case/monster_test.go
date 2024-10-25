package demo03_case

import "testing"

/*
单元测试综合案例：demo03_case
1) .编写一个Monster结构体,字段Name, Age, Skill
2) 给Monster绑定方法Store,可以将一个Monster变量(对象),序列化后保存到文件中
3) 给Monster绑定方法ReStore,可以将一个序列化的Monster,从文件中读4)।取,并反序列化为Monster对象,检查反序列化,名字正确编程测试用例文件store_test.go,编写测试用例函数TestStore和TestRestore进行测试。
4)
*/

func TestMonster_MarshalAndWriteFile(t *testing.T) {
	monster := &Monster{
		Name:   "红孩儿",
		Age:    18,
		Gender: "女",
	}

	ret := monster.MarshalAndWriteFile()
	if ret != nil {
		t.Errorf("TestMonster_MarshalAndWriteFile:%v", ret.Error())
	}
	t.Logf("TestMonster_MarshalAndWriteFile:success ")
}

func TestMonster_UnmarshalAndReadFile(t *testing.T) {
	monster := &Monster{}
	ret := monster.UnmarshalAndReadFile()
	if ret != nil {
		t.Errorf("TestMonster_UnmarshalAndReadFile:%v", ret.Error())
	}
	if monster.Name != "红孩儿" {
		t.Errorf("TestMonster_UnmarshalAndReadFile:%v", monster.Name)
	}
	t.Logf("TestMonster_UnmarshalAndReadFile: success")

}
