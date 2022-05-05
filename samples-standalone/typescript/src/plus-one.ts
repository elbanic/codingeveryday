namespace plus_one {
    function plusOne(digits: number[]): number[] {

        if (digits.length == 0) {
            return new Array(1)
        }

        if (digits[digits.length-1] < 9) {
            digits[digits.length-1] += 1
            return digits
        }

        let carry: number = 1
        const newDigits: number[] = new Array()
        for (let i = digits.length-1; i>=0; i--) {
            let remain = (digits[i] + carry) % 10
            newDigits.unshift(remain)
            carry = Math.floor((digits[i] + carry) / 10)
            
        }
        if (carry == 1) {
            newDigits.unshift(carry)
        }
        return newDigits
    };

    const digits = [9,8,9]
    console.log(plusOne(digits))
}