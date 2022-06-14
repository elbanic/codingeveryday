namespace valid_palindrome {
    function isPalindrome(s: string): boolean {

        const refined = s.toLowerCase().replace(/[^a-z\d]/g, '');
        for (let i=0; i<refined.length; i++) {
            if (refined[i] != refined[refined.length-1-i]){
                return false
            }
        }
        return true
    };

    const s1 = "aabbaa"
    console.log(isPalindrome(s1))

    const s2 = "A man, a plan, a canal: Panama"
    console.log(isPalindrome(s2))
}