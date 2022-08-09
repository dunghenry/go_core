package main

import "fmt"

func main() {
	//Array
	nums := []int{1, 2, 3}
	for index, item := range nums {
		fmt.Println(index, item)
	}

	fmt.Printf("%T", nums)
	fmt.Println()

	for index, item := range nums {
		if item == 3 {
			fmt.Println("value :", item)
			fmt.Println("index :", index)
		}
	}

	//or
	
	for _, item := range nums {
		if item == 3 {
			fmt.Println("value :", item)
		}
	}

	//Map (object javascript)
	mapstr := map[string]string{"2": "Mango", "1": "Apple"}
	for i, v := range mapstr {
		//sort big > small
		fmt.Printf("%s -> %s\n", i, v)
	}

	for k := range mapstr {
		//sort big > small
		fmt.Println("key:", k)
	}

	//String
	for i, c := range "Hi" {
		//convert number ASCII
		fmt.Println(i, c)
		//H : 72
		// i : 105
	}
}
