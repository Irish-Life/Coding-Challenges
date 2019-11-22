def isNumInString(numStrings :list):
    matches = []
    for entry in numStrings:
        strArr = list(entry)
        for char in strArr:
            if char.isdigit():
                matches.append(entry)
    print(matches)
    return matches
    
  
#numString = ["1a", "a", "2b", "b"] 
numString = ["abc", "abc10"] 

isNumInString(numString)
