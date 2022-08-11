namespace validate_binary_search_tree {

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

    function createTreeNode(nums: number[]): TreeNode {
        const root = new TreeNode(nums[0])
        const queue = new Array(root)

        for (let i = 1; i < nums.length; i += 2) {
            const top = queue.shift()
            if (nums[i] != null) {
                top.left = new TreeNode(nums[i])
                queue.push(top.left)
            }
            if (i + 1 < nums.length && nums[i + 1] != null) {
                top.right = new TreeNode(nums[i + 1])
                queue.push(top.right)
            }
        }
        return root
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

    function isValidBST(root: TreeNode | null): boolean {
        return isValidBSTHelper(root, null, null)
    }

    function isValidBSTHelper(root: TreeNode, low: number, high: number): boolean {
        if (root == null) {
            return true
        }
        if ((low != null && root.val <= low) || (high != null && root.val >= high)) {
            return false
        }
        return isValidBSTHelper(root.right, root.val, high) && isValidBSTHelper(root.left, low, root.val)
    }

    const root = [5,4,6,null,null,3,7]
    const rootTree = createTreeNode(root)

    console.log(isValidBST(rootTree))
}