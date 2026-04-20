var names  = "Juma";
var courseName = "JavaScript";
var isLoggedIn = true;
var loggedIn = 34;
var paymentMode; 
paymentMode="Credit Card"
console.log(names,courseName,isLoggedIn,loggedIn,paymentMode);



// various ways to add a set of aray to another array

var users =['john','james']
var newUser =['felix','joel']

//spread operator
var allUsers=[...users,...newUser]
console.log(allUsers)
// push
var allUsers=users.push(...newUser)
console.log(users)


//pop removes an item from an array
var popped = users.pop()
console.log(popped)