namespace non_decreasing_array {
    function checkPossibility(nums: number[]): boolean {
        let a: number[], b: number[]
        let aRes = true, bRes = true
        for (let i = 0; i < nums.length - 1; i++) {
            if (nums[i] > nums[i + 1]) {
                a = [...nums]
                b = [...nums]
                a.splice(i, 1)
                b.splice(i + 1, 1)
                break
            }
        }
        for (let i = 0; i < a.length - 1; i++) {
            if (a[i] > a[i + 1]) {
                aRes = false
                break
            }
        }

        for (let i = 0; i < b.length - 1; i++) {
            if (b[i] > b[i + 1]) {
                bRes = false
                break
            }
        }
        return aRes || bRes
    };

    const nums = [4, 2, 3]
    console.log(checkPossibility(nums))

    const nums2 = [4, 2, 1]
    console.log(checkPossibility(nums2))
}