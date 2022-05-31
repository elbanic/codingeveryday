
namespace parallel_courses {

    function minimumSemesters(n: number, relations: number[][]): number {

        if (relations.length == 0) {
            return -1
        }

        relations.sort((a, b) => b[1] - a[1])
        const lastCourse = relations[0][1]

        const visited = new Map<number, number>()
        for (let i = 0; i <= n; i++) {
            visited.set(i, 0)
        }

        for (let node = 0; node <= n; node++) {
            if (isCycle(relations, node, visited) == -1) {
                return -1
            }
        }

        visited.forEach((value, key) => visited.set(key, 0))
        let maxLength = 1
        for (let node = 0; node <= n; node++) {
            const length = maxPathDFS(relations, node, visited);
            maxLength = Math.max(length, maxLength);
        }
        return maxLength
    };

    function maxPathDFS(relations: number[][], node: number, visited: Map<number, number>): number {

        if (visited.get(node) != 0) {
            return visited.get(node)
        }

        let maxLength = 1
        const children = relations.filter(x => x[1] == node).map(x => x[0])
        for (let child of children) {
            const length = maxPathDFS(relations, child, visited)
            maxLength = Math.max(length + 1, maxLength)
        }
        visited.set(node, maxLength)
        return maxLength
    }

    function isCycle(relations: number[][], node: number, visited: Map<number, number>): number {

        if (visited.get(node) != 0) {
            return visited.get(node)
        } else {
            visited.set(node, -1)
        }

        const children = relations.filter(x => x[1] == node).map(x => x[0])
        for (let child of children) {
            if (isCycle(relations, child, visited) == -1) {
                return -1
            }
        }
        visited.set(node, 1)
        return 1
    }

    const n = 3, relations = [[1, 3], [2, 3]]
    console.log(minimumSemesters(n, relations))
}