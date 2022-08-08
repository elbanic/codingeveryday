
namespace binary_tree_vertical_order_traversal {

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

    class NodeInfo{
        constructor(public tr: TreeNode, public row: number) {}
    }

    function verticalOrder(root: TreeNode | null): number[][] {

        if (root == null) {
            return new Array()
        }
        const pair: Map<number, NodeInfo[]> = new Map()
        const [minDepth, maxDepth] = traversal(root, 0, 0, 0, 0, pair)

        const res: number[][] = new Array()
        let depth = minDepth
        while (depth <= maxDepth) {
            const nodesWithRow = pair.get(depth)
            nodesWithRow.sort((a, b) => b.row - a.row)
            res.push(nodesWithRow.map(i => i.tr.val))
            depth++
        }
        return res
    };    

    function traversal(node: TreeNode, col: number, row: number, minDepth: number, maxDepth: number, pair: Map<number, NodeInfo[]>): [number, number] {

        if (node == null) {
            return
        }

        if (minDepth > col) {
            minDepth = col
        }
        if (maxDepth < col) {
            maxDepth = col
        }

        if (pair.has(col)) {
            pair.get(col).push(new NodeInfo(node, row))
        } else {
            pair.set(col, new Array(new NodeInfo(node, row)))
        }
        if (node.left) {
            [minDepth, maxDepth] = traversal(node.left, col - 1, row - 1, minDepth, maxDepth, pair)
        }
        if (node.right) {
            [minDepth, maxDepth] = traversal(node.right, col + 1, row - 1, minDepth, maxDepth, pair)
        }
        return [minDepth, maxDepth]
    }

    const root = []
    const rootTree = createTreeNode(root)
    console.log(verticalOrder(rootTree))

}