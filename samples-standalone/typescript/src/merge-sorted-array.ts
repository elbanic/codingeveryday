namespace merge_sorted_array {

    /**
     Do not return anything, modify nums1 in-place instead.
    */
    function merge(nums1: number[], m: number, nums2: number[], n: number): void {

        if (nums1.length == 0 || nums2.length == 0) {
            return
        }

        let p1 = 0, p2 = 0
        while (p1 < nums1.length || p2 < nums2.length) {

            if (p1 >= nums1.length) {
                nums1.push(nums2[p2])
                p2++
                break
            }

            if (p2 >= nums2.length) {
                return
            }

            if (p1 >= m) {
                nums1[p1] = nums2[p2]
                p1++
                p2++
            } else if (nums1[p1] > nums2[p2]) {
                nums1.pop()
                nums1.splice(p1, 0, nums2[p2])
                p1++
                p2++
                m++
            } else {
                p1++
            }
        }
    };

    function merge2(nums1: number[], m: number, nums2: number[], n: number): void {

        if (nums1.length == 0 || nums2.length == 0) {
            return
        }

        const size = nums1.length
        for (let i = 0; i < size - m; i++) {
            nums1.pop()
            console.log(i, nums1)
        }

        nums2.forEach(x => nums1.push(x))
        nums1.sort((a, b) => a - b)
    };

    const nums1 = [1, 2, 3, 0, 0, 0], m = 3, nums2 = [2, 5, 6], n = 3
    merge2(nums1, m, nums2, n)

    console.log(nums1)
}