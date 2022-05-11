namespace count_sorted_vowel_strings {

    //brute force
    function countVowelStrings(n: number): number {
        if (n == 0) return 0

        let comb: string[] = new Array()
        "aeiou".split('').forEach(x => comb.push(x))

        for (let i = 1; i < n; i++) {
            const temp = comb
            comb = new Array()
            const vowel = "aeiou".split('')
            for (let v of vowel) {
                for (let c of temp) {
                    if (c[c.length - 1] <= v) {
                        comb.push(c + v)
                    }
                }
            }
        }
        return comb.length
    };

    //better solution - 
    function countVowelStrings2(n: number): number {
        if (n == 0) return 0

        const cnt: number[] = [1, 1, 1, 1, 1]
        for (let i = 0; i < n - 1; i++) {
            for (let j = 1; j < cnt.length; j++) {
                cnt[j] = cnt[j - 1] + cnt[j]
            }
        }
        return cnt.reduce((sum, x) => sum + x)
    }

    const n = 5
    console.log(countVowelStrings2(n))
}