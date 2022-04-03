/**
 * @param {string} s
 * @return {string[]}
 */
var expand = function (s) {
    return findAllWords(s, 0);
};

var storeFirstOptions = function (s, startPos, firstOptions) {

    if (s.charAt(startPos) != '{') {
        firstOptions.push(s.charAt(startPos));
    } else {
        while (s.charAt(startPos) != '}') {
            if (s.charAt(startPos) >= 'a' && s.charAt(startPos) <= 'z') {
                firstOptions.push(s.charAt(startPos));
            }
            startPos++;
        }
        firstOptions.sort();
    }
    return startPos + 1;
};

var findAllWords = function (s, startPos) { 
    if (startPos == s.length) {
        return [''];
    }

    let firstOpstions = new Array();
    let remStringStartPos = storeFirstOptions(s, startPos, firstOpstions);
    let wordsWithRemString = findAllWords(s, remStringStartPos);
    let expandedWords = new Array();

    for (const c of firstOpstions) {
        for (const word of wordsWithRemString) {
            expandedWords.push(c + word);
        }        
    }
    return expandedWords
};

let s = "{a,b}c{d,e}f"
console.log(expand(s))
