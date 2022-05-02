namespace sort_array_by_parity {

    function sortArrayByParity(nums: number[]): number[] {
        const odd: number[] = new Array()
        const even: number[] = new Array()
        while (nums.length > 0) {
            const n = nums.pop()
            if (n % 2 == 0) {
                even.unshift(n)
            } else {
                odd.unshift(n)
            }
        }
        for (let o of odd) {
            even.push(o)
        }
        return even
    };

    function sortArrayByParity2(nums: number[]): number[] {
        const odd: number[] = new Array()
        for (let i=0; i<nums.length; i++) {
            if (nums[i] % 2 == 1) {
                odd.push(nums[i])
                nums.splice(i, 1)
                i--
            }
        }

        for (let o of odd) {
            nums.push(o)
        }
        return nums
    };

    const nums = [3,1,2,4]
    console.log(sortArrayByParity2(nums))
}