"""
An ISBN number is legal if it consists of 10 digits and
d_1 + 2*d_2 + 3*d_3 + ... + 10*d_10 is a multiple of 11.
The ISBN number 0-201-31452-5 is valid.
1*5 + 2*2 + 3*5 + 4*4 + 5*1 + 6*3 + 7*1 + 8*0 + 9*2 + 10*0 = 88
and 88 is a multiple of 11.
Write a program which validates whether an ISBN number.
For Ex:
1) if user inputs 8535902775 then it should print "Valid"
2) if user inputs 1843369283 then it should print "Not Valid"
"""
# the nicely laid out one
def is_valid(isbn):

    if len(isbn) != 10:
        return False

    isbn = reversed(isbn)
    coefficients = [(ord(c) - 48) * (i+1) for i, c in enumerate(isbn)]
    return sum(coefficients) % 11 == 0

print(is_valid("0201314525"))
print(is_valid("1843369283"))

# the one liner
def one_line(isbn):
    return sum([(ord(c) - 48) * (i+1) for i, c in enumerate(reversed(isbn))]) % 11 == 0 and len(isbn) == 10

print(one_line("0201314525"))
print(one_line("1843369283"))
