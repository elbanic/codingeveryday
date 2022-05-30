namespace divide_two_integers {

    //Repeated Subtraction
    function divide(dividend: number, divisor: number): number {

        if (dividend == -2147483648 && divisor == -1) {
            return 2147483647;
        }

        let _divisor = Math.abs(divisor)
        let _dividend = Math.abs(dividend)
        const minus1 = _divisor != divisor
        const minus2 = _dividend != dividend

        let cnt = 0
        while (_dividend >= _divisor) {
            cnt++
            _dividend -= _divisor
        }
        return minus1 !== minus2 ? -cnt : cnt
    };

    //Repeated Exponential Searches
    function divide2(dividend: number, divisor: number): number {

        const HALF_INT_MIN = -1073741824

        if (dividend == -2147483648 && divisor == -1) {
            return 2147483647;
        }
        let negatives = 2
        if (dividend > 0) {
            negatives--
            dividend = -dividend
        }
        if (divisor > 0) {
            negatives--
            divisor = -divisor
        }
    
        let quotient = 0
        while (divisor >= dividend) {

            let powerOfTwo = -1;
            let value = divisor;

            while (value >= HALF_INT_MIN && value + value >= dividend) {
                value += value;
                powerOfTwo += powerOfTwo;
            }

            quotient += powerOfTwo;
            dividend -= value;
        }
    
        if (negatives != 1) {
            return -quotient;
        }
        return quotient;

    }

    const dividend = -2147483648, divisor = 4
    console.log(divide2(dividend, divisor))
}