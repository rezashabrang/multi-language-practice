const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});
readline.question("",n => {
    for (let i = 0; i < parseInt(n) ; i++)
        print()
    readline.close();
});

function print() {
    console.log("man khoshghlab hastam");
}