package go_basic

import (
	"fmt"
	"testing"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func TestSelect(t *testing.T) {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go addNumberToChan(chan1)
	go addNumberToChan(chan2)
	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
			//default:
			//	fmt.Printf("No element in chan1 and chan2.\n")
			//	time.Sleep(1 * time.Second)
		}
	}
}

func TestSelect1(t *testing.T) {
	testSelect1()
}

func testSelect1() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}

	fmt.Println("main exit.")
}

func TestSelect2(t *testing.T) {
	testSelect2()
}

func testSelect2() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	writeFlag := false
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if writeFlag {
				chan2 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
}

func TestSelect3(t *testing.T) {
	testSelect3()
}

func testSelect3() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")

}

func TestSelect4(t *testing.T) {
	select {}
}

// 对于空的select语句，程序会被阻塞，准确的说是当前协程被阻塞，同时Golang自带死锁检测机制，当 发现当前协程再也没有机会被唤醒时，则会panic。所以上述程序会panic。
