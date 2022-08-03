namespace my_calendar_i {

    class MyCalendar {

        calendar: number[][]
        constructor() {
            this.calendar = new Array()
        }

        book(start: number, end: number): boolean {

            for (let arr of this.calendar) {
                if (arr[0] < end && start < arr[1]) {
                    return false
                }
            }
            this.calendar.push([start, end])
            return true
        }
    }

    const obj = new MyCalendar()
    console.log(obj.book(10, 20))
    console.log(obj.book(15, 25))
    console.log(obj.book(20, 30))
}