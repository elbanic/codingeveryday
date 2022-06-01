namespace minimum_knight_moves {

    const dir = [[1, 2], [2, 1], [2, -1], [1, -2], [-1, -2], [-2, -1], [-2, 1], [-1, 2]]

    class Coord {
        constructor(public x: number, public y: number) { }
        getPath(): Coord[] {
            return dir.map(i => new Coord(this.x + i[0], this.y + i[1]))
        }
    }

    //brute force
    function minKnightMoves(x: number, y: number): number {

        if (x == 0 && y == 0) {
            return 0
        }

        const start = new Coord(0, 0)
        const target = new Coord(x, y)
        const pathsQueue = start.getPath().map(i => new Coord(i.x, i.y))
        const visited: boolean[][] = new Array()

        for (let i = 0; i <= 607; i++) {
            const v: boolean[] = new Array()
            for (let j = 0; j <= 607; j++) {
                v.push(false)
            }
            visited.push(v)
        }

        let count = 1
        while (pathsQueue.length > 0) {

            const size = pathsQueue.length
            for (let i = 0; i < size; i++) {
                const path = pathsQueue.shift()
                if (path.x == target.x && path.y == target.y) {
                    return count
                }
                path.getPath().forEach(i => {
                    if (!visited[i.x + 302][i.y + 302]) {
                        pathsQueue.push(i)
                        visited[i.x + 302][i.y + 302] = true
                    }
                })
            }
            count++
        }
        return count
    };

    const x = -84, y = 170
    console.log(minKnightMoves(x, y))
}