package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	b := [4]int{10, 20, 30, 40}
	reverse(&a) // Меняет порядок элементов на обратный
	reverse0(&b)
	fmt.Println(a)
	fmt.Println(b)
}

func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr) - 1 - index] = value
	}
}

func reverse0(arr *[4]int) {
	for i, j := 0, len(*arr) - 1; i < j; i, j = i + 1, j - 1 {
    	arr[i], arr[j] = arr[j], arr[i]
	}
}