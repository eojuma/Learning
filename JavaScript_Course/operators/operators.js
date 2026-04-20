var num1 = 7;
var num2 = 6;

console.log(num1+num2);
console.log(num1*num2);
console.log(num1%num2);
console.log(num1-num2);
console.log(num1>num2);
console.log(num1<num2);

var listPrice = 799;
var sellingPrice = 199;
percentageDiscount=((listPrice-sellingPrice)/listPrice)*100;
percentagedis=Math.round(percentageDiscount);
console.log(percentagedis+"% off");