namespace queue_reconstruction_by_height {

    function reconstructQueue(people: number[][]): number[][] {

        people.sort((a, b) => {
            if (b[0] > a[0]) {
                return 1;
            } else if (b[0] < a[0]) {
                return -1;
            }

            if (b[1] > a[1]) {
                return -1;
            } else if (b[1] < a[1]) {
                return 1;
            } else {
                return 0;
            }
        });

        const output: number[][] = new Array()
        for (let i = 0; i < people.length; i++) {
            output.splice(people[i][1], 0, people[i]);
        }

        return output
    };

    const people = [[7, 0], [4, 4], [7, 1], [5, 0], [6, 1], [5, 2]]
    console.log(reconstructQueue(people))

}