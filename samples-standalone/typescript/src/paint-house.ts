namespace paint_house {

    function minCost(costs: number[][]): number {

        for (let n = costs.length - 2; n >= 0; n--) {
            costs[n][0] += Math.min(costs[n + 1][1], costs[n + 1][2])
            costs[n][1] += Math.min(costs[n + 1][0], costs[n + 1][2])
            costs[n][2] += Math.min(costs[n + 1][0], costs[n + 1][1])
        }
        if (costs.length == 0) {
            return 0
        }

        return Math.min(Math.min(costs[0][0], costs[0][1]), costs[0][2])
    };

    const costs = [[17, 2, 17], [16, 16, 5], [14, 3, 19]]
    console.log(minCost(costs))
}