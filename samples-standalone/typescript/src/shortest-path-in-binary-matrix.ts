
namespace shortest_path_in_binary_matrix {

    class Coord {
        constructor(public x: number, public y: number) { }
        getNeighbour(grid: number[][], visited: Set<string>): boolean {
            return this.x >= 0 && this.y >= 0 && this.x < grid.length && this.y < grid.length && !visited.has(JSON.stringify(this))
        }
    }

    const neighbours = [[1, 1], [1, -1], [-1, 1], [-1, -1], [0, 1], [1, 0], [-1, 0], [0, -1]]

    function shortestPathBinaryMatrix(grid: number[][]): number {
        if (grid.length == 0) {
            return -1
        }

        if (grid[0][0] == 1) {
            return -1
        }

        const queue: Coord[] = new Array(new Coord(0, 0))
        const visited: Set<string> = new Set<string>()
        const n = grid.length
        grid[0][0] = 1

        while (queue.length > 0) {
            const top = queue.shift()

            if (top.x == n - 1 && top.y == n - 1) {
                return grid[top.x][top.y]
            }

            visited.add(JSON.stringify(top))
            const depth = grid[top.x][top.y]

            for (let i of neighbours) {
                const curCoord = new Coord(top.x+i[0], top.y+i[1])
                if (curCoord.getNeighbour(grid, visited) && grid[curCoord.x][curCoord.y] == 0) {
                    grid[curCoord.x][curCoord.y] = depth + 1
                    queue.push(curCoord)
                }
            }
        }
        return -1
    };

    const grid = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
    console.log(shortestPathBinaryMatrix(grid))
}