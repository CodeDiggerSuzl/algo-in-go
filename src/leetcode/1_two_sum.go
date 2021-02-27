package leetcode

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。


链接：https://leetcode-cn.com/problems/two-sum
*/
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for idx, num := range nums {
        if p, ok := m[target-num]; ok {
            return []int{idx, p}
        }
        m[num] = idx
    }
    return nil
}
