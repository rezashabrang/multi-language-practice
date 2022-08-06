const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false,
})

function firstAngry(k) {
    if (k % 2 == 0) {
        return "Bala Barare"
    }
    return "Payin Barare"

}

rl.on("line", function (line) {
    var K = parseInt(line, 10);
    var answer = firstAngry(K);
    console.log(answer);
    rl.close();
})