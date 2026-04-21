function isEven(number) {
  if (number % 2 === 0) {
    return true;
  }
  return false;
}

console.log(isEven(54));

// the arrow function

var addition = (num1,num2) => {
  return num1+num2;
};

console.log(addition(3,8));


//callback  function

var res=[2,4,6,8].every(isEven);
console.log(res);