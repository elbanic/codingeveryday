let ALPHABET_SIZE = 26;

class TrieNode {
    isEndOfWord: boolean
    children: TrieNode[]
    constructor() {
        this.isEndOfWord = false
        this.children = new Array()
        for (let i = 0; i < ALPHABET_SIZE; i++) {
            this.children.push(null)
        }
    }
}

export class Trie {
    root: TrieNode
    constructor() { 
        this.root = new TrieNode()
    }
    insert(key: string) {
        let index: number
        let pCrawl = this.root

        for (let level = 0; level < key.length; level++) {
            index = key[level].charCodeAt(0) - "a".charCodeAt(0)
            if (pCrawl.children[index] == null) {
                pCrawl.children[index] = new TrieNode()
            }
            pCrawl = pCrawl.children[index]
        }

        pCrawl.isEndOfWord = true
    }

    search(key: string): boolean {
        let index: number
        let pCrawl = this.root

        for (let level = 0; level < key.length; level++) {
            index = key[level].charCodeAt(0) - 'a'.charCodeAt(0)
            if (pCrawl.children[index] == null) {
                return false
            }
            pCrawl = pCrawl.children[index]
        }
        return pCrawl.isEndOfWord
    }
}

