namespace length_of_last_word {
 
    function lengthOfLastWord(s: string): number {

        const splitted = s.split(" ")
        const filtered = splitted.filter(x => x.length != 0)
        return filtered[filtered.length-1].length
    };

    const s = "   fly me   to   the moon  "
    console.log(lengthOfLastWord(s))
}