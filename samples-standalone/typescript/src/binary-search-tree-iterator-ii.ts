
import { TreeNode, createTree } from "./common-tools";

namespace binary_search_tree_iterator_ii {

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

    let nums = [1];
    let root = createTree(nums);

    var obj = new BSTIterator(root)

    console.log(obj.hasPrev());
    console.log(obj.hasNext());
    console.log(obj.next());
    console.log(obj.hasPrev());
    console.log(obj.hasNext());
}

