var counted = document.getElementById("counter");
var follower = document.getElementById("followers");

let count = 1;
let followers = "Followers on Instagram"

setInterval(()=>{
  if (count<1000){
    count++;
    counted.innerText = count;
  }
},5)

setTimeout(()=>{
  follower.innerText = followers;
},5000)