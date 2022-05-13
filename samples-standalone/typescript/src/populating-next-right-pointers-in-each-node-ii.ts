namespace populating_next_right_pointers_in_each_node_ii {

    class Node {
        val: number
        left: Node | null
        right: Node | null
        next: Node | null
        constructor(val?: number, left?: Node, right?: Node, next?: Node) {
            this.val = (val === undefined ? 0 : val)
            this.left = (left === undefined ? null : left)
            this.right = (right === undefined ? null : right)
            this.next = (next === undefined ? null : next)
        }
    }


    function connect(root: Node | null): Node | null {

        if (root == null) {
            return null
        }

        const queueA: Node[] = new Array(root)
        const queueB: Node[] = new Array()

        while (queueA.length > 0 || queueB.length > 0) {
            if (queueA.length > 0) {
                let cur = queueA.shift()
                if (cur.left) queueB.push(cur.left)
                if (cur.right) queueB.push(cur.right)
                while (queueA.length > 0) {
                    cur.next = queueA.shift()
                    cur = cur.next
                    if (cur.left) queueB.push(cur.left)
                    if (cur.right) queueB.push(cur.right)
                }
                cur.next = null
            } else if (queueB.length >0) {
                let cur = queueB.shift()
                if (cur.left) queueA.push(cur.left)
                if (cur.right) queueA.push(cur.right)
                while (queueB.length > 0) {
                    cur.next = queueB.shift()
                    cur = cur.next
                    if (cur.left) queueA.push(cur.left)
                    if (cur.right) queueA.push(cur.right)
                }
                cur.next = null
            }
        }
        return root
    };

    function printNode(root:Node) {

        if (root == null) {
            return
        }

        const queue: Node[] = new Array(root)
        while (queue.length > 0) {
            const node = queue.shift()
            console.log(node.val)
            if (node.left) queue.push(node.left)
            if (node.right) queue.push(node.right)
        }
    }

    function createTree(nums: number[]): Node | null{

        if (nums.length == 0) {
            return null
        }

        const root: Node = new Node(nums[0])
        const queue: Node[] = new Array(root)

        for (let i=1; i<nums.length; i+=2) {
            const top = queue.shift()
            top.left = new Node(nums[i])
            queue.push(top.left)
            if (i+1 < nums.length) {
                top.right = new Node(nums[i+1])
                queue.push(top.right)
            }
        }
        return root
    }

    const root = [1,2,3,4,5,null,7]
    const rootNode = createTree(root)
    printNode(rootNode)
    connect(rootNode)
    printNode(rootNode)
}