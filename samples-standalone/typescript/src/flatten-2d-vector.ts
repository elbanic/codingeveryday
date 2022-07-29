namespace flatten_2d_vector {
    class Vector2D {
        i: number
        j: number
        vec: number[][]
        constructor(vec: number[][]) {
            this.vec = vec, this.i = 0, this.j = 0
        }

        next(): number {
            if (this.hasNext()) {
                const i = this.i, j = this.j
                if (this.j < this.vec[this.i].length - 1) {
                    this.j++
                } else {
                    this.i++
                    this.j=0
                }
                return this.vec[i][j]
            } else {
                return null
            }
        }

        hasNext(): boolean {
            if (this.i < this.vec.length) {
                if (this.j >= this.vec[this.i].length) {
                    this.i++
                    this.j = 0
                }
                if (this.vec[this.i][this.j] != null) {
                    return true
                }
                return this.hasNext()
            } else {
                return false
            }
        }
    }

    const vector = [[],[]]
    var obj = new Vector2D(vector)
    console.log(obj.next())
    console.log(obj.hasNext())
    console.log(obj.next())
    console.log(obj.next())
    console.log(obj.next())
    console.log(obj.next())
    console.log(obj.hasNext())
}