package classic

func removeElement(nums []int, val int) int {
	length := len(nums)
	cur := 0
	for times := 0; times < length; times++ {
		if nums[cur] != val {
			cur++
			continue
		}

		temp := nums[cur+1:]
		temp = append(temp, nums[cur])
		nums = append(nums[:cur], temp...)
	}
	return cur
}

func removeElement2(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
