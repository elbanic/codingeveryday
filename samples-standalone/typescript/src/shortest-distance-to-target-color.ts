namespace shortest_distance_to_target_color {
    function shortestDistanceColor(colors: number[], queries: number[][]): number[] {

        const dist: number[] = new Array()
        while (queries.length) {
            const cur = queries.shift()

            if (bsearch([...colors], cur[1]) == -1) {
                dist.push(-1)
                continue
            }

            const left = colors.slice(0, cur[0])
            const right = colors.slice(cur[0])
            const foundl = left.reverse().findIndex(x => x === cur[1])
            const foundr = right.findIndex(x => x === cur[1])
            if (foundl == -1 && foundr == -1) {
                dist.push(-1)
                continue
            }
            if (foundl == -1) { dist.push(foundr) }
            else if (foundr == -1) { dist.push(foundl+1) }
            else { dist.push(Math.min(foundl+1, foundr)) }
        }
        return dist
    };
    
    function bsearch(nums: number[], target: number): number {

        nums.sort((a,b) => a - b)
        let left = 0, right = nums.length - 1
        while (left < right) {
            const mid = Math.floor(left + (right - left) / 2)
            if (nums[mid] == target) {
                return mid
            } else if (nums[mid] < target) {
                left = mid+1
            } else {
                right = mid-1
            }
        }
        return -1
    }

    console.log(bsearch([1,1,2,1,3,2,2,3,3], 1))
    console.log(bsearch([1,1,2,1,3,2,2,3,3], 4))
    const colors = [1,1,2,1,3,2,2,3,3], queries = [[1,3],[2,2],[6,1]]
    console.log(shortestDistanceColor(colors, queries))
}