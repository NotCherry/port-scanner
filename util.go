package main

import "strconv"

func Conv_array_to_int(t []string)([]int) {
	var t2 = []int{}

	for _, i := range t {
			j, err := strconv.Atoi(i)
			if err != nil {
					panic(err)
			}
			t2 = append(t2, j)
	}
	return t2
}