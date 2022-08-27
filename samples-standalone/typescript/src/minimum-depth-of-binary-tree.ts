namespace minimum_depth_of_binary_tree {

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
            if (nums[i] != null) {
                top.left = new TreeNode(nums[i])
                queue.push(top.left)
            }

            if (i + 1 < nums.length) {
                if (nums[i + 1] != null) {
                    top.right = new TreeNode(nums[i + 1])
                    queue.push(top.right)
                }

            }
        }
        return root
    }

    function pushData(q1: TreeNode[], q2: TreeNode[]): boolean {

        const top = q1.shift()
        if (top.left == null && top.right == null) {
            return false
        }
        if (top.left != null && top.left.val != null) {
            q2.push(top.left)
        }
        if (top.right != null && top.right.val != null) {
            q2.push(top.right)
        }
        return true
    }

    function minDepth(root: TreeNode | null): number {

        if (root == null) {
            return 0
        }

        const queue1 = new Array(root)
        const queue2: TreeNode[] = new Array()

        let depth = 1
        while (queue1.length > 0 || queue2.length > 0) {
            if (queue1.length > 0) {
                while (queue1.length > 0) {
                    if (!pushData(queue1, queue2)) {
                        return depth
                    }
                }
            } else {
                while (queue2.length > 0) {
                    if (!pushData(queue2, queue1)) {
                        return depth
                    }
                }
            }
            depth++
        }
        return depth
    };

    const root = [0, 0, 0, 0, null, null, 0, null, null, null, 0]
    const rootTree = createTree(root)
    console.log(minDepth(rootTree))
}