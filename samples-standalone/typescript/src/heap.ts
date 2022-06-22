export class MinimumHeap {
    heap: number[] = new Array();
    constructor() { }

    peak(): number{
        return this.heap[0]
    }

    insert(x: number){
        var contain = false
        for (var i = 0; i < this.heap.length; i++) {
            if (this.heap[i] > x) {                
                this.heap.splice(i, 0, x)
                contain = true
                break
            }
        }
        if (!contain) {
            this.heap.push(x);
        }
    }

    pop(): number {
        if (this.heap.length == 0) {
            return null
        }
        return this.heap.shift()
    }
}

export class MaximumHeap {
    heap: number[] = new Array();
    constructor() { }

    peak(): number{
        return this.heap[0]
    }

    insert(x: number){
        var contain = false
        for (var i = 0; i < this.heap.length; i++) {
            if (this.heap[i] < x) {                
                this.heap.splice(i, 0, x)
                contain = true
                break
            }
        }
        if (!contain) {
            this.heap.push(x);
        }
    }

    pop(): number {
        if (this.heap.length == 0) {
            return null
        }
        return this.heap.shift()
    }
}