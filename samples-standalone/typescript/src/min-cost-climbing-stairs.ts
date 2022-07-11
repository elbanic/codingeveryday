namespace min_cost_climbing_stairs {

    function minCostClimbingStairs(cost: number[]): number {

        let minimumCost: number[] = new Array()
        for (let i = 0; i <= cost.length; i++) {
            minimumCost.push(0)
        }

        for (let i = 2; i < minimumCost.length; i++) {
            const takeOneStep = minimumCost[i - 1] + cost[i - 1]
            const takeTwoStep = minimumCost[i - 2] + cost[i - 2]
            minimumCost[i] = Math.min(takeOneStep, takeTwoStep)
        }
        return minimumCost.pop()
    };


    const cost = [10, 15, 20]
    console.log(minCostClimbingStairs(cost))
}