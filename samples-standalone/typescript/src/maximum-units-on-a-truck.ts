namespace maximum_units_on_a_truck {

    function maximumUnits(boxTypes: number[][], truckSize: number): number {

        boxTypes.sort(function (a, b) {
            return b[1] - a[1]
        });

        let sum: number = 0
        let i = 0;
        let count: number = 0

        while (count <= truckSize && i < boxTypes.length) {
            if (count+boxTypes[i][0] <= truckSize) {
                sum += boxTypes[i][0] * boxTypes[i][1]
                count += boxTypes[i][0]
            } else {
                const remain = truckSize - count
                sum += remain * boxTypes[i][1]
                break
            }
            i++
        }
        return sum
    }

    let boxTypes = [[1,3],[5,5],[2,5],[4,2],[4,1],[3,1],[2,2],[1,3],[2,5],[3,2]], truckSize = 35
    console.log(maximumUnits(boxTypes, truckSize))
}