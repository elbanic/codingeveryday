namespace permutations_ii{

    function permuteUnique(nums: number[]): number[][] {
        const output = new Array()
        backtrack(nums.length, nums, output, 0)

        const seen: Set<number[]> = new Set<number[]>()
        const unique = multiDimensionalUnique(output)
        
        return unique
    };

    function multiDimensionalUnique(arr: number[][]): number[][] {
        var uniques: number[][] = new Array()
        var itemsFound = {}
        for(var i = 0, l = arr.length; i < l; i++) {
            var stringified = JSON.stringify(arr[i])
            if(itemsFound[stringified]) { 
                continue
            }
            uniques.push(arr[i])
            itemsFound[stringified] = true
        }
        return uniques;
    }

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

    const nums = [1,1,2]
    console.log(permuteUnique(nums))

    const nums2 = [1,2,3]
    console.log(permuteUnique(nums2))
}