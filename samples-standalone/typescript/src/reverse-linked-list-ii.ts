namespace reverse_linked_list_ii {

    class ListNode {
        val: number
        next: ListNode | null
        constructor(val?: number, next?: ListNode | null) {
            this.val = (val === undefined ? 0 : val)
            this.next = (next === undefined ? null : next)
        }
    }

    function createList(nums: number[]): ListNode {
        const head = new ListNode(nums[0])
        let curr = head
        for (let i = 1; i < nums.length; i++) {
            curr.next = new ListNode(nums[i])
            curr = curr.next
        }
        return head
    }

    function printList(head: ListNode) {
        while (head) {
            console.log(head.val)
            head = head.next
        }
    }

    function reverseBetween(head: ListNode | null, left: number, right: number): ListNode | null {
        if (left == 0) {
            return null
        }

        let prev = head, post = head, curr = head
        let i = 0, mid = left == 1

        const arr: number[] = new Array()
        while (curr) {
            if (i == left-1) { 
                prev = curr
                mid = true
            }
            if (mid == true) {
                arr.push(curr.val)
            }
            if (i == right-1) {
                post = curr
                break
            }
            curr = curr.next
            i++            
        }
        
        while (arr.length > 0) {
            prev.val = arr.pop()
            prev = prev.next
        }
        return head
    };

    const head = [1,2,3,4,5], l = 2, r = 4
    printList(reverseBetween(createList(head), l, r))
}