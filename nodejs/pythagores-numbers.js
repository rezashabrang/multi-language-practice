const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false,
})

function checkRightTriangle(numbers) {
    if (!numbers.every((num) => {return num > 0}))
        return "NO"
    numbers.sort((a,b) => a - b)
    let a = numbers[0], b = numbers[1], c = numbers[2]
    if (Math.pow(a, 2) + Math.pow(b, 2) === Math.pow(c, 2)) {
        return "YES"
    }

    return "NO"
}
var cnt = 0;
var input = [];

rl.on("line", function (line) {
    input.push(parseInt(line, 10));
    cnt++;
    if (cnt === 3) {
        console.log(
            checkRightTriangle(input)
        )
        rl.close()
    }
})

// const main = async () => {
//     let input = []
//     for (let i = 0; i < 3; i++) {
//         inp = await question()
//         input.push(inp)
//     }
//     rl.close()
//     input = input.map((x) => {return parseInt(x, 10)})
//     res = checkRightTriangle(input)
//     console.log(res)
// }

// main()