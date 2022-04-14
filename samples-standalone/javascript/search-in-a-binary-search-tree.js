
/**
 * @param {TreeNode} root
 * @param {number} val
 * @return {TreeNode}
 */

function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
};

var searchBST = function (root, val) {
    let cur = root;
    while (cur != null) {
        if (cur.val == val) {
            return cur;
        } else if (cur.val < val) {
            if (cur.right != null) {
                cur = cur.right;
            } else {
                return null;
            }
        } else {
            if (cur.left != null) {
                cur = cur.left;
            } else {
                return null;
            }
        }
    }
    return null;
};

var createBST = function (nums) {
    if (nums.length == 0) {
        return null
    }

    let node = new TreeNode(nums[0]);
    let queue = new Array(node);

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
    return node
};

let nums = [62,2,93,null,30,null,null,15,null,null,null];
let val = 15;
let root = createBST(nums);

console.log(searchBST(root, val));