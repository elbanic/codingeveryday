import { map } from "ramda"


namespace max_number_of_k_sum_pairs {
    function maxOperations(nums: number[], k: number): number {

        const m: Map<number, number> = new Map<number, number>()

        for (let n of nums) {
            m.set(n, m.has(n) ? m.get(n) + 1 : 1)
        }

        let cnt = 0
        for (let [key, value] of m) {
            if (m.has(k - key)) {

                if (key == k-key) {
                    m.set(key, value % 2)
                    cnt += Math.floor(value / 2)
                    continue
                }

                const value1 = value, value2 = m.get(k - key)
                if (value >= 2 && value2 >= 2) {
                    if (value > value2) {
                        m.set(key, value - value2)
                        m.delete(k - key)
                    } else {
                        m.set(k - key, value2 - value)
                        m.delete(key)
                    }
                    cnt+=Math.min(value1, value2)
                } else {
                    if (value > 1) {
                        m.set(key, value - 1)
                    } else {
                        m.delete(key)
                    }
                    if (value2 > 1) {
                        m.set(k - key, value2 - 1)
                    } else {
                        m.delete(k - key)
                    }
                    cnt++
                }
            }
        }
        return cnt
    };
    
    const nums = [2,5,4,4,1,3,4,4,1,4,4,1,2,1,2,2,3,2,4,2], k = 3
    console.log(maxOperations(nums, k))
}
