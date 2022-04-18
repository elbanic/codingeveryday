import { TreeNode, createTree } from "./common-tools";

namespace increasing_order_search_tree {
    
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

    let nums = [2,1,4,null, null,3];
    let root = createTree(nums);

    console.log(increasingBST(root));
}
