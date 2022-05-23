
namespace ones_and_zeroes {

    function findMaxForm(strs: string[], m: number, n: number): number {

        const memo: number[][][] = new Array()
        for (let i = 0; i<strs.length; i++) {
            const x: number[][] = new Array()
            for (let j = 0; j<=m; j++) {
                const y: number[] = new Array()
                for (let k =0; k<=n; k++) {
                    y.push(0)
                }
                x.push(y)
            }
            memo.push(x)
        }

        return calculate(strs, 0, m, n, memo)
    };

    function calculate(strs: string[], i: number, zeros: number, ones: number, memo: number[][][]): number {

        if (i == strs.length) {
            return 0
        }
        if (memo[i][zeros][ones] != 0) {
            return memo[i][zeros][ones]
        }
        const count: ZerosOnes = countZerosOnes(strs[i])
        let taken = -1

        if (zeros - count.z >= 0 && ones - count.o >= 0) {
            taken = calculate(strs, i+1, zeros - count.z, ones - count.o, memo) + 1
        }
        const notTaken = calculate(strs, i+1, zeros, ones, memo)
        memo[i][zeros][ones] = Math.max(taken, notTaken)

        return memo[i][zeros][ones]
    }
    class ZerosOnes {
        constructor(public z: number, public o: number) { }
    }

    function countZerosOnes(s: string): ZerosOnes{
        return new ZerosOnes(
                s.split('').filter(x => x == '0').length, 
                s.split('').filter(x => x == '1').length)
    }
    
    const strs = ["10","0001","111001","1","0"], m = 3, n = 4
    console.log(findMaxForm(strs, m, n))

}