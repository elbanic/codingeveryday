namespace two_sum_ii_input_array_is_sorted {

    function twoSum(numbers: number[], target: number): number[] {

        const map = new Map<number, number>()
        for (let i =0;i<numbers.length; i++) {
            map.set(numbers[i], i)
        }
        for (let i = 0; i< numbers.length; i++){
            if (map.has(target - numbers[i])) {
                return new Array(i+1, map.get(target - numbers[i])+1)
            }
        }
        return null
    };

    const numbers = [2,7,11,15], target = 9
    console.log(twoSum(numbers, target))
}