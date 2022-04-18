import { TreeNode, createTree } from "./common-tools";

namespace kth_smallest_element_in_a_bst {

    function kthSmallest(root: TreeNode | null, k: number): number {

        let stack: TreeNode[] = new Array()
        let res: TreeNode[] = new Array()

        let node = root
        while (res.length != 0 || node != null) {
            while (node != null) {
                res.push(node)
                node = node.left
            }
            let left = res.pop()
            stack.push(left)

            if (stack.length == k) {
                break
            }
            if (left.right != null) {
                node = left.right
            }
        }
        return stack.pop().val
    };

    let nums = [3, 1, 4, null, 2], k = 1
    let root = createTree(nums)

    console.log(kthSmallest(root, k))
}