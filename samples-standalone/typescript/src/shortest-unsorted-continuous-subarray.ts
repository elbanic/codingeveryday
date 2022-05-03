import { min } from "ramda"

namespace shortest_unsorted_continuous_subarray {
    function findUnsortedSubarray(nums: number[]): number {

        const copyNums = [...nums]
        nums.sort(function(a: number, b:number){
            return a - b
        })

        let len: number = 0, start: number = -1, end: number = -1
        for (let i = 0; i < copyNums.length; i++) {
            if (copyNums[i] != nums[i]) {
                start = i
                break
            }
        }
        for (let i = copyNums.length - 1; i >= 0; i--) {
            if (copyNums[i] != nums[i]) {
                end = i
                break
            }
        }
        if (end == start) {
            return 0
        }
        return end - start + 1
    };

    const nums = [1,2,3,4]
    console.log(findUnsortedSubarray(nums))
}