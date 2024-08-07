function fizzBuzzTest(firstNumber, lastNumber) {
    console.log("FizzBuzz with Js")
    for (let i = firstNumber; i <= lastNumber; i++) {
        if (i % 15 == 0) {
            console.log("FizzBuzz");
        } else if (i % 5 == 0) {
            console.log("Buzz");
        } else if (i % 3 == 0) {
            console.log("Fizz");
        } else {
            console.log(i);
        }
    }
}

fizzBuzzTest(1, 100);