const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});
readline.question("",n => {
    console.log(`${factorial(n)}`);
    readline.close();
});

function factorial(n) {
    if (n == 1) return n

    return n * factorial(n - 1);
}