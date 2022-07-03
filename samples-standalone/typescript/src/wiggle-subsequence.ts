namespace wiggle_subsequence {

    function wiggleMaxLength(nums: number[]): number {

        if (nums.length < 2) {
            return nums.length
        }

        let prevDiff = nums[1] - nums[0]
        let count = prevDiff != 0 ? 2 : 1

        for (let i=2; i< nums.length; i++) {
            const diff = nums[i] - nums[i - 1]
            if ((diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0)) {
                count++
                prevDiff = diff
            }
        }

        return count
    };

    const nums = [1,2,3,4,5,6,7,8,9]
    console.log(wiggleMaxLength(nums))
}