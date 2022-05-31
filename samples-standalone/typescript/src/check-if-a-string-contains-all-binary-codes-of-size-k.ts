
namespace check_if_a_string_contains_all_binary_codes_of_size_k {

    function hasAllCodes2(s: string, k: number): boolean {
        let need = 1 << k
        const got : Set<string> = new Set<string>()

        for (let i=k; i<=s.length; i++) {
            let a = s.slice(i-k, i)
            if (!got.has(a)) {
                got.add(a)
                need--

                if (need == 0) {
                    return true
                }
            }
        }
        return false
    }

    function hasAllCodes(s: string, k: number): boolean {

        if (s.length < k) {
            return false
        }

        const codes = getCodes(k)
        const subs: string[] = new Array()
        for (let i=0; i<=s.length-k; i++) {
            if (codes.has(s.slice(i, i+k))){
                codes.set(s.slice(i, i+k), 0)
            }
        }

        for (let i of codes) {
            if (i[1] == 1) {
                return false
            }
        }
        return true
    };

    function getCodes(k: number): Map<string, number> {
        const codes: Map<string, number> = new Map<string, number>()
        for (let n of [...Array(Math.pow(2, k)).keys()]) {
            codes.set((n >>> 0).toString(2).padStart(k, '0'), 1)
        }
        return codes
    }
    
    const s = "1011110111001001101001110001100111101111010101011100111001110010010001000111010110101110000110101001011100100010100110011101011110001000100010101101011", k = 20
    console.log(hasAllCodes(s, k))
}