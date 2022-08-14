const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false,
})

function firstPowerTwo(number) {
    for (let i = 0; i < 1000; i++) {
        let powerTwo = Math.pow(2, i)
        if (powerTwo >= number){
            return powerTwo
        }
    }
}

rl.on("line", function (line) {
    res = firstPowerTwo(parseInt(line, 10));
    console.log(res);
        rl.close()
    }
)