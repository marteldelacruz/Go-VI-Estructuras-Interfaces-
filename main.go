package main

import "fmt"

func main() {
	var opt string

	for opt != "0" {
		fmt.Println("")
		fmt.Println("1-. Add image")
		fmt.Println("2-. Add audio")
		fmt.Println("3-. Add video")
		fmt.Println("4-. Show content")
		fmt.Println("0-. Exit")
		fmt.Print("\nSelect an option:")
		fmt.Scanln(&opt)

		switch opt {
		case "1":
			break
		case "2":
			break
		case "3":
			break
		case "4":
			break
		}
	}
}
