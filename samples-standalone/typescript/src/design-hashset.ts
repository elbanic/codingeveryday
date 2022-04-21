namespace design_hashset {

    class MyHashSet {
        arr: boolean[]
        constructor() {
            this.arr = new Array();
        }

        add(key: number): void {
            this.arr[key] = true
        }

        remove(key: number): void {
            delete this.arr[key]
        }

        contains(key: number): boolean {
            return this.arr[key] != undefined
        }
    }

    var obj = new MyHashSet()
    console.log(obj.add(1))
    console.log(obj.add(2))
    console.log(obj.contains(1))
    console.log(obj.contains(3))
    console.log(obj.add(2))
    console.log(obj.contains(2))
    console.log(obj.remove(2))
    console.log(obj.contains(2))
}