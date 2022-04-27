namespace number_of_days_in_a_month {

    function numberOfDays(year: number, month: number): number {
        return new Date(year, month, 0).getDate()
    };
    
    console.log(numberOfDays(1992, 7))
}
