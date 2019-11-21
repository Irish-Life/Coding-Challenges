// numInStr(["1a", "a", "2b", "b"]) ➞ ["1a", "2b"]
// numInStr(["abc", "abc10"]) ➞ ["abc10"]
// numInStr(["abc", "ab10c", "a10bc", "bcd"]) ➞ ["ab10c", "a10bc"]
// numInStr(["this is a test", "test1"]) ➞ ["test1"]

const strings = ["abc", "ab10c", "a10bc", "bcd"];

const result = strings.filter(strings => strings.match(/\d/));
//filter the strings that match the regex \d\
// \d matches any single digit in most regex grammar styles,

console.log(result);
