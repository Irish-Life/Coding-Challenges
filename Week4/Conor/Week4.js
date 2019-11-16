//var isbn = 8535902775 + "";

var isbn = 1843369283 + "";

var arr1 = isbn.split("");
var arr2 = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1];
var results = [];

for (var i = 0; i < 10; i++) {
  results.push(arr1[i] * arr2[i]);
}

let finalSum = results.reduce((a, b) => a + b, 0);

if (finalSum % 11 === 0) {
  console.log("This is a valid ISBN");
} else {
  console.log("This is an invalid ISBN");
}
