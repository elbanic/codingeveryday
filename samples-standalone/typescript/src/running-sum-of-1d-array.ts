
namespace running_sum_of_1d_array {

    function runningSum(nums: number[]): number[] {
        const sums: number[] = new Array()
        nums.reduce((prev, curr) => {sums.push(prev); return prev+curr})
        return sums
    };

    const nums = [1,2,3,4]
    console.log(runningSum(nums))
}