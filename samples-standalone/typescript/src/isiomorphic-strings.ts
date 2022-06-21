namespace isiomorphic_strings {
    function isIsomorphic(s: string, t: string): boolean {
        if (s.length != t.length) {
            return false
        }
        return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
    };

    function isIsomorphicHelper(s: string, t: string): boolean {
        let seen = new Set<string>()
        let map = new Map<string, string>()
        for (let i = 0; i < s.length; i++) {
            if (!seen.has(s[i])) {
                seen.add(s[i])
                map.set(s[i], t[i])
            } else {
                if (map.get(s[i]) != t[i]) {
                    return false
                }
            }
        }
        return true
    }

    console.log(isIsomorphic("egg", "add"))
    console.log(isIsomorphic("bbbaaaba", "aaabbbba"))
    console.log(isIsomorphic("badc", "baba"))
    console.log(isIsomorphic("paper", "title"))

}