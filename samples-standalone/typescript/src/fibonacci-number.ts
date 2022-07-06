namespace fibonacci_number {
 
    function fib(n: number): number {

        if (n == 0 || n == 1) {
            return n
        }

        return fib(n-1) + fib(n-2)
    };

    const n = 4
    console.log(fib(4))
}