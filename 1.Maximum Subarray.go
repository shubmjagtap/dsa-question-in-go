func maxSubArray(nums []int) int {
    
    var maxSum int
    var currentSum int
    maxSum = -10000
    currentSum = 0

    for i:= 0 ; i < len(nums) ; i++ {
        if nums[i] >= currentSum + nums[i]{
            currentSum = nums[i]
        } else {
            currentSum = currentSum + nums[i]
        }
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }    
    return maxSum

}
