namespace inorder_successor_in_bst_ii {

    class Node {
        val: number
        left: Node | null
        right: Node | null
        parent: Node | null
        constructor(val?: number, left?: Node | null, right?: Node | null, parent?: Node | null) {
            this.val = (val === undefined ? 0 : val)
            this.left = (left === undefined ? null : left)
            this.right = (right === undefined ? null : right)
            this.parent = (parent === undefined ? null : parent)
        }
    }

    function createNode(nums: number[]): Node{
        const root = new Node(nums[0])
        const queue = new Array(root)
        for (let i = 1; i<nums.length; i+=2) {
            const top = queue.shift()
            top.left = new Node(nums[i])
            top.left.parent = top
            queue.push(top.left)
            if (i+1 < nums.length){
                top.right = new Node(nums[i+1])
                top.right.parent = top
                queue.push(top.right)
            }
        }
        return root
    }

    function inorderSuccessor(node: Node | null): Node | null {

        if (node.right != null) {
            node = node.right
            while (node.left != null) {
                node = node.left
            }
            return node
        }

        while (node.parent != null && node == node.parent.right) {
            node = node.parent
        }
        return node.parent
    };

    const tree = [2,1,3], node = 1
    console.log(inorderSuccessor(createNode(tree)))
}