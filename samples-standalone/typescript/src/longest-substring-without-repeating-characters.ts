namespace longest_substring_without_repeating_characters {
 
    function lengthOfLongestSubstring(s: string): number {
        if (s.length == 0) {
            return 0
        }

        let st = 0, end = 1
        let max = end - st
        const map = new Map<string, number>()
        map.set(s[st], st)

        while (end < s.length) {
            if (map.has(s[end])) {
                st = Math.max(st, map.get(s[end]) + 1)
            }
            map.set(s[end], end)
            end++

            max = Math.max(max, end - st)
        }

        return max
    };

    const s = "abba"
    console.log(lengthOfLongestSubstring(s))
}