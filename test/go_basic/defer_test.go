package go_basic

import (
	"fmt"
	"testing"
)

func deferFuncParameter() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	aInt = aInt * 3
	return
}

func testReturn1() int {
	var i int

	defer func() {
		i++ //defer里面对i增1
		fmt.Println("test defer, i = ", i)
	}()

	return i
}

// 返回值改为命名返回值
func testReturn2() (i int) {
	//var i int

	defer func() {
		i++
		fmt.Println("test defer, i = ", i)
	}()

	return i
}

func TestDeferReturn2(t *testing.T) {
	ret := testReturn2()
	fmt.Println("test return:", ret)
}

func TestDeferReturn1(t *testing.T) {
	ret := testReturn1()
	fmt.Println("test return:", ret)
}

func TestDeferFuncParameter(t *testing.T) {
	deferFuncParameter()
}

func TestDefer(t *testing.T) {
	testErr()
}

func testErr() {
	var aInt = 1

	defer func() {
		fmt.Println(aInt)
		if aInt == 2 {
			fmt.Println("触发了defer")
		}
	}()

	aInt++
	aInt = aInt * 3

	return
}
