namespace jump_game_vi {
    function maxResult(nums: number[], k: number): number {
        let sum = nums[0]
        for (let i = 1; i < nums.length; i++) {
            let max = nums[i]
            const orgI = i
            for (let j = orgI; j < orgI + k; j++) {
                if (j >= nums.length) {
                    break
                }
                if (max < nums[j]) {
                    max = nums[j]
                    i = j
                }
            }
            sum += max
        }
        return sum
    };

    function maxResult2(nums: number[], k: number): number {

        const n = nums.length
        const score: number[] = new Array()
        for (let i = 0; i < n; i++) {
            score.push(0)
        }
        const queue: number[] = new Array()
        queue.push(0)

        for (let i = 1; i < n; i++) {
            while (queue.length > 0 && queue[0] < i - k) {
                queue.shift()
            }
            score[i] = score[queue[0]] + nums[i]
            while (queue.length > 0 && score[i] >= score[queue[queue.length - 1]]) {
                queue.pop()
            }
            queue.push(i)
        }
        return score[n - 1]
    }

    const nums = [10, -5, -2, 4, 0, 3], k = 3
    console.log(maxResult2(nums, k))
}