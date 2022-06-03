
namespace range_sum_query_2d_immutable {

    class NumMatrix2 {
        dp: number[][]
        constructor(matrix: number[][]) {
            if (matrix.length == 0 || matrix[0].length == 0) { 
                return 
            }
            this.dp = new Array()

            for (let i=0; i<= matrix.length; i++) {
                const row = new Array()
                for (let j =0; j<= matrix[0].length; j++) {
                    row.push(0)
                }
                this.dp.push(row)
            }

            for (let row=0; row<matrix.length; row++) {
                for (let col=0; col<matrix[0].length; col++) {
                    this.dp[row+1][col+1] = this.dp[row+1][col] + this.dp[row][col+1] + matrix[row][col] - this.dp[row][col]
                }
            }
        }

        sumRegion(row1: number, col1: number, row2: number, col2: number): number {
            return this.dp[row2 + 1][col2 + 1] - this.dp[row1][col2 + 1] - this.dp[row2 + 1][col1] + this.dp[row1][col1];
        }
    }

    //caching
    class NumMatrix {
        cache: Map<string, number>
        matrix: number[][]
        n: number
        m: number
        constructor(matrix: number[][]) {
            this.matrix = matrix
            this.m = this.matrix.length, this.n = this.matrix[0].length
            this.cache = new Map<string, number>()
        }

        getRegion(row1: number, col1: number, row2: number, col2: number): number[][] {
            if (row1 >= this.m || row2 >= this.m || col1 >= this.n || col2 >= this.n) {
                return null
            }

            const newMatrix: number[][] = new Array()
            for (let i = row1; i <= row2; i++) {
                const row = new Array()
                for (let j = col1; j <= col2; j++) {
                    row.push(this.matrix[i][j])
                }
                newMatrix.push(row)
            }
            return newMatrix
        }

        sumRegion(row1: number, col1: number, row2: number, col2: number): number {

            if (this.cache.has(JSON.stringify({ row1, col1, row2, col2 }))) {
                return this.cache.get(JSON.stringify({ row1, col1, row2, col2 }))
            }

            const newMatrix = this.getRegion(row1, col1, row2, col2)
            let sum = 0
            for (let row of newMatrix) {
                row.forEach(x => sum += x)
            }
            this.cache.set(JSON.stringify({ row1, col1, row2, col2 }), sum)
            return sum
        }
    }


    const matrix = [[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]
    var obj = new NumMatrix(matrix)
    console.log(obj.sumRegion(2, 1, 4, 3))
    console.log(obj.sumRegion(1, 1, 2, 2))
    console.log(obj.sumRegion(1, 2, 2, 4))
}