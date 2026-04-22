var testArray = [2, 3, 4, 5, 6, 7, 8, 9];

//fill function

// Just as in other languages js respects index starting from 0
//the start index is inclusive and the end index is exclusive

// In the example below the testArray elements from index 2
// all the way to the 5 index will be filled with h instead
// of the actual values in these indexees
console.log(testArray.fill("h", 2, 6));

//filter function

//This function creates a new array excluding the specified elements

// In the example below the number 60 is not part of the res variable
const numbers = [60, 12, 34, 5, 6, 70, 90, 60, 60];

const res = numbers.filter((num) => num != 60);

console.log(res);
