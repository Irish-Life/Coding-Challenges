//8535902775

var isbn = 8535902775 +'';
//var isbn = 1843369283 +'';

var arr1 = isbn.split('');

//var arr1 = [8,5,3,5,9,0,2,7,7,5];
var arr2 = [10,9,8,7,6,5,4,3,2,1];

var result1 = 0;
var result2 = 0;
var result3 = 0;
var result4 = 0;
var result5 = 0;
var result6 = 0;
var result7 = 0;
var result8 = 0;
var result9 = 0;
var result10 = 0;

for (var i=0; i < 1; i++) {
  result1 += arr1[0] * arr2[0];
  result2 += arr1[1] * arr2[1];
  result3 += arr1[2] * arr2[2];
  result4 += arr1[3] * arr2[3];
  result5 += arr1[4] * arr2[4];
  result6 += arr1[5] * arr2[5];
  result7 += arr1[6] * arr2[6];
  result8 += arr1[7] * arr2[7];
  result9 += arr1[8] * arr2[8];
  result10 += arr1[9] * arr2[9];


}
let finalSum = (result1 + result2 + result3
  + result4
  + result5
  + result6
  + result7
  + result8
  + result9
  + result10)

if(finalSum % 11 === 0) {
  console.log("This is a valid ISBN")
}
else {
  console.log("This is invalid")
}
/*console.log(finalSum);
console.log(result1 , result2 , result3
  , result4
  , result5
  , result6
  , result7
  , result8
  , result9
  , result10);

console.log(result1 + result2 + result3
  + result4
  + result5
  + result6
  + result7
  + result8
  + result9
  + result10
  );*/