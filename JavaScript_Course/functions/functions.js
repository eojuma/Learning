function sayHello(name){

    console.log(`Hello there,${name}`);
}

sayHello("Juma");
greetings=helloInSwahili();
console.log(greetings);

function helloInSwahili(){
    return "Habari gani?"
}


function getUserRole(name,role){
    switch (role) {
        case "admin":
            return `${name} is admin with all access`
            break;
    
            case "subadmin":
            return `${name} is subadmin with access to create and delete courses`
            break;

            case "testpreps":
            return `${name} is testpreps with access to create and delete tests`
            break;

            case "user":
            return `${name} is user with access to consume the content`
            break;
        default:
            return `${name} is a trial user`
            break;
    }
}

console.log(getUserRole("Juma","admin"))

//notes
// function declarations are scanned and made available
// variable declarations are scanned and made undefined

// when you declare a function as a variable then the function must be called after the variable
