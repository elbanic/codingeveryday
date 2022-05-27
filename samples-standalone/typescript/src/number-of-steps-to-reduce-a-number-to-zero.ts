namespace number_of_steps_to_reduce_a_number_to_zero{

    function numberOfSteps(num: number): number {

        let count = 0
        while (num > 0) {
            if (num % 2 == 0) {
                num /= 2
            } else {
                num -= 1
            }
            count++
        }
        return count
    };

    const num = 14
    console.log(numberOfSteps(num))
}