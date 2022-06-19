

namespace search_suggestions_system {

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

    class Trie {
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

            pCrawl.isEndOfWord = true;
        }

        searchLike(key: string): string[] {
            let index: number
            let pCrawl = this.root;

            for (let level = 0; level < key.length; level++) {
                index = key[level].charCodeAt(0) - 'a'.charCodeAt(0);
                if (pCrawl.children[index] == null) {
                    return new Array()
                }
                pCrawl = pCrawl.children[index];
            }

            let output: string[] = new Array()
            let currKey = key

            while (pCrawl) {
                if (pCrawl.isEndOfWord) {
                    output.push(currKey)
                }
                pCrawl.children
            }
            return output
        }
    }

    function suggestedProducts(products: string[], searchWord: string): string[][] {

        const trie = new Trie()
        for (let prod of products) {
            trie.insert(prod)
        }

        let word: string
        let output: string[][] = new Array()
        for (let i = 0; i < searchWord.length; i++) {
            word += searchWord[i]
            const found = trie.searchLike(word) 
            if (found.length == 0) { break }
            output.push(found)
        }
        return output
    };

    const products = ["mobile", "mouse", "moneypot", "monitor", "mousepad"], searchWord = "mouse"
    console.log(suggestedProducts(products, searchWord))
}