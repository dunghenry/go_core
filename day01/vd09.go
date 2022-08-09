package main

import ("fmt")

func main() {
	var n int
	loop:
		fmt.Println("Nhap n : ")
		fmt.Scan(&n)
	if n <= 0{
		goto loop
	}else{
		fmt.Println("Successfully")
		fmt.Println(n)  
	}
}