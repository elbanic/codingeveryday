
namespace design_underground_system {

    class Passenger {
        endStation: string
        endTime: number
        constructor(public id: number, public startStation: string, public startTime: number) { }

        setEndStation(endStation: string, endTime: number): void {
            this.endStation = endStation
            this.endTime = endTime
        }

        getRoute(): Route {
            return new Route(this.startStation, this.endStation, this.endTime - this.startTime)
        }
    }

    class Route {
        constructor(public startStation: string, public endStation: string, public time: number) { }
    }

    class UndergroundSystem {

        passengers: Map<number, Passenger>
        routes: Route[]

        constructor() {
            this.passengers = new Map<number, Passenger>()
            this.routes = new Array()
        }

        checkIn(id: number, stationName: string, t: number): void {
            this.passengers[id] = new Passenger(id, stationName, t)
        }

        checkOut(id: number, stationName: string, t: number): void {
            if (this.passengers[id]) {
                this.passengers[id].setEndStation(stationName, t)
                this.routes.push(this.passengers[id].getRoute())

                delete this.passengers[id]
            }
        }

        getAverageTime(startStation: string, endStation: string): number {
            let sum = 0
            let cnt = 0
            for (let route of this.routes) {
                if (route.startStation == startStation && route.endStation == endStation) {
                    sum += route.time
                    cnt++                    
                }
            }
            if (cnt > 0) {
                return sum / cnt
            }
            return 0
        }
    }

    var obj = new UndergroundSystem()

    console.log(obj.checkIn(45, "Leyton", 3))
    console.log(obj.checkIn(32, "Paradise", 8))
    console.log(obj.checkIn(27, "Leyton", 10))
    console.log(obj.checkOut(45, "Waterloo", 15))
    console.log(obj.checkOut(27, "Waterloo", 20))
    console.log(obj.checkOut(32, "Cambridge", 22))
    console.log(obj.getAverageTime("Paradise", "Cambridge"))
    console.log(obj.getAverageTime("Leyton", "Waterloo"))
    console.log(obj.checkIn(10, "Leyton", 24))
    console.log(obj.getAverageTime("Leyton", "Waterloo"))
    console.log(obj.checkOut(10, "Waterloo", 38))
    console.log(obj.getAverageTime("Leyton", "Waterloo"))

}