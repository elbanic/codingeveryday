

namespace search_suggestions_system {

    let ALPHABET_SIZE = 26;
    let ALPHABETS = "abcdefghijklmnopqrstuvwxyz".split("")

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

        findAllWords(trie: TrieNode, key: string, output: string[]): string[] {
            if (trie.isEndOfWord) {
                output.push(key)
            }
            if (output.length >= 3) {
                return output
            }
            for (let child =0; child<trie.children.length; child++) {
                if (trie.children[child]) {
                    output = this.findAllWords(trie.children[child], key + ALPHABETS[child], output)
                }
            }
            return output
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

            return this.findAllWords(pCrawl, key, new Array())
        }
    }

    function suggestedProducts(products: string[], searchWord: string): string[][] {

        const trie = new Trie()
        for (let prod of products) {
            trie.insert(prod)
        }

        let word: string = ""
        let output: string[][] = new Array()
        for (let i = 0; i < searchWord.length; i++) {
            word += searchWord[i]
            const found = trie.searchLike(word)
            output.push(found) 
        }
        return output
    };

    const products = ["mobile", "mouse", "moneypot", "monitor", "mousepad"], searchWord = "mouse"
    console.log(suggestedProducts(products, searchWord))
}