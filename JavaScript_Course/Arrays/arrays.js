var names = ["Silas","Maxwell"]

console.log(names)

// adding elements  to an array using spread operator
newNames =["Joel","Caleb"]
allNames=[...names,...newNames]
console.log(allNames)

// adding elements to an array using the shift function

allNames.unshift("James") //add an element in the first position[0]
console.log(allNames)



//removing an element in an existing array

//using pop function

allNames.pop(); //removes the last element in the array
console.log(allNames)

//using the shift function

allNames.shift();  //equally removes the last name in the array
console.log(allNames);
