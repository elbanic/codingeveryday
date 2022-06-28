namespace minimum_deletions_to_make_character_frequencies_unique {
    function minDeletions(s: string): number {

        const map = new Map<string, number>()

        for (let c of s) {
            if (map.has(c)) {
                map.set(c, map.get(c) + 1)
            } else {
                map.set(c, 1)
            }
        }
        return helper(map, 0)
    };

    function helper(m: Map<string, number>, minCost: number): number {

        const set = new Set<number>()
        let key = ''
        for (let [k, v] of m) {
            if (v == 0) {
                m.delete(k)
            } else if (set.has(v)) {
                key = k
                break
            } else {
                set.add(v)
            }
        }
        if (key != '') {
            m.set(key, m.get(key) - 1)
            minCost += 1
            return helper(m, minCost)
        }
        return minCost
    }

    const s = "ceabaacb"
    console.log(minDeletions(s))
}