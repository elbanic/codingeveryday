namespace find_k_length_substrings_with_no_repeated_characters {

    function numKLenSubstrNoRepeats(s: string, k: number): number {

        if (s.length < k) {
            return 0
        }

        let count = 0
        let wind = s.slice(0, k)
        for (let i = 0; i < s.length; i++) {
            wind = s.slice(i, i+k)
            const set = new Set<string>()
            for (let c of wind) {
                if (set.has(c)) {
                    break
                }
                set.add(c)
            }
            if (set.size == k) {
                count++
            }
        }
        return count
    };

    const s = "havefunonleetcode", k = 5
    console.log(numKLenSubstrNoRepeats(s, k))
}