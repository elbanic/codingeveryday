/**
 * @param {number[]} height
 * @return {number}
 */
//brute force
var maxArea = function (height) {

    let max = 0;
    for (let i = 0; i < height.length - 1; i++) {
        for (let j = i + 1; j < height.length; j++) {

            const area = Math.min(height[i], height[j]) * (j - i)
            if (max < area) {
                max = area;
            }
        }
    }
    return max;
};

// two pointer 
var maxArea2 = function (height) {

    let left = 0;
    let right = height.length - 1;

    let max = 0;
    while (left < right) {
        const area = Math.min(height[left], height[right]) * (right - left)
        if (max < area) {
            max = area;
        }
        if (height[left] < height[right]) {
            left++;
        } else {
            right--;
        }
    }
    return max;
};

height = [1, 8, 6, 2, 5, 4, 8, 3, 7];
console.log(maxArea2(height))