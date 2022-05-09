namespace letter_combinations_of_a_phone_number {

    const m: Map<number, string[]> = new Map<number, string[]>()
    m.set(2, "abc".split(''))
    m.set(3, "def".split(''))
    m.set(4, "ghi".split(''))
    m.set(5, "jkl".split(''))
    m.set(6, "mno".split(''))
    m.set(7, "pqrs".split(''))
    m.set(8, "tuv".split(''))
    m.set(9, "wxyz".split(''))

    function letterCombinations(digits: string): string[] {

        if (digits == "") {
            return new Array()
        }

        let comb = new Array()
        const digitArr = digits.split('')

        for (let i of m.get(Number(digitArr[0]))) {
            comb.push(i)
        }

        for (let i = 1; i<digitArr.length; i++) {
            const prevComb = comb
            comb = new Array()
            for (let j of m.get(Number(digitArr[i]))) {
                for (let k of prevComb) {
                    comb.push(k+j)
                }
            }
        }
        return comb
    };
    
    const digits = "234"
    console.log(letterCombinations(digits))
}