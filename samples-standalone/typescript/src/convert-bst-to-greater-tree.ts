import { TreeNode, createTree } from "./common-tools";

namespace convert_bst_to_greater_tree {

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
    
    const nums: number[] = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8];
    const root: TreeNode = createTree(nums);
    
    console.log(convertBST(root));
    let a: number = 0;
}
