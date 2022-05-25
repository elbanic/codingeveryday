namespace construct_binary_tree_from_string {

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

    function str2tree(s: string): TreeNode | null {
        if (s.length == 0) {
            return null
        }

        const str = s.split(/[(|)]+/).map(x => Number(x))
        const stack: TreeNode[] = new Array(new TreeNode(str[0]))
        let index = 1

        for (let i = 0; i < s.length; i++) {
            if (s[i] == '(') {
                const parent = stack.pop()
                const child = new TreeNode(str[index])
                if (parent.left) {
                    parent.right = child
                } else {
                    parent.left = child
                }
                stack.push(parent)
                stack.push(child)
                index++
            } else if (s[i] == ')') {
                stack.pop()
            }
        }
        return stack.pop()
    };

    function printTree(tr: TreeNode) {
        const queue: TreeNode[] = new Array(tr)
        while (queue.length > 0) {
            const top = queue.shift()
            console.log(top.val)

            if (top.left) queue.push(top.left)
            if (top.right) queue.push(top.right)
        }
    }

    const s = "4(2(3)(1))(6(5))"
    const tree = str2tree(s)
    printTree(tree)
}