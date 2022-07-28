namespace valid_anagram {
    function isAnagram(s: string, t: string): boolean {

        if (s.length != t.length) {
            return false
        }

        const map = new Map<string, number>()
        for (let i of s) {
            if (map.has(i)) {
                map.set(i, map.get(i)+1)
            } else {
                map.set(i, 1)
            }
        }

        for (let i of t) {
            if (map.has(i)) {
                if (map.get(i) > 1) {
                    map.set(i, map.get(i) - 1)
                } else {
                    map.delete(i)
                }
            } else {
                return false
            }
        }
        return true
    };

    const s = "anagram", t = "nagaram"
    console.log(isAnagram(s,t ))
}