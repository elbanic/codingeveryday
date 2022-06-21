namespace furthest_building_you_can_reach {
    function furthestBuilding(heights: number[], bricks: number, ladders: number): number {
        return helper(heights, 0, 0, bricks, ladders)
    };

    function helper(heights: number[], index: number, max: number, bricks: number, ladders: number): number {

        if (index == heights.length - 1) {
            return index
        }

        if (heights[index + 1] - heights[index] < 0) {
            return helper(heights, index + 1, max, bricks, ladders)
        }

        if (bricks >= heights[index + 1] - heights[index]) {
            max = Math.max(max, helper(heights, index + 1, max, bricks - (heights[index + 1] - heights[index]), ladders))
            if (ladders > 0) {
                max = Math.max(max, helper(heights, index + 1, max, bricks, ladders - 1))
            }
        } else if (ladders > 0) {
            max = Math.max(max, helper(heights, index + 1, max, bricks, ladders - 1))
        }
        return Math.max(max, index)
    }

    const heights = [14,3,19,3], bricks = 17, ladders = 0
    console.log(furthestBuilding(heights, bricks, ladders))
}