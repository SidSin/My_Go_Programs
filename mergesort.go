package pftp

import (
	"fmt"
	"strconv"
)

const globaldebuglevel = 0

func writedebugarray(debuglevel int, arr []int) {
	if debuglevel > 0 {
		fmt.Println(arr)
	}
}

func writedebugmessage(debuglevel int, message ...string) {
	if debuglevel > 0 {
		fmt.Println(message)
	}
}

//this func assumes that the arrays are sorted
func merge(arr1, arr2 []int) []int {

	var resultarr []int
	var dbglvl int = globaldebuglevel

	writedebugmessage(dbglvl, " ***** Arr1 length = ", strconv.Itoa(len(arr1)), "  Arr2 Length = ", strconv.Itoa(len(arr2)))

	if len(arr1) == 0 {
		resultarr = arr2
	} else if len(arr2) == 0 {
		resultarr = arr1
	} else {
		resultarr = buildresultarr(arr1, arr2)
	}
	writedebugmessage(dbglvl, "Ending merge")
	writedebugarray(dbglvl, resultarr)

	return resultarr

}

func buildresultarr(firstarray []int, secondarray []int) []int {

	var i, j, k int
	totalsize := len(firstarray) + len(secondarray)
	var resultarr = make([]int, totalsize)
	var dbglvl int = globaldebuglevel

	copyfirstarray := false
	copysecondarray := false

	writedebugmessage(dbglvl, " ***** firstarray length = ", strconv.Itoa(len(firstarray)))

	for i = 0; i < totalsize; {

		writedebugmessage(dbglvl, " ***** i = ", strconv.Itoa(i), " , j = ", strconv.Itoa(j), " , k = ", strconv.Itoa(k))

		if firstarray[i] < secondarray[j] {

			writedebugmessage(dbglvl, " ***** (1) small = ", strconv.Itoa(firstarray[i]))
			resultarr[k] = firstarray[i]
			writedebugmessage(dbglvl, " ***** (1) resultarr[k] = ", strconv.Itoa(resultarr[k]))
			i++
			k++

			if i >= len(firstarray) {
				copysecondarray = true
				break
			}

		} else {
			writedebugmessage(dbglvl, " ***** (2) small = ", strconv.Itoa(secondarray[j]))
			resultarr[k] = secondarray[j]
			writedebugmessage(dbglvl, " ***** (2) resultarr[k] = ", strconv.Itoa(resultarr[k]))
			j++
			k++

			if j >= len(secondarray) {
				copyfirstarray = true
				break
			}
		}
	}

	writedebugmessage(dbglvl, " array looping ends", " i = ", strconv.Itoa(i), " , j = ", strconv.Itoa(j), " , k = ", strconv.Itoa(k))

	if copyfirstarray {
		copyremainingarrelements(i, firstarray, resultarr, k)
	}

	if copysecondarray {
		copyremainingarrelements(j, secondarray, resultarr, k)
	}

	writedebugmessage(dbglvl, " ending buildresultarr")
	writedebugarray(dbglvl, resultarr)

	return resultarr
}

func copyremainingarrelements(a int, remarray []int, resultarr []int, k int) {

	var dbglvl int = globaldebuglevel

	for a < len(remarray) {
		resultarr[k] = remarray[a]
		writedebugmessage(dbglvl, " k = ", strconv.Itoa(k))
		k++
		a++
	}
}

func MergeSort(unsortedarray []int) []int {

	arr1 := make([]int, len(unsortedarray))
	arr2 := make([]int, len(unsortedarray))

	var dbglvl int = 1 //globaldebuglevel

	if len(unsortedarray) == 1 {

		return unsortedarray

	} else {

		mid := len(unsortedarray) / 2

		arr1 = unsortedarray[0:mid]
		arr2 = unsortedarray[mid:]

		arr1 = MergeSort(arr1)
		arr2 = MergeSort(arr2)
	}

	writedebugmessage(dbglvl, " calling merge with arr1 = ")
	writedebugarray(dbglvl, arr1)
	writedebugmessage(dbglvl, " arr2 = ")
	writedebugarray(dbglvl, arr2)

	arr3 := merge(arr1, arr2)

	writedebugmessage(dbglvl, " Final sorted array is ")
	writedebugarray(dbglvl, arr3)

	return arr3

}
