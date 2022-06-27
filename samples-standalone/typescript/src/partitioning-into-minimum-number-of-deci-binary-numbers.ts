namespace partitioning_into_minimum_number_of_deci_binary_numbers {

    function minPartitions(n: string): number {

        let max = 0
        const nums = n.split("")

        for (let n of nums) {
            if (max < +n){
                max = +n
            }
        }
        return max
    };

    const n = "27346209830709182346"
    console.log(minPartitions(n))
}