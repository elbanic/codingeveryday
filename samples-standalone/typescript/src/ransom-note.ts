namespace ransom_note {
    function canConstruct(ransomNote: string, magazine: string): boolean {

        const map = new Map<string, number>()

        for (let c of magazine) {
            if (map.has(c)) {
                map.set(c, map.get(c) + 1)
            } else {
                map.set(c, 1)
            }
        }

        for (let c of ransomNote) {
            if (map.has(c)) {
                if (map.get(c) > 1) {
                    map.set(c, map.get(c) - 1)
                } else {
                    map.delete(c)
                }
            } else {
                return false
            }
        }
        return true
    };

    const ransomNote = "aa", magazine = "aab"
    console.log(canConstruct(ransomNote, magazine))
}