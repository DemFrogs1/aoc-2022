package lib

func DeleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
