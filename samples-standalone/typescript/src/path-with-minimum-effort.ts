namespace path_with_minimum_effort {

    function minimumEffortPath(heights: number[][]): number {

        const paths = getPath(heights, 0, 0, new Array(), new Array())
        let minDist = Number.MAX_SAFE_INTEGER

        for (let path of paths) {
            let leastDist = 0
            for (let i = 0; i < path.length - 1; i++) {
                if (path[i + 1] - path[i] > 0 && path[i + 1] - path[i] > leastDist) {
                    leastDist = path[i + 1] - path[i]
                }
            }
            minDist = leastDist < minDist ? leastDist : minDist
        }
        return minDist
    };

    function getPath(heights: number[][], i: number, j: number, path: number[], paths: number[][]): number[][] {

        if (i < 0 || i >= heights.length || j >= heights.length) {
            return paths
        }

        if (i == heights[0].length - 1 && j == heights.length - 1) {
            path.push(heights[i][j])
            paths.push(path)
            return paths
        }

        path.push(heights[i][j])

        const right = getPath(heights, i, j + 1, [...path], paths)
        const bottom = getPath(heights, i + 1, j, [...path], paths)

        return paths
    }


    const heights = [[1, 2, 3], [3, 8, 4], [5, 3, 5]]
    console.log(minimumEffortPath(heights))

}