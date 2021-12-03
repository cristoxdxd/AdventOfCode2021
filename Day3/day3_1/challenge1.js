const fs = require('fs');
const data = fs.readFileSync('./input.txt', 'utf-8',);
const input = data.split("\n");

function mapNumbers(data){
    let ones = 0;
    let zeros = 0;
    let map = {};
    for (let i = 0; i < data.length; i++){
        let line = data[i];
        for (let j = 0; j < line.length-1; j++){
            let number = line[j];
            if (map[j]){
                map[j].push(number);
            }else{
                map[j] = [
                    number
                ];
            }
        }
    }
    let Gamma
    let Epsilon
    for (let i; i < 12; i++){
        for (let j; j < map[i].length; j++){
            console.log(map[j])
            if (map[i][j] == 1){
                ones++;
            }else{
                zeros++;
            }
        }
        if (ones > zeros){
            Gamma[i] = [0];
            Epsilon[i] = [0];
        }else{
            Gamma[i] =[0];
            Epsilon[i] = [1];
        }
        let decimalGamma = parseInt(Gamma, 2);        
        let decimalEpsilon = parseInt(Epsilon, 2);
    return decimalGamma*decimalEpsilon       
}

console.log(mapNumbers(input));
