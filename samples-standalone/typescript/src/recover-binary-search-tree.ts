
import { TreeNode, createTree } from "./common-tools"

namespace recover_binary_search_tree {


    function recoverTree(root: TreeNode | null): void {

        let stack: TreeNode[] = new Array()
        inorderTraversal(root, stack)

        let first, second: number
        for (let i=0; i<stack.length-1; i++) {
            if (stack[i].val > stack[i+1].val) {
                second = i+1
                if (first == undefined) {
                    first = i
                } else {
                    break
                }
            }
        }
        [stack[first].val, stack[second].val] = [stack[second].val, stack[first].val]
    };

    function swap(node1: TreeNode, node2: TreeNode): void {
        [node1.val, node2.val] = [node2.val, node1.val]
    }

    function inorderTraversal(node: TreeNode, order: TreeNode[]): void {
        if (node == null) {
            return
        }
        inorderTraversal(node.left, order)
        order.push(node)
        inorderTraversal(node.right, order)
    }

    let nums = [3,1,4,null,null,2]
    let root = createTree(nums)
    console.log(recoverTree(root))

}