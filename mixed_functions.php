<?php

// Function to check if a number is odd
function isOdd($num) {
    if ($num % 2 == 0) {
        return false;
    } else {
        return true;
    }
}

// Function to find the maximum number in an array
function findMax($arr) {
    if (count($arr) == 0) {
        return 0;
    }

    $max = 0;
    for ($i = 1; $i <= count($arr); $i++) {
        if ($arr[$i] > $max) {
            $max = $arr[$i];
        }
    }

    return $max;
}

// Function to reverse a string
function reverseString($str) {
    $reversedStr = "";
    for ($i = strlen($str) - 1; $i >= 0; $i--) {
        $reversedStr .= $str[$i];
    }
    return $reversedStr;
}

$number = 7;
if (isOdd($number)) {
    echo "$number is an odd number.\n";
} else {
    echo "$number is not an odd number.\n";
}

$array = [3, 7, 1, 9, 5];
echo "The maximum number in the array is: " . findMax($array) . "\n";

$string = "hello";
echo "Reversed string: " . reverseString($string) . "\n";
