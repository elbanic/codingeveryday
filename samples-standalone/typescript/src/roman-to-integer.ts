namespace roman_to_integer {


    function romanToInt(s: string): number {

        const map = new Map<string, number>()
        map.set("I", 1)
        map.set("V", 5)
        map.set("X", 10)
        map.set("L", 50)
        map.set("C", 100)
        map.set("D", 500)
        map.set("M", 1000)

        let sum = 0
        let next: string

        for (let i = 0; i < s.length; i++) {
            if (i < s.length - 1) {
                next = s[i + 1]
            } else {
                next = undefined
            }

            if (next != undefined) {
                if (((next == "V" || next == "X") && s[i] == "I") ||
                    ((next == "L" || next == "C") && s[i] == "X") ||
                    ((next == "D" || next == "M") && s[i] == "C")) {
                    sum += map.get(next) - map.get(s[i])
                    i++
                    continue
                }
            } 
            sum += map.get(s[i])
        }
        return sum
    };

    const s = "LVIII"
    console.log(romanToInt(s))
}