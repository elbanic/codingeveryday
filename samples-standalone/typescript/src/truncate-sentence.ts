namespace truncate_sentence {
    function truncateSentence(s: string, k: number): string {
        return s.split(" ").slice(0, k).join(" ")
    };

    const s = "Hello how are you Contestant", k = 4
    console.log(truncateSentence(s, k))
}