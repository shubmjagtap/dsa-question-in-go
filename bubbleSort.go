package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func bubbleSort(arrPtr *[]int, size int) {
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if (*arrPtr)[j] > (*arrPtr)[j+1] {
				swap(&((*arrPtr)[j]), &((*arrPtr)[j+1]))
			}
		}
	}
}

func main() {

	var size int
	fmt.Printf("Enter the size of array : ")
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth number : ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Unsorted array : ", arr)

	bubbleSort(&arr, size)

	fmt.Println("Sorted array : ", arr)

}
