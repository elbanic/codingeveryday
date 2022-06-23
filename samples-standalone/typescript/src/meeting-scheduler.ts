namespace meeting_scheduler {

    function minAvailableDuration(slots1: number[][], slots2: number[][], duration: number): number[] {

        slots1.sort((a,b) => a[0]-b[0])
        slots2.sort((a,b) => a[0]-b[0])
        for (let slot1 of slots1) {
            for (let slot2 of slots2) {
                if ((slot1[0] >= slot2[0] && slot1[0] <= slot2[1]) ||
                    (slot1[1] >= slot2[0] && slot1[1] <= slot2[1]) ||
                    (slot2[0] >= slot1[0] && slot2[0] <= slot1[1]) ||
                    (slot2[1] >= slot1[0] && slot2[1] <= slot1[1])
                ) {
                    const left = Math.max(slot1[0], slot2[0])
                    const right = Math.min(slot1[1], slot2[1])
                    if (right - left >= duration) {
                        return new Array(left, left + duration)
                    }
                }
            }
        }
        return new Array()
    };

    const slots1 = [[10, 50], [60, 120], [140, 210]], slots2 = [[0, 15], [60, 70]], duration = 12
    console.log(minAvailableDuration(slots1, slots2, duration))
}