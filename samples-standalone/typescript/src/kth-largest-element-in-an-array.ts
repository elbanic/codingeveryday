import { MaximumHeap } from "./heap";

namespace kth_largest_element_in_an_array {

    function findKthLargest(nums: number[], k: number): number {

        const maxHeap = new MaximumHeap()
        for (let n of nums) {
            maxHeap.insert(n)
        }

        for (let i=1; i<k; i++) {
            maxHeap.pop()
        }
        return maxHeap.pop()
    };

    const nums = [3,2,1,5,6,4], k = 2
    console.log(findKthLargest(nums, k))
}