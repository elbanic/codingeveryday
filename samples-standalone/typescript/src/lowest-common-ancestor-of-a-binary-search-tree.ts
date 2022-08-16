namespace lowest_common_ancestor_of_a_binary_search_tree {

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

    function lowestCommonAncestor(root: TreeNode | null, p: TreeNode | null, q: TreeNode | null): TreeNode | null {

        const queue: TreeNode[][] = new Array()
        queue.push(new Array(root))

        let pAnc: TreeNode[], qAnc: TreeNode[]

        while (queue.length > 0) {
            const last = queue.shift()
            if (last[last.length - 1].val == p.val) {
                pAnc = [...last]
            }
            if (last[last.length - 1].val == q.val) {
                qAnc = [...last]
            }
            if (pAnc != undefined && qAnc != undefined) {
                break
            }

            if (last[last.length - 1].left != null) {
                const temp = [...last]
                temp.push(last[last.length - 1].left)
                queue.push(temp)
            }
            if (last[last.length - 1].right != null) {
                const temp = [...last]
                temp.push(last[last.length - 1].right)
                queue.push(temp)
            }
        }

        let cur = pAnc[0]
        for (let i = 0; i < Math.min(pAnc.length, qAnc.length); i++) {
            if (pAnc[i].val == qAnc[i].val) {
                cur = pAnc[i]
            } else {
                break
            }
        }
        return cur
    };



    const root = [6, 2, 8, 0, 4, 7, 9, null, null, 3, 5], p = 2, q = 4
    const rootTree = createTreeNode(root)
    console.log(lowestCommonAncestor(rootTree, new TreeNode(p), new TreeNode(q)))

}