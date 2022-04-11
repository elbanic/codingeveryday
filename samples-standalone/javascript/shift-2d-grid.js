/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var shiftGrid = function (grid, k) {

    let n = grid.length;
    let m = grid[0].length;

    const arr = new Array();
    for (let i of grid) {
        for (let j of i) {
            arr.push(j);
        }
    }
    
    k %= arr.length;
    for (let i = 0; i<k; i++) {
        let s = arr.pop();
        arr.splice(0, 0, s);
    }

    for (let i =0; i<n; i++) {
        for (let j =0; j<m; j++) {
            grid[i][j] = arr.shift();
        }
    }
    return grid
};

let grid = [[1,2,3],[4,5,6],[7,8,9]];
let k = 1;
console.log(shiftGrid(grid, k));