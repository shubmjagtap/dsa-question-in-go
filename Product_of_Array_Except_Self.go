func productExceptSelf(nums []int) []int {

    n := len(nums)

    // init prefix array
    prefix := make([]int, n)
    prefix[0] = nums[0];

    // init postfix array
    postfix := make([]int, n)
    postfix[n-1] = nums[n-1];

    // init result array
    result := make([]int, n)
    result[0] = nums[0];
    result[n-1] = nums[n-1];

    // make prefix array
    for i := 1 ; i < n ; i++ {
        prefix[i] = prefix[i-1] * nums[i];
    } 

    // make postfix array
    for i := n-2 ; i >= 0 ; i-- {
        postfix[i] = postfix[i+1] * nums[i];
    } 

    // make result array
    for i := 0 ; i < n ; i++ {

        var left int;
        if i == 0 {
            left = 1;
        }else{
            left = prefix[i-1]
        }

        var right int;
        if i == n-1 {
            right = 1;
        }else{
            right = postfix[i+1]
        }

        result[i] = left*right;
    } 

    // return result array
    return result
}
