package main

import(
	"fmt"
)

//以下的变量都是nil
func theseAreNil(){
	var ptr 	*int
	var arr 	[]int
	var mp		map[string]int
	var cha		chan int
	var fun		func()
	var err		error

	fmt.Println(ptr, arr, mp, cha, fun, err)
}

func sliceBasic(){

	var arr = [3]int{1,2}
	var arr1 = [...]int32{1, 2, 3}
	for i, v := range arr{
		fmt.Println(i, v)
	}

	_ = arr1
}

func main(){
	sliceBasic()
	theseAreNil()
}


