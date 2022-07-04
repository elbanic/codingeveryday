namespace symmetric_tree {

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

        const root = new TreeNode(nums[0])
        const queue: TreeNode[] = new Array()

        queue.push(root)
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

    function isSymmetric(root: TreeNode | null): boolean {

        const queueL: TreeNode[] = new Array()
        const queueR: TreeNode[] = new Array()

        queueL.push(root)
        queueR.push(root)
        while (queueL.length > 0) {

            const topL = queueL.shift()
            const topR = queueR.shift()
            if (topL.val != topR.val) {
                return false
            }
            if (topL.left) {
                queueL.push(topL.left)
                if (!topR.right) return false
                queueR.push(topR.right)
            }
            if (topL.right) {
                queueL.push(topL.right)
                if (!topR.left) return false
                queueR.push(topR.left)
            }
        }
        return true
    };

    const root = createTree([1, 2, 2, 3, 4, 4, 3])
    console.log(isSymmetric(root))
}