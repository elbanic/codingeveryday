namespace longest_consecutive_sequence {

    function longestConsecutive(nums: number[]): number {
        if (nums.length == 0) {
            return 0
        }
        if (nums.length == 1) {
            return 1
        }

        nums.sort((a, b) => a - b)
        let p1 = 0, p2 = 1
        let max = p2 - p1

        while (p2 < nums.length) {
            if (nums[p2] - nums[p2 - 1] == 1) {
                max = Math.max(max, nums[p2] - nums[p1] + 1)
            } else if (nums[p2] - nums[p2 - 1] > 1) {
                p1 = p2
            }
            p2++
        }
        return max
    };

    const nums = [0, 1, 2, 0]
    console.log(longestConsecutive(nums))

    const nums2 = [100, 4, 200, 1, 3, 2]
    console.log(longestConsecutive(nums2))
}