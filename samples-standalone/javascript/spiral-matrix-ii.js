/**
 * @param {number} n
 * @return {number[][]}
 */
const Right = 0
const Down = 1
const Left = 2
const Up = 3
var generateMatrix = function (n) {

    const arr = new Array();
    const visited = new Array();
    for (let i = 0; i < n; i++) {
        arr[i] = new Array();
        visited[i] = new Array();
        for (let j = 0; j < n; j++) {
            visited[i].push(false);
            arr[i].push(0);
        }
    }

    insertVal(1, 0, 0, arr, visited, Right)
    return arr;
};

var insertVal = function (n, i, j, arr, visited, dir) {
    if (visited[i][j]) {
        return
    }

    while (i >= 0 && i < arr.length && j >= 0 && j < arr.length && !visited[i][j]) {
        arr[i][j] = n;
        visited[i][j] = true;
        n++;
        switch (dir) {
            case Right:
                j++;
                break;
            case Down:
                i++;
                break;
            case Up:
                i--;
                break;
            case Left:
                j--;
                break;
        }
    }
    const nextDir = (dir + 1) % 4;
    switch (nextDir) {
        case Right:
            if (i >= 0 && i < arr.length && j >= 0 && j < arr.length && visited[i][j]) {
                i++;                
            }
            j++;
            break;
        case Down:
            if (i >= 0 && i < arr.length && j >= 0 && j < arr.length && visited[i][j]) {
                j--;
            }
            i++;
            break;
        case Up:
            if (i >= 0 && i < arr.length && j >= 0 && j < arr.length && visited[i][j]) {
                j++;
            }
            i--;
            break;
        case Left:
            if (i >= 0 && i < arr.length && j >= 0 && j < arr.length && visited[i][j]) {
                i--;
            }
            j--;
            break;
    }
    i = Math.max(Math.min(arr.length - 1, i), 0);
    j = Math.max(Math.min(arr.length - 1, j), 0);
    insertVal(n, i, j, arr, visited, nextDir);
};


console.log(generateMatrix(1))