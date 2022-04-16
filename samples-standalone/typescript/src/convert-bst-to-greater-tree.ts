
class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}


function convertBST(root: TreeNode | null): TreeNode | null {
    let sum: number = 0;
    let node: TreeNode = root;
    let stack: TreeNode[] = new Array();
    while (stack.length != 0 || node != null) {
        while (node != null) {
            stack.push(node);
            node = node.right;
        }
        node = stack.pop();
        sum += node.val;
        node.val = sum;

        node = node.left;
    }
    return root;
};

function createBST(nums: number[]): TreeNode {

    if (nums.length == 0 ) {
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
            if (nums[i+1] == null) {
                top.right = null;
            } else {
                let right: TreeNode = new TreeNode(nums[i + 1]);
                top.right = right;
                queue.push(right);
            }
        }
    }
    return root;
};


const nums: number[] = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8];
const root: TreeNode = createBST(nums);

console.log(convertBST(root));
let abc: number = 0;