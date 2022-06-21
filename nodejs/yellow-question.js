const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});
readline.question("",name => {
    console.log(`W${"o".repeat(parseInt(name))}w!`);
    readline.close();
});