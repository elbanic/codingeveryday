import { rootCertificates } from "tls"

namespace convert_sorted_array_to_binary_search_tree {


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

    function printTreeNode(root: TreeNode) {
        const queue = new Array(root)
        while (queue.length > 0) {
            const cur = queue.shift()
            console.log(cur.val)
            if (cur.left) {
                queue.push(cur.left)
            }
            if (cur.right) {
                queue.push(cur.right)
            }
        }
    }

    function sortedArrayToBST(nums: number[]): TreeNode | null {
        const root = sortedArrayToBSThelper(nums, 0, nums.length - 1)
        return root
    };

    function sortedArrayToBSThelper(nums: number[], left: number, right: number): TreeNode {
        if (left > right) {
            return null
        }
        let pointer = Math.floor((left + right) / 2)
        if ((left + right) % 2 == 1) {
            pointer++
        }
        const root = new TreeNode(nums[pointer])
        root.left = sortedArrayToBSThelper(nums, left, pointer - 1)
        root.right = sortedArrayToBSThelper(nums, pointer + 1, right)
        return root
    }

    const nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    const root = sortedArrayToBST(nums)
    printTreeNode(root)
}