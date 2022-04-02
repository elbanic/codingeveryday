/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function(s) {

    let ret = isPalindrome(s);
    if (ret[0]) {
        return true
    } else {
        const s1 = s.slice(0, ret[1]) + s.slice(ret[1]+1);
        const s2 = s.slice(0, s.length - 1 - ret[1]) + s.slice(s.length - ret[1]);

        return isPalindrome(s1)[0] || isPalindrome(s2)[0]
    }
};

var isPalindrome = function(s) {
    var i;
    for (i = 0; i<s.length/2; i++) {
        if (s[i] != s[s.length - 1 - i]) {
            return [false, i]
        }
    }
    return [true, i]
}


console.log(validPalindrome("aba"))
console.log(validPalindrome("abca"))
console.log(validPalindrome("abc"))

