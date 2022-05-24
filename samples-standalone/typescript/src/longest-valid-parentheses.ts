namespace longest_valid_parentheses {
 
    function longestValidParentheses(s: string): number {

        let max = 0
        const dp: number[] = new Array()
        for (let i=0; i<s.length; i++) {
            dp.push(0)
        }

        for (let i=1; i<s.length; i++) {
            if (s[i] == ')') {
                if (s[i-1] == '(') {
                    dp[i] = (i >= 2 ? dp[i-2] : 0) + 2
                } else if (i - dp[i-1] >0 && s[i - dp[i-1] - 1] == '(') {
                    dp[i] = dp[i-1] + ((i- dp[i-1]) >= 2? dp[i - dp[i-1] -2 ] : 0) + 2
                }
                max = Math.max(max, dp[i])
            }
        }
        return max
    };

    const s = "()(()"
    console.log(longestValidParentheses(s))

    const s2 = ")()())"
    console.log(longestValidParentheses(s2))
}