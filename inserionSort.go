package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func insertionSort(arrPtr *[]int, size int) {
	for i := 0 ; i <= size - 1; i++ {
		j := i
		for j > 0 && (*arrPtr)[j-1] > (*arrPtr)[j] {
			swap(&((*arrPtr)[j-1]),&((*arrPtr)[j]))
			j--;
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

	insertionSort(&arr, size)

	fmt.Println("Sorted array : ", arr)

}
