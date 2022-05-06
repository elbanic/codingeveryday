namespace remove_all_adjacent_duplicates_in_string_ii {


    function removeDuplicates(s: string, k: number): string {

        if (s.length < k) {
            return ""
        }

        const m: Map<string, number> = new Map<string, number>()
        let wind = 0
        for (let i = 0; i < s.length; i++) {

            if (m.has(s[i])) {
                m.set(s[i], m.get(s[i]) + 1)
            } else {
                m.set(s[i], 1)
            }

            if (wind == k) {
                if (m.get(s[i - k]) > 1) {
                    m.set(s[i - k], m.get(s[i - k]) - 1)
                } else {
                    m.delete(s[i - k])
                }
            } else {
                wind++
            }

            if (m.size == 1 && m.get(s[i]) == k) {
                s = s.replace(s[i].repeat(k), '')
                m.clear()
                i = -1
                wind = 0
            }
        }
        return s
    };

    // using stack
    function removeDuplicates2(s: string, k: number): string {
        const cnt: number[] = new Array()

        for (let i=0; i<s.length; i++) {
            if (i == 0 || s[i] != s[i-1]) {
                cnt.push(1)
            } else {
                const incremented = cnt.pop() + 1
                if (incremented == k) {
                    s = s.replace(s.slice(i-k+1, i+1), '')
                    i = i - k
                } else {
                    cnt.push(incremented)
                }
            }
        }
        return s
    }

    const s = "deeedbbcccbdaa", k = 3
    console.log(removeDuplicates2(s, k))
}