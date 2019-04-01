package unittest

import (
	"testing"
	"fmt"
)

func TestAddUpper(t *testing.T)  {
	res := addUpper(3)
	if res != 3 {
		t.Fatalf("AddUpper(3) 执行错误，期望值=%v 实际值=%v\n", 6, res)
	}
	//如果正确，输出日志
	t.Logf("AddUpper(3) 执行正确...")
}

func TestHello(t *testing.T) {

	fmt.Println("TestHello被调用..")

}