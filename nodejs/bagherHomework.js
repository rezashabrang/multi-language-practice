const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});
readline.question("",n => {
    canBeTriangle(n);
    readline.close();
});
/**
 * 
 * @param {string} n 
 */
function canBeTriangle(n) {
    let angles = n.split(" ");
    let sumOfAngles = angles.reduce((result, number) => {
        return parseInt(result) + parseInt(number)
    });
    if (sumOfAngles === 180 && angles.every(i => i > 0))
        console.log("Yes")
    else
        console.log("No")
}