package unittest

import (
	"testing"
)

func TestStore(t *testing.T)  {
	monster := &Monster{
		Name : "孙悟空",
		Age : 6000,
		Skill : "筋斗云",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster store() error,希望为=%v,实际是=%v",true, res)
	}
	t.Logf("monster store() 测试成功")
}

func TestReStore(t *testing.T)  {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster TestReStore() error,希望为=%v,实际是=%v",true, res)
	}
	if monster.Name != "孙悟空" {
		t.Fatalf("monster TestReStore() error,希望为=%v,实际是=%v","孙悟空", monster.Name)
	}
	t.Logf("monster restore() 测试成功")
}