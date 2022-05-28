namespace missing_number {
    function missingNumber(nums: number[]): number {

        nums.sort()
        if (nums[nums.length - 1] != nums.length) {
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
        return nums[mid] - 1
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

    const nums = [45,35,38,13,12,23,48,15,44,21,43,26,6,37,1,19,22,3,11,32,4,16,28,49,29,36,33,8,9,39,46,17,41,7,2,5,27,20,40,34,30,25,47,0,31,42,24,10,14]
    console.log(missingNumber2(nums))
}