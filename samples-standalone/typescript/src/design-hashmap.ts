namespace design_hashmap {
    class MyHashMap {
        data: number[]
        constructor() {
            this.data = new Array()
        }

        put(key: number, value: number): void {
            this.data[key] = value
        }

        get(key: number): number {
            if (this.data[key] != undefined) {
                return this.data[key]
            } else {
                return -1
            }
        }

        remove(key: number): void {
            delete this.data[key]
        }
    }

    var obj = new MyHashMap()

    console.log(obj.put(1,1))
    console.log(obj.put(2,2))
    console.log(obj.get(1))
    console.log(obj.get(3))
    console.log(obj.put(2,1))
    console.log(obj.get(2))
    console.log(obj.remove(2))
    console.log(obj.get(2))
}
