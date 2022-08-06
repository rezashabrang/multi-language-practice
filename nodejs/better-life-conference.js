const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false,
})

function giveDirections(r, c) {
    let direction = c <= 10 ? "Right" : "Left"
    let row = 11 - r
    let chair = Math.min(c, 21 - c)
    return `${direction} ${row} ${chair}`
}
var r = 0;
var c = 0;

rl.on("line", function (line) {
    r_c = line.split(" ").map(element => parseInt(element, 10))
    r = r_c[0]
    c = r_c[1]
    console.log(giveDirections(r, c))
    rl.close()
})