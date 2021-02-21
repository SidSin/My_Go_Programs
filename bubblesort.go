package udemyproblems

import "fmt"

//BubbleSort sorts the given array
func BubbleSort(array []int) {

	bound := len(array)
	t := bound

	for t > 0 {

		fmt.Println("Start Loop")
		t = 0

		for j := 0; j < (bound - 1); j++ {

			fmt.Println("Comparing ", array[j], " and ", array[j+1])

			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				t = j
			}
		}

		fmt.Println("t = ", t)

		if t == 0 {
			break
		} else {
			bound = t + 1
		}

		fmt.Println("bound = ", bound)
		fmt.Println("**** Intermidiate array  = ", array)
	}

} // func ends
