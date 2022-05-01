

namespace backspace_string_compare {
    function backspaceCompare(s: string, t: string): boolean {

        let stackS: string[] = new Array()
        let stackT: string[] = new Array()

        for (let i of s) {
            if (i != "#") {
                stackS.push(i)
            } else {
                stackS.pop()
            }
        }

        for (let i of t) {
            if (i != "#") {
                stackT.push(i)
            } else {
                stackT.pop()
            }
        }

        if (stackS.length != stackT.length) {
            return false
        }

        while (stackS.length > 0) {
            if (stackS.pop() != stackT.pop()) {
                return false
            }   
        }
        return true
    };

    let s = "a#c", t = "b"

    console.log(backspaceCompare(s, t))
}