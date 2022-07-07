namespace interleaving_string {

    function isInterleave(s1: string, s2: string, s3: string): boolean {
        if (s1 == "" && s2 == "" && s3 == "") {
            return true
        }

        if (s3 == "") {
            return false
        }

        if (s3[0] == s1[0]) {
            if (isInterleave(s1.slice(1), s2, s3.slice(1))) {
                return true
            }
        }
        if (s3[0] == s2[0]) {
            if (isInterleave(s1, s2.slice(1), s3.slice(1))) {
                return true
            }
        }
        return false
    };

    //dynamic programing
    function isInterleave2(s1: string, s2: string, s3: string): boolean {
        if (s3.length != s1.length + s2.length) {
            return false
        }

        const dp: boolean[][] = new Array()
        for (let i = 0; i <= s1.length; i++) {
            const temp: boolean[] = new Array()
            for (let j = 0; j <= s2.length; j++) {
                temp.push(false)
            }
            dp.push(temp)
        }

        for (let i = 0; i <= s1.length; i++) {
            for (let j = 0; j <= s2.length; j++) {
                if (i == 0 && j == 0) {
                    dp[i][j] = true
                } else if (i == 0) {
                    dp[i][j] = dp[i][j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1)
                } else if (j == 0) {
                    dp[i][j] = dp[i - 1][j] && s1.charAt(i - 1) == s3.charAt(i + j - 1)
                } else {
                    dp[i][j] = (dp[i - 1][j] && s1.charAt(i - 1) == s3.charAt(i + j - 1)) || (dp[i][j - 1] && s2.charAt(j - 1) == s3.charAt(i + j - 1))
                }
            }
        }
        return dp[s1.length][s2.length]
    }

    const s1 = "a", s2 = "b", s3 = "b"
    console.log(isInterleave2(s1, s2, s3))
}