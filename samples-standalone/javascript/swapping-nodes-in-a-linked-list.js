/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */

 var swapNodes = function (head, k) {

    let arr = new Array();
    let cur = head;
    while (cur != null) {
        arr.push(cur)
        cur = cur.next;
    }

    const temp = arr[k-1];
    arr[k-1] = arr[arr.length-k];
    arr[arr.length-k] = temp;

    if (arr.length > k && k > 1) {
        arr[k-2].next = arr[k-1];
        arr[k-1].next = arr[k];
        arr[arr.length-k-1].next = arr[arr.length-k];
        arr[arr.length-k].next = arr[arr.length-k+1];    
    } else {
        for (let i =0; i<arr.length-1; i++) {
            arr[i].next = arr[i+1];
        }
        head = arr[0];
        arr[arr.length-1].next = null;
    }

    return head
};

var createList = function (nums) {

    const head = new ListNode();
    let cur = head;
    for (n of nums) {
        cur.next = new ListNode(n);
        cur = cur.next;
    }
    return head.next;
}

function ListNode(val, next) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
}

const nums = [1];
const k = 1;

const head = createList(nums);
console.log(head);

const newHead = swapNodes(head, k);
console.log(newHead);