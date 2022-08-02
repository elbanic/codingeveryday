namespace kth_smallest_element_in_a_sorted_matrix {

    function kthSmallest(matrix: number[][], k: number): number {
        //flatten
        return matrix.reduce((accumulator, value) => accumulator.concat(value), []).sort((a, b) => a - b)[k - 1]
    };

    const matrix = [[1, 5, 9], [10, 11, 13], [12, 13, 15]], k = 8
    console.log(kthSmallest(matrix, k))

}