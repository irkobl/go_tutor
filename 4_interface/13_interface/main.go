package main

import "fmt"

func ValueInfo(obj interface{}) {
	switch val := obj.(type) {
	case string:
		fmt.Printf("Длина строки: %d\n", len(val))
	case []int:
		fmt.Println("Длина строки:", cap(val))
	default:
		fmt.Printf("Tip %T\n", val)
	}

}

func main() {
	var obj interface{}
	obj = 1
	fmt.Println(obj)
	obj = "str"
	fmt.Println(obj)

	var slc []interface{}
	slc = append(slc, 1)
	slc = append(slc, "str")
	slc = append(slc, []int{2, 4})
	slc = append(slc, 3.14)
	fmt.Println(slc)

	mp := map[int]interface{}{
		1: "go",
		2: 1234,
		3: false,
	}

	mp[1] = "gogo"

	value, ok := mp[2].(int)
	if !ok {
		fmt.Println("error")
	}
	mp[2] = value + 100

	ValueInfo("str")
	ValueInfo([]int{1, 2})
	ValueInfo(true)
}
