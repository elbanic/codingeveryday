

function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @param {number} low
 * @param {number} high
 * @return {TreeNode}
 */
var trimBST = function (root, low, high) {
    if (root == null) return root;
    if (root.val > high) return trimBST(root.left, low, high);
    if (root.val < low) return trimBST(root.right, low, high);

    root.left = trimBST(root.left, low, high);
    root.right = trimBST(root.right, low, high);
    return root;
};

var createNode = function (nums) {
    if (nums.length == 0) {
        return null;
    }
    let root = new TreeNode(nums[0]);
    let queue = new Array(root);

    for (let i = 1; i < nums.length; i += 2) {
        let top = queue.shift();
        if (nums[i] == null) {
            top.left = null;
        } else {
            let left = new TreeNode(nums[i]);
            top.left = left;
            queue.push(left);
        }
        if (i + 1 < nums.length) {
            if (nums[i+1] == null) {
                top.right = null;
            } else {
                let right = new TreeNode(nums[i + 1]);
                top.right = right;
                queue.push(right);
            }
        }
    }
    return root;
}

let num = [1, 0, 2];
let low = 1, high = 2;

let root = createNode(num);
console.log(trimBST(root, low, high));


