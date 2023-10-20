package sorts

import "fmt"

func bubblesort() {
	array := [8]int{1,4,5,2,6,3,8,1}
	fmt.Println(array)
	for i := 0;i <len(array);i++{
		for j :=0;j<len(array);j++{
			if array[i]<=array[j]{
				aux := array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	fmt.Println(array)
}