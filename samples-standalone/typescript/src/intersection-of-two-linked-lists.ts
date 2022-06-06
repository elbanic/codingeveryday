namespace intersection_of_two_linked_lists {

    class ListNode {
        val: number
        next: ListNode | null
        constructor(val?: number, next?: ListNode | null) {
            this.val = (val === undefined ? 0 : val)
            this.next = (next === undefined ? null : next)
        }
        print() {
            if (this.next) {
                process.stdout.write(this.val + ' -> ');
                this.next.print()
            } else {
                process.stdout.write(''+ this.val + '\n')
            }
        }
    }

    function getIntersectionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {

        const map: Map<number, ListNode> = new Map<number, ListNode>()
        let curr = headA
        while (curr) {
            map.set(curr.val, curr)
            curr = curr.next
        }
        curr = headB
        while (curr) {
            if (map.has(curr.val)){
                return map.get(curr.val)
            }
            curr = curr.next
        }
        return null
    };

    function createList(nums: number[]): ListNode {
        const head = new ListNode()
        let curr = head
        for (let n of nums){
            curr.next = new ListNode(n)
            curr = curr.next
        }
        return head.next
    }

    const listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
    const lista = createList(listA); lista.print()
    const listb = createList(listB); listb.print()
    const inters = getIntersectionNode(lista, listb)
    if (inters != null) inters.print()
}