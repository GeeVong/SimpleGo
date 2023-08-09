package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	// 5, 9, 3, 1, 7, 6

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

			showArrInfo(arr)
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	showArrInfo(arr)
	return i + 1
}

func showArrInfo(arr []int) {
	fmt.Println("arr info:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func main() {
	arr := []int{5, 9, 3, 1, 7, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
