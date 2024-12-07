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

func TestSlice1(t *testing.T) {
	var array [10]int

	var slice = array[5:6]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])

}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func TestSlice2(t *testing.T) {

	var slice []int
	slice = append(slice, 1, 2, 3)

	fmt.Println(cap(slice))
	fmt.Println(slice)
	newSlice := AddElement(slice, 4)
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))

	fmt.Println(slice)

	fmt.Println(&slice[0] == &newSlice[0])

}

func TestSlice3(t *testing.T) {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
