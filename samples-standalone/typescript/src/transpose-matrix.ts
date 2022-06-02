namespace transpose_matrix {

    function transpose(matrix: number[][]): number[][] {

        if (matrix.length == 0) {
            return null
        }
        const m = matrix.length, n = matrix[0].length
        const newMatrix: number[][] = new Array()
        for (let i=0; i<n; i++) {
            const row: number[] = new Array()
            for (let j=0; j<m; j++) {
                row.push(0)
            }
            newMatrix.push(row)
        }

        for (let i=0; i<m; i++) {
            for (let j=0; j<n; j++) {
                newMatrix[j][i] = matrix[i][j]
            }
        }

        return null
    };

    const matrix = [[1,2,3],[4,5,6]]
    console.log(transpose(matrix))    
}