namespace path_sum {

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

    function hasPathSum(root: TreeNode | null, targetSum: number): boolean {

        if (!root) {
            return false
        }
        if (!root.left && !root.right) {
            return root.val == targetSum
        }
        if (root.left) {
            root.left.val += root.val
        }
        if (root.right) {
            root.right.val += root.val
        }
        return hasPathSum(root.left, targetSum) || hasPathSum(root.right, targetSum)
    };

    function createTreeNode(nums: number[]): TreeNode {

        const root = new TreeNode(nums[0])
        const queue: TreeNode[] = new Array()
        queue.push(root)

        for (let i = 1; i < nums.length; i += 2) {
            const top = queue.shift()
            if (i < nums.length) {
                top.left = new TreeNode(nums[i])
                queue.push(top.left)
            }
            if (i + 1 < nums.length) {
                top.right = new TreeNode(nums[i + 1])
                queue.push(top.right)
            }
        }
        return root
    }

    const root = [], targetSum = 22
    const rootTree = createTreeNode(root)
    console.log(hasPathSum(rootTree, targetSum))
}