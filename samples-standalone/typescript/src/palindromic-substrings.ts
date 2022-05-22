
namespace palindromic_substrings {
    //brute force
    function countSubstrings(s: string): number {

        let cnt = s.length
        for (let size = 2; size<=s.length; size++) {

            for (let i =0; i<=s.length-size; i++) {
                const sub = s.slice(i, i+size)
                if (isPalindrom(sub)) {
                    cnt++
                }
            }
        }
        return cnt
    };

    function isPalindrom(s: string): boolean {
        for (let i =0; i<s.length/2; i++) {
            if (s[i] != s[s.length - 1 - i]) {
                return false
            }
        }
        return true
    }

    const s = "aaa"
    console.log(countSubstrings(s))
}