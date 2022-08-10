package main

import (
	"fmt"
	"demo/controller/demo"
	"demo/controller/lib"
)

func main() {
	fmt.Println("Hi")
	fmt.Println(demo.PI)
	var a int
	var b int
	fmt.Print("Nhap so a : ")
	fmt.Scan(&a)
	fmt.Print("Nhap so b : ")
	fmt.Scan(&b)
	fmt.Printf("Tong so %v va %v la: %v", a, b, demo.Tong(a, b))
	fmt.Println()
	fmt.Printf("Hieu so %v va %v la: %v", a, b, demo.Hieu(a, b))
	fmt.Println()
	fmt.Printf("Tich so %v va %v la: %v", a, b, demo.Nhan(a, b))
	fmt.Println()
	if(a > b){
		fmt.Printf("Random number: %v", lib.Random(a, b))
		fmt.Println()
	}else{
		fmt.Printf("Random number: %v", lib.Random(b, a))
		fmt.Println()
	}

	fmt.Printf("Pow number: %v", lib.Pow(a, b))
}