var users = ["Juma","Silas","John","Maxwell"]


// slice function

//The slice function in js trims the specified elements excluding the last index
// Rhe second log will print from index 2 to the end
console.log(users.slice(1,4));
console.log(users.slice(2));


//splice function

// It accepts three values the start index,the count to the right
//and the value(s)(could be as many as you like) to replace the specified elements
var users1 = ["Bonface","Catherine","Dorine","Caren"]
updated=users1.splice(0,2,"Olama");
console.log(users1);