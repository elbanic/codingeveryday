namespace reverse_bits {

    function reverseBits(n: number): number {
        let ret = 0, power = 24
        const cache = new Map<number, number>()
        
        while (n != 0) {
            ret += reverseByte(n & 0xff, cache) << power
            n =n >> 8
            power -= 8
        }
        return ret
    };

    function reverseByte(byte: number, cache: Map<number, number>): number {
        if (cache.has(byte)) {
            return cache[byte]
        }
        const value = (byte * 0x0202020202 & 0x010884422010) % 1023
        cache[byte] = value
        return value
    }

    const n = "00000010100101000001111010011100"

    console.log(reverseBits(parseInt(n, 2)))
}