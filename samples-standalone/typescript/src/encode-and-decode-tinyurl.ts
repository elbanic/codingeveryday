

namespace encode_and_decode_tinyurl {

    const urls = new Map<string, string>();
    let inc = 0

    function encode(longUrl: string): string {

        if (urls.get(longUrl)) {
            return urls.get(longUrl)
        } else {
            urls[longUrl] = "http://tinyurl.com/" + inc.toString()
            urls["http://tinyurl.com/" + inc.toString()] = longUrl
            return "http://tinyurl.com/" + inc.toString()
        }
    };

    function decode(shortUrl: string): string {
        return urls[shortUrl]
    };

    let url = "https://leetcode.com/problems/design-tinyurl"

    let tiny = encode(url)
    let ans = decode(tiny)
    console.log(tiny)
    console.log(ans)

}