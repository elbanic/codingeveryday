
namespace binary_search_tree_iterator_ii {

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


    class BSTIterator {
        stack: TreeNode[]
        pointer: number
        constructor(root: TreeNode | null) {
            this.pointer = -1
            this.stack = new Array();
            let tmpStack: TreeNode[] = new Array()
            let node = root
            while (tmpStack.length != 0 || node != null) {
                while (node != null) {
                    tmpStack.push(node)
                    node = node.left
                }
                let left = tmpStack.pop()
                this.stack.push(left)
                if (left.right != null) {
                    node = left.right
                }
            }
        }

        hasNext(): boolean {
            return this.pointer < this.stack.length - 1
        }

        next(): number {
            this.pointer = Math.min(this.pointer + 1, this.stack.length)
            return this.stack[this.pointer].val
        }

        hasPrev(): boolean {
            return this.pointer > 0
        }

        prev(): number {
            this.pointer = Math.max(this.pointer - 1, 0)
            return this.stack[this.pointer].val
        }
    }

    function createTree(nums: number[]): TreeNode {
        if (nums.length == 0) {
            return null;
        }

        let root: TreeNode = new TreeNode(nums[0]);
        let queue: TreeNode[] = new Array(root);

        for (let i = 1; i < nums.length; i += 2) {
            let top: TreeNode = queue.shift();
            if (nums[i] == null) {
                top.left = null;
            } else {
                let left: TreeNode = new TreeNode(nums[i]);
                top.left = left;
                queue.push(left);
            }
            if (i + 1 < nums.length) {
                if (nums[i + 1] == null) {
                    top.right = null;
                } else {
                    let right: TreeNode = new TreeNode(nums[i + 1]);
                    top.right = right;
                    queue.push(right);
                }
            }
        }
        return root;
    }

    let nums = [1];
    let root = createTree(nums);

    var obj = new BSTIterator(root)

    console.log(obj.hasPrev());
    console.log(obj.hasNext());
    console.log(obj.next());
    console.log(obj.hasPrev());
    console.log(obj.hasNext());
}

