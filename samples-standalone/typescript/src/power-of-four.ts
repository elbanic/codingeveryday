namespace power_of_four {
 
    function isPowerOfFour(n: number): boolean {
        if (n == 0) {
            return false
        }

        while (n > 1){
            n /= 4
            if (n != Math.floor(n)) {
                return false
            }
        }
        return n == 1
    };

    console.log(isPowerOfFour(16))
    console.log(isPowerOfFour(5))
    console.log(isPowerOfFour(1))

    console.log(isPowerOfFour(4))
}