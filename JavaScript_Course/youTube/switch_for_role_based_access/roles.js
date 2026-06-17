var user ="admin"
switch (user){
    case "admin":
        console.log("Gets full access");
        break;

        case "subadmin":
        console.log("Gets access to create or delete courses");
        break;

        case "testprep":
        console.log("Gets access to create or delete tests");
        break;

        case "user":
        console.log("Gets access to consume the content");
        break;

        default:
            console.log("Trial user")
            break;
}