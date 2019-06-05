func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    var ls int
    for i := range nums {
        ls = m[nums[i]]
        if ls != 0 {
            return []int{ls - 1, i}
        }
        m[target - nums[i]] = i + 1
    }
    return nil;
}
