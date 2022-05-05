namespace implement_stack_using_queues {
    class MyStack {
        queue: number[]
        constructor() {
            this.queue = new Array()
        }

        push(x: number): void {
            this.queue.push(x)
        }

        pop(): number {
            return this.queue.pop()
        }

        top(): number {
            return this.queue[this.queue.length-1]
        }

        empty(): boolean {
            return this.queue.length == 0
        }
    }

    const obj = new MyStack()
    console.log(obj.push(1))
    console.log(obj.push(2))
    console.log(obj.top())
    console.log(obj.pop())
    console.log(obj.empty())

}