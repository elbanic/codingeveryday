namespace unique_paths_ii {

    //brute force
    function uniquePathsWithObstacles(obstacleGrid: number[][]): number {
        if (obstacleGrid[0][0] != 0) {
            return 0
        }
        return pathCount(obstacleGrid)
    };

    function pathCount(grid: number[][]): number {
        if (grid.length == 1 && grid[0].length == 1) {
            return 1
        }

        //down
        const downGrid = grid.slice(1, grid.length)
        const downCount = downGrid.length > 0 && downGrid[0][0] == 0 ? pathCount(downGrid) : 0

        //right
        const rightGrid = grid.map(x => x.slice(1, x.length))
        const rightCount = rightGrid.length > 0 && rightGrid[0][0] == 0 ? pathCount(rightGrid) : 0

        return downCount + rightCount
    };

    //dynamic programming
    function uniquePathsWithObstacles2(obstacleGrid: number[][]): number {
        const R = obstacleGrid.length
        const C = obstacleGrid[0].length

        if (obstacleGrid[0][0] == 1) {
            return 0
        }

        obstacleGrid[0][0] = 1

        for (let i=1; i < R; i++) {
            obstacleGrid[i][0] = (obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1) ? 1: 0
        }
        for (let i=1; i < C; i++) {
            obstacleGrid[0][i] = (obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1) ? 1: 0
        }

        for (let i = 1; i<R; i++) {
            for (let j = 1; j<C; j++) {
                if (obstacleGrid[i][j] == 0) {
                    obstacleGrid[i][j]= obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
                } else {
                    obstacleGrid[i][j] = 0
                }
            }
        }
            
        return obstacleGrid[R - 1][C - 1]
    };

    const obstacleGrid = [[1,0]]
    console.log(uniquePathsWithObstacles2(obstacleGrid))
}