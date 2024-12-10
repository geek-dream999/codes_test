package go_basic

import (
	"fmt"
	"testing"
)

func TestDeferReturn2(t *testing.T) {
	ret := testReturn2()
	fmt.Println("test return:", ret)
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

/*
test defer, i =  1
test return: 1
*/

func TestDeferReturn1(t *testing.T) {
	ret := testReturn1()
	fmt.Println("test return:", ret)
}

func testReturn1() int {
	var i int

	defer func() {
		i++ //defer里面对i增1
		fmt.Println("test defer, i = ", i)
	}()
	fmt.Println(i)
	i = i + 2

	return i
}

/*
0
test defer, i =  3
test return: 2
*/

func TestDeferFuncParameter(t *testing.T) {
	deferFuncParameter()
}
func deferFuncParameter() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	aInt = aInt * 3
	return
}

/*
1
*/

func TestDefer(t *testing.T) {
	testErr()
}

func testErr() {
	var aInt = 1

	defer func() {
		fmt.Println(aInt)
	}()

	aInt++
	aInt = aInt * 3

	return
}

/*
6
*/

func TestDeferFuncReturn(t *testing.T) {
	ret := deferFuncReturn()
	fmt.Println("test return:", ret)
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func TestDeferReturn(t *testing.T) {
	ret := testReturn()
	fmt.Println("test return:", ret)
}

func testReturn() int {
	a := 1

	defer func() {
		a = a + 1 // defer 中修改 a 的值
	}()

	a = a * 2 // 修改 a 的值

	return a // 返回 a，实际返回的值是 2 还是 3？
}
