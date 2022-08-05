namespace combination_sum_iv {
    function combinationSum4(nums: number[], target: number): number {
        const res = helperCombinationSum4(nums, target, new Map())
        return res
    };

    function helperCombinationSum4(nums: number[], target: number, memo: Map<number, number>): number {
        if (memo.has(target)) {
            return memo.get(target)
        }
        if (target == 0) {
            return 1
        }

        let sum = 0
        for (let n of nums) {
            if (n <= target) {
                const cur = helperCombinationSum4(nums, target - n, memo)
                memo.set(target - n, cur)
                sum += cur
            }
        }
        return sum
    };

    const nums = [3,2,1], target = 35
    console.log(combinationSum4(nums, target))
}