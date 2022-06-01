
namespace running_sum_of_1d_array {

    function runningSum(nums: number[]): number[] {
        const sums: number[] = new Array()
        const final = nums.reduce((prev, curr) => {sums.push(prev); return prev+curr})
        sums.push(final)
        return sums
    };

    const nums = [1,2,3,4]
    console.log(runningSum(nums))
}