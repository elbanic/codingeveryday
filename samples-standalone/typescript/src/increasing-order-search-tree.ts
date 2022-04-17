
namespace increasing_order_search_tree {

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
    
    function increasingBST(root: TreeNode | null): TreeNode | null {
    
        let node = root;
        let stack: TreeNode[] = new Array();
        let res: TreeNode[] = new Array();

        while (stack.length != 0 || node != null) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            let left = stack.pop();
            res.push(left);
            if (left.right != null) {
                node = left.right;
            }
        }
        root = res.shift();
        node = root;
        while (res.length != 0) {
            node.left = null;
            node.right = res.shift();
            node = node.right;
        }
        node.left = null;
        node.right = null;
        return root;
    };

    function createTree(nums: number[]): TreeNode {
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
    }

    let nums = [2,1,4,null, null,3];
    
    let root = createTree(nums);

    console.log(increasingBST(root));
}
