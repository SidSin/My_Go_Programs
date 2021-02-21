package udemyproblems

import "fmt"

//BubbleSort sorts the given array using bubble sort algorithm
func BubbleSort(array []int) {

	bound := len(array)
	t := bound

	for t > 0 {

		fmt.Println("Start Loop")
		// this indicates no swap has happened so far
		// t = 0 in program means swap happened when
		// very first element is greater than second
		// and rest of the array is sorted
		t = -1

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

		if t == -1 {
			break
		} else {
			bound = t + 1
		}

		fmt.Println("bound = ", bound)
		fmt.Println("**** Intermidiate array  = ", array)
	}

} // func ends
