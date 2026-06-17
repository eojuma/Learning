
const myNumbers= ["one","Two","Three","Four","Five"]
// 1. for loop
for (let index = 0; index<myNumbers.length; index++) {
    console.log(myNumbers[index]);
    
}
console.log()

//  2.while loop
let i=0;
while(i<myNumbers.length){
    console.log(myNumbers[i]);
    i++;
}
console.log();
// 3.do while loop
let j=0;
do{
    console.log(myNumbers[j]);
    j++;
}while(j<myNumbers.length);


console.log();
// 4.for each
myNumbers.forEach((s)=>(console.log(s)));


console.log()

// 5.for of (arrays)
for(const n of myNumbers){
console.log(n);
}
const names={
    it:"information technology",
    yt:"youtube",
    JSON:"JavaCript Object Notation"
}
console.log()
// 6.for in  (objects)
for(const a in names){
console.log(a); // prints out the keys for the objects in js
console.log(names[a]); // prints the values

}