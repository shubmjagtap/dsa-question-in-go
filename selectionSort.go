package main

import "fmt"

func swap(a *int, b *int)  {
	temp := *a
	*a = *b
	*b = temp
}

func selectionSort(arrPtr *[]int, size int)  {
	for i := 0; i < size - 1; i++ {
		minIndex := i
		for j := i; j < size  ; j++ {
			if((*arrPtr)[j] < (*arrPtr)[minIndex]){
				minIndex = j
			}
		}
		swap(&(*arrPtr)[minIndex],&(*arrPtr)[i])
	}
}

func main()  {

	var size int
	fmt.Printf("Enter the size of array : ")
	fmt.Scan(&size)

	arr := make([]int,size)

	for i := 0 ; i < size  ; i++ {
		fmt.Printf("Enter the %dth number : ",i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Unsorted array : ",arr)

	selectionSort(&arr,size);

	fmt.Println("Sorted array : ",arr)	

}
