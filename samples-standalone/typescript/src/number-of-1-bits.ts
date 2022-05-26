namespace number_of_1_bits {
    function hammingWeight(n: number): number {
        let count = 0
        for (let i of n.toString(2)) {
            if (i == '1') {
                count++
            }
        }
        return count
    };

    const n = 11
    console.log(hammingWeight(n))
}