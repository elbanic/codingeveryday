namespace find_a_corresponding_node_of_a_binary_tree_in_a_clone_of_that_tree {

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

    function copyTree(root: TreeNode): TreeNode {
        const newRoot = new TreeNode(root.val)

        let queueForOrg: TreeNode[] = new Array(root);
        let queueForNew: TreeNode[] = new Array(newRoot);
        while (queueForOrg.length > 0) {
            const topForOrg = queueForOrg.shift()
            const topForNew = queueForNew.shift()
            if (topForOrg.left) {
                topForNew.left = new TreeNode(topForOrg.left.val)
                queueForOrg.push(topForOrg.left)
                queueForNew.push(topForNew.left)
            }
            if (topForOrg.right) {
                topForNew.right = new TreeNode(topForOrg.right.val)
                queueForOrg.push(topForOrg.right)
                queueForNew.push(topForNew.right)
            }
        }
        return newRoot
    }

    function getTargetCopy(original: TreeNode | null, cloned: TreeNode | null, target: TreeNode | null): TreeNode | null {

        if (original == null || cloned == null) {
            return null
        }

        let queueForOrg: TreeNode[] = new Array(original);
        let queueForCld: TreeNode[] = new Array(cloned);
        while (queueForOrg.length > 0) {
            const topForOrg = queueForOrg.shift()
            const topForCld = queueForCld.shift()
            if (target.val == topForOrg.val) {
                return topForCld
            }

            if (topForOrg.left) {
                queueForOrg.push(topForOrg.left)
                queueForCld.push(topForCld.left)
            }
            if (topForOrg.right) {
                queueForOrg.push(topForOrg.right)
                queueForCld.push(topForCld.right)
            }
        }

        return null
    };

    const tree = [7, 4, 3, null, null, 6, 19]
    const treeRoot = createTree(tree)
    const treeCopy = copyTree(treeRoot)
    const target = new TreeNode(3)

    console.log(getTargetCopy(treeRoot, treeCopy, target))
}