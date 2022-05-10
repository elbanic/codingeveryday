namespace combination_sum_iii {

    function combinationSum3(k: number, n: number): number[][] {

        const comb = combination(k, Array.from(Array(9).keys()).map(x => x + 1), new Array(), new Array())
        const filtered = comb.filter(x => x.reduce((sum, current) => sum + current, 0) == n)
        return filtered
    };

    function combination(depth: number, candidate: number[], prev: number[], comb: number[][]): number[][] {

        if (prev.length == depth) {
            const temp = [...prev]
            comb.push(temp)
            return comb
        }

        for (let i=0; i<candidate.length; i++) {
            prev.push(candidate[i])
            comb = combination(depth, candidate.slice(i+1), prev, comb)
            prev.pop()
        }
        return comb
    }

    const k = 9, n = 45
    console.log(combinationSum3(k, n))
}