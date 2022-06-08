namespace remove_palindromic_subsequences {

    function removePalindromeSub(s: string): number {

        return helper(s, 0)
    };

    function helper(s: string, depth: number): number {
        if (isPalindrome(s)) {
            return depth+1
        }

        if (s == '') {
            return depth
        }
    
        let ch = s[0]
        s = s.slice(1, s.length)
        for (let i=s.length-1; i>=0; i--) {
            if (ch == s[i]) {
                s = s.slice(0, i) + s.slice(i+1);
                ch = s[0]
                s = s.slice(1, s.length)
            }
        }
        
        return helper(s, depth+1)
    };

    function isPalindrome(s: string): boolean {

        for (let i=0; i<s.length/2; i++) {
            if (s[i] != s[s.length-1-i]) {
                return false
            }
        }
        return true
    }

    const s = "abb"
    console.log(removePalindromeSub(s))
}