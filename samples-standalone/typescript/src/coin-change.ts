namespace coin_change {

    function coinChange(coins: number[], amount: number): number {
        if (amount < 1) {
            return 0
        }

        const count:number[] = new Array()
        for (let i =0; i<amount; i++) {
            count.push(0)
        }
        return helper(coins, amount, count)
    };

    function helper(coins: number[], rem: number, count: number[]): number {
        if (rem < 0) { 
            return -1 
        }
        if (rem == 0) { 
            return 0 
        }
        if (count[rem - 1] != 0) { 
            return count[rem - 1] 
        }

        let min = Number.MAX_SAFE_INTEGER
        for (let coin of coins) {
            const res = helper(coins, rem - coin, count)
            if (res >= 0 && res < min) {
                min = 1 + res
            }
        }
        count[rem - 1] = (min == Number.MAX_SAFE_INTEGER) ? -1 : min
        return count[rem - 1]
    }

    const coins = [186, 419, 83, 408], amount = 6249
    console.log(coinChange(coins, amount))
}