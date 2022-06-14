namespace delete_operation_for_two_stArray {

    function minDistance(word1: string, word2: string): number {
        const memo: number[][] = new Array()
        for (let i=0; i<=word1.length; i++) {
            const row: number[] = new Array()
            for (let j=0; j<=word2.length; j++) {
                row.push(0)
            }
            memo.push(row)
        }
        return word1.length + word2.length - 2 * lcs(word1, word2, word1.length, word2.length, memo)
    };

    function lcs(s1: string, s2: string, m :number,  n: number, memo: number[][]): number{
        if (m == 0 || n == 0)
            return 0;
        if (memo[m][n] > 0)
            return memo[m][n];
        if (s1.charAt(m - 1) == s2.charAt(n - 1))
            memo[m][n] = 1 + lcs(s1, s2, m - 1, n - 1, memo);
        else
            memo[m][n] = Math.max(lcs(s1, s2, m, n - 1, memo), lcs(s1, s2, m - 1, n, memo));
        return memo[m][n];
    }

    const word1 = "leetcode", word2 = "etco"
    console.log(minDistance(word1, word2))
}