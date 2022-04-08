/**
 * @param {number} k
 * @param {number[]} nums
 */
var KthLargest = function (k, nums) {
    nums.sort(function (a, b) { return a - b });
    this.k = k
    this.nums = nums;
};

/** 
 * @param {number} val
 * @return {number}
 */
KthLargest.prototype.add = function (val) {

    let left = 0;
    let right = this.nums.length - 1;

    while (left <= right) {
        const mid = Math.floor(left + (right - left) / 2);

        if (this.nums[mid] == val) {
            left = mid;
            break;
        } else if (this.nums[mid] < val) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    this.nums.splice(left, 0, val);
    return this.nums[this.nums.length - this.k];
};

/** 
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */

const nums = [-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1];
var obj = new KthLargest(7, nums)
console.log(obj.add(3))
console.log(obj.add(2))
console.log(obj.add(3))
console.log(obj.add(1))
console.log(obj.add(2))
console.log(obj.add(4))
console.log(obj.add(5))
console.log(obj.add(5))
console.log(obj.add(6))
console.log(obj.add(7))
console.log(obj.add(7))
console.log(obj.add(8))
console.log(obj.add(2))
console.log(obj.add(3))
console.log(obj.add(1))
console.log(obj.add(1))
console.log(obj.add(1))
console.log(obj.add(10))
console.log(obj.add(11))
console.log(obj.add(5))
console.log(obj.add(6))
console.log(obj.add(2))
console.log(obj.add(4))
console.log(obj.add(7))
console.log(obj.add(8))
console.log(obj.add(5))
console.log(obj.add(6))
