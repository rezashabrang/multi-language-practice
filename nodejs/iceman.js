const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});
readline.question("",n => {
    iceManAnswer(parseInt(n))
    readline.close();
});

function iceManAnswer(n) {
    if (n > 100)
        console.log("Steam")
    else if (n < 0)
        console.log("Ice")
    else
        console.log("Water")
    return
}