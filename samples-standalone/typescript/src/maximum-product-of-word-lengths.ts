namespace maximum_product_of_word_lengths {

    //brute force
    function maxProduct(words: string[]): number {

        let max = 0
        for (let i of words) {
            const iMap = new Set<string>()
            for (let ii of i) {
                iMap.add(ii)
            }

            for (let j of words) {
                let has = false
                for (let jj of j) {
                    if (iMap.has(jj)) {
                        has = true
                        break
                    }
                }
                if (!has) {
                    const product = i.length * j.length
                    if (product > max) {
                        max = product
                    }
                }
            }
        }
        return max
    };

    const words = ["a","ab","abc","d","cd","bcd","abcd"]
    console.log(maxProduct(words))
}