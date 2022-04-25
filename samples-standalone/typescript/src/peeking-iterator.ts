namespace peeking_iterator {

    class Iterator {
        pointer: number
        constructor(public nums: number[] | null) {
            this.pointer = -1
        }

        hasNext(): boolean {
            return this.pointer < this.nums.length - 1
        }

        next(): number {
            this.pointer = Math.min(this.pointer + 1, this.nums.length)
            return this.nums[this.pointer]
        }
    }

    // class PeekingIterator {
    //     nums: number[]
    //     pointer: number
    //     constructor(iterator: Iterator) {
    //         this.nums = new Array()
    //         while (iterator.hasNext()) {
    //             this.nums.push(iterator.next())
    //         }
    //         this.pointer = -1
    //     }

    //     peek(): number {
    //         return this.nums[this.pointer+1]
    //     }

    //     hasNext(): boolean {
    //         return this.pointer < this.nums.length - 1
    //     }

    //     next(): number {
    //         this.pointer = Math.min(this.pointer + 1, this.nums.length)
    //         return this.nums[this.pointer]
    //     }
    // }

    class PeekingIterator {
        iter: Iterator
        pointer: number
        peeked: number
        constructor(iterator: Iterator) {
            this.iter = iterator
            this.pointer = -1
            this.peeked = undefined
        }

        peek(): number {
            if (this.peeked == undefined) {
                if (this.iter.hasNext()) {
                    this.peeked = this.iter.next()
                }
            }
            return this.peeked
        }

        hasNext(): boolean {            
            return this.peeked != undefined || this.iter.hasNext()
        }

        next(): number {
            if (this.peeked != undefined) {
                const temp = this.peeked
                this.peeked = undefined
                return temp
            }
            return this.iter.next()
        }
    }

    let iterator = new Iterator([1, 2, 3])
    let obj = new PeekingIterator(iterator)

    console.log(obj.next())
    console.log(obj.peek())
    console.log(obj.next())
    console.log(obj.next())
    console.log(obj.hasNext())
}