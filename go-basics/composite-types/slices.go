package main

import (
	"fmt"
	"slices"
)

func main() {

	// SLICE ( similar to vectors in cpp)
	// slices are mostly used when we need to hold sequence of values, over arrays in go
	// slices are not comparable like arrays, e.g. slicearr1 == slicear2 cannot be done
	// slices can be compared only to 'nil' value ,e..g slicearr1 == nil

	var sliceArr = []int{10, 20, 30, 40, 50}

	// 2d slices
	var twodSlice [][]int

	// Zero value for slices is 'nil'
	//
	var sliceNilArr []int // intialized to 'nil' value

	// **** IN-built slices package functions
	//
	x := []int{1, 2, 3, 5, 6, 7}
	y := []int{1, 2, 3, 5, 6, 7}
	z := []int{0, 2, 3, 5, 6, 2}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	fmt.Println(slices.Equal(y, z))

	fmt.Println(sliceArr, twodSlice, sliceNilArr)

	fmt.Println(len(x))

	x = append(x, 20, 30)
	fmt.Println(x)

	// appending 2 slices togehter using ...  operator
	x = append(x, y...)
	fmt.Println(x)

	///Capacity
	fmt.Println(cap(x))

	// make() to create slices

	makeSliceStyle1 := make([]int, 5)     // (type,lenght)
	makeSliceStyle2 := make([]int, 5, 10) // (type,length,cap)

	makeSliceStyle3 := make([]int, 0)

	fmt.Println(makeSliceStyle1, cap(makeSliceStyle1), makeSliceStyle2, cap(makeSliceStyle2), makeSliceStyle3)

	// emptying a slice using clear()
	// sets elements in the slice values to 0
	namesSlice := []string{"saish", "pawar", "saish"}
	fmt.Println(namesSlice)
	clear(namesSlice) // in this case, "" for string elements as that is their zero value in go

	fmt.Println(namesSlice, len(namesSlice), cap(namesSlice))

	//slicing slices
	// /// Leaving starting offset means start from 0 index, leaving ending offset means till last index
	// fist offset inclusive, second offset excluded

	a := []string{"a", "b", "c", "d"}
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	e := a[:]
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// Copying slices:
	/// copy(destinationSlice, sourceSlice)
	sliceCopy1 := []int{1, 2, 3, 4}
	sliceCopy2 := make([]int, 2)
	sliceCopy4 := make([]int, 2)

	sliceCopy3 := copy(sliceCopy2, sliceCopy1) // copies min(len(sliceCopy2),len(sliceCopy1))

	copy(sliceCopy4, sliceCopy1[2:])

	fmt.Println("sliceCopy1 : ", sliceCopy1)
	fmt.Println("sliceCopy2 : ", sliceCopy2)
	fmt.Println("sliceCopy3 : ", sliceCopy3)
	fmt.Println("sliceCopy4 (slices of slice : ", sliceCopy4)

	// Convertign arrays <--> slices
	//
	// Array to slice
	exArr := [5]int{1, 2, 3, 4, 5}
	exArrSliec := exArr[:]
	fmt.Printf("%T\n", exArrSliec)

	// Slice to array
	exArr = [5]int(exArrSliec)
	fmt.Printf("%T\n", exArr)

}
