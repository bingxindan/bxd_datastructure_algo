package main

func main() {

}

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	index := len(postorder) - 1
	rootValue := postorder[index]

	for k, v := range postorder {
		if index == len(postorder)-1 && v > rootValue {
			index = k
		}
		if index != len(postorder)-1 && rootValue > v {
			return false
		}
	}

	return verifyPostorder(postorder[:index]) && verifyPostorder(postorder[index:len(postorder)-1])
}
