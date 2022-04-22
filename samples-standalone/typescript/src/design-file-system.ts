namespace design_file_system {

    class TreeNode {
        name: string
        isDir: boolean
        val: number
        children: TreeNode[]
        constructor(name: string, isDir: boolean, val?: number) {
            this.name = name
            this.isDir = isDir
            this.val = (val === undefined ? 0 : val)
            this.children = new Array()
        }
    }

    class FileSystem {
        tree: TreeNode
        constructor() {
            this.tree = new TreeNode("/", true)
        }

        createPath(path: string, value: number): boolean {
            if (this.get(path) == -1) {
                let paths = path.split("/")
                let curr = this.tree
                let i = 1
                while (curr != null) {
                    if (i >= paths.length) {
                        return false
                    }

                    let next = curr.children.find(n => n.name == paths[i])
                    if (next) {
                        curr = next
                        i++
                    } else {
                        break
                    }
                }
                if (i == paths.length-1) {
                    curr.children.push(new TreeNode(paths[paths.length - 1], false, value))
                    return true
                } else {
                    return false
                }
                
            } else {
                return false
            }
        }

        get(path: string): number {
            let paths = path.split("/")
            let curr = this.tree
            let i = 1
            while (curr != null) {
                if (i >= paths.length) {
                    break
                }

                let next = curr.children.find(n => n.name == paths[i])
                if (next) {
                    curr = next
                    if (i == paths.length-1) {
                        return curr.val
                    }
                    i++
                } else {
                    break
                }
            }
            return -1
        }
    }

    var obj = new FileSystem()
    console.log(obj.createPath("/leet", 1))
    console.log(obj.createPath("/leet/code", 2))
    console.log(obj.get("/leet/code"))
    console.log(obj.createPath("/c/d", 1))
    console.log(obj.get("/c"))
}