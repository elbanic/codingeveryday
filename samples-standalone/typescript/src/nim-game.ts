namespace nim_game {
    function canWinNim(n: number): boolean {
        return (n % 3) != 1
    };

    console.log(canWinNim(4))
}