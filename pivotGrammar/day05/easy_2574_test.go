package day05

import (
	"fmt"
	"testing"
)

/*
*

给你一个下标从 0 开始的整数数组 nums ，请你找出一个下标从 0 开始的整数数组 answer ，其中：

answer.length == nums.length
answer[i] = |leftSum[i] - rightSum[i]|
其中：

leftSum[i] 是数组 nums 中下标 i 左侧元素之和。如果不存在对应的元素，leftSum[i] = 0 。
rightSum[i] 是数组 nums 中下标 i 右侧元素之和。如果不存在对应的元素，rightSum[i] = 0 。
返回数组 answer 。

输入：nums = [10,4,8,3]
输出：[15,1,11,22]
解释：数组 leftSum 为 [0,10,14,22] 且数组 rightSum 为 [15,11,3,0] 。
数组 answer 为 [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22] 。
*/

func Sum(nums []int) int {
	var tmp int
	for _, num := range nums {
		tmp += num
	}
	return tmp
}

func Abs(value int) int {
	if value < 0 {
		return -1 * value
	} else {
		return value
	}

}

func leftRigthDifference(nums []int) []int {
	/*
	   *

	   		leftSum[i] 是数组 nums 中下标 i 左侧元素之和。如果不存在对应的元素，leftSum[i] = 0 。
	   	  rightSum[i] 是数组 nums 中下标 i 右侧元素之和。如果不存在对应的元素，rightSum[i] = 0 。
	*/
	var answers []int
	var left int
	var right int = Sum(nums)
	i := len(nums)
	for index := range nums {

		rightAdd := nums[index]
		if index != 0 {
			left += nums[index-1]
		} else {
			left = 0
		}

		if index == i-1 {
			right = 0
		} else {
			right -= rightAdd
		}
		answers = append(answers, Abs(left-right))

	}
	fmt.Println(answers)
	return answers

}

func TestName(t *testing.T) {
	var nums = []int{10, 4, 8, 3}
	leftRigthDifference(nums)
}
