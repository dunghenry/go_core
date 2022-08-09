package main

import (
	"fmt"
)

func main() {
	var day int
	fmt.Println("Nhap thu so tu 1 den 7")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Chu nhat")
		break
	case 2:
		fmt.Println("Thu hai")
		break
	default:
		fmt.Println("Khong hop le")
	}

	//! Note fallthrough
	i := 45
    switch {
    case i < 10:
        fmt.Println("i is less than 10")
        fallthrough
    case i < 50:
        fmt.Println("i is less than 50")
        fallthrough
    case i < 100:
        fmt.Println("i is less than 100")
    }
	//? result :
	// i is less than 50
	// i is less than 100
}
