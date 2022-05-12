namespace permutations {

    //backtracking
    function permute(nums: number[]): number[][] {
        const output = new Array()
        backtrack(nums.length, nums, output, 0)
        return output
    };

    function backtrack(n: number, nums: number[], output: number[][], first: number) {
        if (first == n) {
            output.push([...nums])
        }

        for (let i = first; i < n; i++) {
            [nums[first], nums[i]] = [nums[i], nums[first]];
            backtrack(n, nums, output, first + 1);
            [nums[first], nums[i]] = [nums[i], nums[first]];
        }
    }

    const nums = [1, 2, 3]

    permute(nums).forEach(x => console.log(x))
}