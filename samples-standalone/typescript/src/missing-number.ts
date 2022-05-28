namespace missing_number {

    //binary search
    function missingNumber(nums: number[]): number {

        nums.sort((a,b) => a - b)
        if (nums[nums.length-1] != nums.length) {
            return nums.length
        }

        let left = 0, right = nums.length - 1
        let mid = left
        while (left <= right) {
            mid = Math.floor(left + (right - left) / 2)
            if (nums[mid] <= mid) {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        if (left == mid) {
            return nums[mid] - 1
        } else {
            return nums[mid] + 1
        }
    };


    function missingNumber2(nums: number[]): number {

        nums.sort((a,b) => a - b)
        for (let i = 0; i<nums.length; i++) {
            if (nums[i] != i) {
                return i
            }
        }
        return nums.length
    }

    const nums = [0,2,3]
    console.log(missingNumber(nums))
}