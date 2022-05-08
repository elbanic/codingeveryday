namespace _132_pattern {

    //brute force
    function find132pattern(nums: number[]): boolean {

        for (let i = 0; i < nums.length - 2; i++) {
            for (let j = i + 1; j < nums.length - 1; j++) {
                for (let z = j + 1; z < nums.length; z++) {
                    if (nums[i] < nums[z] && nums[z] < nums[j]) {
                        return true
                    }
                }
            }
        }
        return false
    };

    //using stack
    function find132pattern2(nums: number[]): boolean {

        if (nums.length < 3) return false

        const stack: number[] = new Array()
        const min: number[] = new Array()

        for (let i of nums) {
            min.push(0)
        }
        min[0] = nums[0]

        for (let i = 1; i < nums.length; i++) {
            min[i] = Math.min(min[i - 1], nums[i])
        }
        for (let j = nums.length - 1; j >= 0; j--) {
            if (nums[j] > min[j]) {
                while (stack.length > 0 && stack[stack.length - 1] <= min[j]) {
                    stack.pop()
                }
                if (stack.length > 0 && stack[stack.length - 1] < nums[j]) {
                    return true
                }
                stack.push(nums[j])
            }
        }

        return false
    };

    const nums = [3, 1, 4, 2]
    console.log(find132pattern2(nums))
}