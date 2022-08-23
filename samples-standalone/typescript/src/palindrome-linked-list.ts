namespace palindrome_linked_list {


    class ListNode {
        val: number
        next: ListNode | null
        constructor(val?: number, next?: ListNode | null) {
            this.val = (val === undefined ? 0 : val)
            this.next = (next === undefined ? null : next)
        }
    }

    function createList(nums: number[]): ListNode {
        if (nums.length == 0) {
            return null
        }

        const head = new ListNode()
        let cur = head
        for (let n of nums) {
            cur.next = new ListNode(n)
            cur = cur.next
        }
        return head.next
    }

    function isPalindrome(head: ListNode | null): boolean {
        if (head == null) {
            return false
        }

        const arr: ListNode[] = new Array()
        let cur = head
        while (cur) {
            arr.push(cur)
            cur = cur.next
        }

        cur = head
        for (let i = arr.length - 1; i >= arr.length / 2; i--) {
            if (cur.val != arr[i].val) {
                return false
            }
            cur = cur.next
        }
        return true
    };

    const head = [1, 2]
    const headList = createList(head)
    console.log(isPalindrome(headList))
}