package sort 

import "racer/model"

func Sort(unsorted *[]model.Row) {
	if len(*unsorted) < 2 {
		return
	}
	quicksort(unsorted, 0, len(*unsorted)-1)
}

func quicksort(unsorted *[]model.Row, p int, r int) {
	if p < r {
		q := partition(unsorted, p, r)
		quicksort(unsorted, p, q-1)
		quicksort(unsorted, q+1, r)
	}
}

func partition(arr *[]model.Row, p int, r int) int {
	x := &(*arr)[r]
	i := p - 1
	j := p
	
	for j <= r-1 {
		if (*arr)[j].DriverTime.Best <= (*x).DriverTime.Best {
			i++
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[i]
			(*arr)[i] = temp 
		}
		j++
	}

	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[j]
	(*arr)[j] = temp
	
	return i+1
}
