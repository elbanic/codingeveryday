/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function (stones) {
    if (stones.length ==1) {
        return stones[0];
    }

    const top1 = Math.max(...stones);
    const id1 = stones.findIndex(x => x  == top1);
    stones.splice(id1, 1);

    const top2 = Math.max(...stones);
    const id2 = stones.findIndex(x => x == top2);
    stones.splice(id2, 1);
    stones.push(Math.abs(top1-top2))

    return lastStoneWeight(stones);
};



stones = [1, 3];
console.log(lastStoneWeight(stones))
