namespace pascals_triangle {
    function generate(numRows: number): number[][] {

        if (numRows == 0) {
            return new Array()
        }

        const res: number[][] = new Array([1])

        for (let i=1; i<numRows; i++){

            const row = new Array()
            for (let j=0; j<=i; j++) {
                if (j == 0 || j == i) {
                    row.push(1)
                } else {
                    row.push(res[i-1][j-1] + res[i-1][j])
                }
            }
            res.push(row)
        }
        return res
    };

    
    generate(5).forEach(x => console.log(x))
}