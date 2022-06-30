namespace minimum_moves_to_equal_array_elements_ii {

    function minMoves2(nums: number[]): number {

        nums.sort((a, b) => a - b)

        const med = nums.length % 2 == 1 ? nums[Math.floor(nums.length / 2)] :
            Math.floor((nums[nums.length / 2 - 1] + nums[nums.length / 2]) / 2)

        let sum = 0
        for (let n of nums) {
            sum += Math.abs(n - med)
        }
        return sum
    };

    const nums = [1, 2]
    console.log(minMoves2(nums))
}