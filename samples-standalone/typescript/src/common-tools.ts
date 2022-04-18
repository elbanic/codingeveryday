
export class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

export function createTree(nums: number[]): TreeNode {
    if (nums.length == 0) {
        return null;
    }

    let root: TreeNode = new TreeNode(nums[0]);
    let queue: TreeNode[] = new Array(root);

    for (let i = 1; i < nums.length; i += 2) {
        let top: TreeNode = queue.shift();
        if (nums[i] == null) {
            top.left = null;
        } else {
            let left: TreeNode = new TreeNode(nums[i]);
            top.left = left;
            queue.push(left);
        }
        if (i + 1 < nums.length) {
            if (nums[i + 1] == null) {
                top.right = null;
            } else {
                let right: TreeNode = new TreeNode(nums[i + 1]);
                top.right = right;
                queue.push(right);
            }
        }
    }
    return root;
}