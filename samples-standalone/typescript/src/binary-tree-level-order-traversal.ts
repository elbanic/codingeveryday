namespace binary_tree_level_order_traversal {

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

    function createTree(nums: number[]): TreeNode {
        if (nums.length == 0) {
            return null
        }
        const root = new TreeNode(nums[0])
        const queue = new Array(root)

        for (let i = 1; i < nums.length; i += 2) {
            const top = queue.shift()
            top.left = new TreeNode(nums[i])
            queue.push(top.left)
            if (i + 1 < nums.length) {
                top.right = new TreeNode(nums[i + 1])
                queue.push(top.right)
            }
        }
        return root
    }

    function levelOrder(root: TreeNode | null): number[][] {
        const result: number[][] = new Array()
        const queue1 = new Array(root)
        const queue2: TreeNode[] = new Array()

        while (queue1.length > 0 || queue2.length > 0) {
            const row: number[] = new Array()
            if (queue1.length > 0) {
                while (queue1.length > 0) {
                    const top = queue1.shift()
                    if (top != null) {
                        row.push(top.val)
                        if (top.left) queue2.push(top.left)
                        if (top.right) queue2.push(top.right)
                    }
                }
            } else {
                while (queue2.length > 0) {
                    const top = queue2.shift()
                    if (top != null) {
                        row.push(top.val)
                        if (top.left) queue1.push(top.left)
                        if (top.right) queue1.push(top.right)
                    }
                }
            }
            result.push(row)
        }

        return result
    };

    const root = createTree([])
    console.log(levelOrder(root))
}