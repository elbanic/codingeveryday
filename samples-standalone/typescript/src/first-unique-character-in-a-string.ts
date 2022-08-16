namespace first_unique_character_in_a_string {

    function firstUniqChar(s: string): number {

        const map = new Map<string, number>()

        for (let c of s) {
            if (map.has(c)) {
                map.set(c, map.get(c)+1)
            } else {
                map.set(c, 1)
            }
        }

        for (let i=0; i<s.length; i++) {
            if (map.get(s[i]) == 1) return i
        }
        return -1
    };

    console.log(firstUniqChar("loveleetcode"))
    console.log(firstUniqChar("aabb"))
}