/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var nextPermutation = function (nums) {
    let i = nums.length - 2;
    while (i >= 0 && nums[i + 1] <= nums[i]) {
        i--;
    }
    if (i >= 0) {
        let j = nums.length - 1;
        while (nums[j] <= nums[i]) {
            j--;
        }
        swap(nums, i, j);
    }
    reverse(nums, i + 1);
};

var reverse = function(nums, start) {
    let i = start;
    let j = nums.length - 1;

    while (i < j ) {
        swap(nums, i, j);
        i++;
        j--;
    }
};


var swap = function(nums, i, j) {
    const temp = nums[i];
    nums[i] = nums[j];
    nums[j] = temp;
};

const nums = [1,2,3];
nextPermutation(nums)
console.log(nums)