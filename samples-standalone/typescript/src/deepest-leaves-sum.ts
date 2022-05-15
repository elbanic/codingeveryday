namespace deepest_leaves_sum {

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

    function deepestLeavesSum(root: TreeNode | null): number {

        class LeaveSum {
            constructor(public trees: TreeNode[], public sum: number) { }
            add(tr: TreeNode) {
                this.trees.push(tr)
                this.sum += tr.val
            }
        }

        const queueA: TreeNode[] = new Array(root)
        const queueB: TreeNode[] = new Array()
        let lastStatus: TreeNode[]

        let deepestLeaves = 0
        while (queueA.length > 0 || queueB.length > 0) {
            if (queueA.length > 0) {
                lastStatus = [...queueA]
                while (queueA.length > 0) {
                    const node = queueA.shift()
                    if (node.left) {
                        queueB.push(node.left)
                    }
                    if (node.right) {
                        queueB.push(node.right)
                    }
                }
            } else if (queueB.length > 0) {
                lastStatus = [...queueB]
                while (queueB.length > 0) {
                    const node = queueB.shift()
                    if (node.left) {
                        queueA.push(node.left)
                    }
                    if (node.right) {
                        queueA.push(node.right)
                    }
                }
            }
        }
        return lastStatus.reduce((sum, x) => { return sum + x.val }, 0);
    };

    function createTreeNode(nums: number[]): TreeNode {
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

    const root = [1, 2, 3, 4, 5, null, 6, 7, null, null, null, null, 8]
    const rootNode = createTreeNode(root)
    console.log(deepestLeavesSum(rootNode))

}