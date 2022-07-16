
namespace out_of_boundary_paths {
    function findPaths(m: number, n: number, maxMove: number, startRow: number, startColumn: number): number {

        const memo: number[][][] = new Array()

        for (let i = 0; i < m; i++) {
            const arr1: number[][] = new Array()
            for (let j = 0; j < n; j++) {
                const arr2: number[] = new Array()
                for (let z = 0; z <= maxMove; z++) {
                    arr2.push(-1)
                }
                arr1.push(arr2)
            }
            memo.push(arr1)
        }
        return findPathsHelper(m, n, maxMove, startRow, startColumn, memo)
    };

    function findPathsHelper(m: number, n: number, maxMove: number, startRow: number, startColumn: number, memo: number[][][]): number {
        const Mod = 1000000007;
        if (startRow == m || startColumn == n || startRow < 0 || startColumn < 0) {
            return 1
        }
        if (maxMove == 0) {
            return 0
        }
        if (memo[startRow][startColumn][maxMove] >= 0) {
            return memo[startRow][startColumn][maxMove]
        }
        memo[startRow][startColumn][maxMove] = (
            (findPathsHelper(m, n, maxMove - 1, startRow - 1, startColumn, memo) + findPathsHelper(m, n, maxMove - 1, startRow + 1, startColumn, memo)) % Mod +
            (findPathsHelper(m, n, maxMove - 1, startRow, startColumn - 1, memo) + findPathsHelper(m, n, maxMove - 1, startRow, startColumn + 1, memo)) % Mod
        ) % Mod
        return memo[startRow][startColumn][maxMove]
    }

    const m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
    console.log(findPaths(m, n, maxMove, startRow, startColumn))
}