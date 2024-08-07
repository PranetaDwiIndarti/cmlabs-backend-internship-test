<?php
function fizzBuzzTest($firstNumber, $lastNumber) {
    echo "FizzBuzz with PHP";
    for ($i = $firstNumber; $i <= $lastNumber; $i++) {
        if ($i % 15 == 0) {
            echo "FizzBuzz\n";
        } elseif ($i % 5 == 0) {
            echo "Buzz\n";
        } elseif ($i % 3 == 0) {
            echo "Fizz\n";
        } else {
            echo $i . "\n";
        }
    }
}

fizzBuzzTest(1, 100);
?>