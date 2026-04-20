var temperature;
temperature = 210;

if (temperature<20){
    console.log("It is cold outside");
}else if (temperature>20){
    console.log("Hot temperature")
}else{
    console.log("Normal")
}

// ternary operator


var authenticated = true;
var count=false;

authenticated && count ? console.log("SignOut"):console.log("Login"); 
// the first  part shows the true condition while the second part is the false condition
// the second part is the basically the fallback in case anything fails