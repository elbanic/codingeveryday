namespace flatten_nested_list_iterator {

    // This is the interface that allows for creating nested lists.
    // You should not implement it, or speculate about its implementation
    class NestedInteger {
        constructor(value?: number) {

        };

        isInteger(): boolean {
            return false
        };

        getInteger(): number | null {
            return null
        };

        setInteger(value: number) {

        };

        add(elem: NestedInteger) {

        };

        getList(): NestedInteger[] {
            return null
        };
    };

    class NestedIterator {
        ints: number[]
        pos: number
        constructor(nestedList: NestedInteger[]) {
            this.ints = new Array()
            this.pos = 0
            this.flatten(nestedList)
        }

        flatten(nestedList: NestedInteger[]) {
            for (let i of nestedList) {
                if (i.isInteger()) {
                    this.ints.push(i.getInteger())
                } else {
                    this.flatten(i.getList())
                }
            }
        }

        hasNext(): boolean {
            return this.pos < this.ints.length
        }

        next(): number {
            if (!this.hasNext()) return null
            return this.ints[this.pos++]
        }
    }

}