/**
 * @param {string[]} ops
 * @return {number}
 */
var calPoints = function (ops) {

    let score = 0;
    let stack = [];

    for (let op of ops) {
        switch (op) {
            case "+":
                let num1 = stack[stack.length-1];
                let num2 = stack[stack.length-2];
                stack.push(parseInt(num1) + parseInt(num2));
                break;
            case "D":
                let cur = stack[stack.length-1];
                stack.push(cur * 2);
                break;
            case "C":
                stack.pop();
                break;
            default:
                stack.push(parseInt(op));
        }
    }

    return stack.reduce((prev, curr) => prev + curr, 0);
};


ops = ["5","2","C","D","+"];
console.log(calPoints(ops));