package minitask3

import "fmt"

func SliceInterger(ageSlice []int, inputSlice int) {
	fmt.Println(ageSlice)
	ageSlice = append(ageSlice[:3], append([]int{inputSlice}, ageSlice[3:]...)...)
	fmt.Println(ageSlice)

	for i := 0; i < len(ageSlice); i++ {
		fmt.Println(ageSlice[i])
	}
}
