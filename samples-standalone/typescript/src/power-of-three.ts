namespace power_of_three {

    function isPowerOfThree(n: number): boolean {
        while(n > 1) {
            if (n != Math.floor(n)) {
                return false
            }
            n /= 3
        }
        return n == 1
    };

    console.log(isPowerOfThree(27))
    console.log(isPowerOfThree(0))
    console.log(isPowerOfThree(9))
}