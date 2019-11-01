def first_vowel(txt):
    # list of vowels
    vowels = ["a", "e", "o", "i", "u"]
    # list concantenation that stores string as list of lowercase letters
    txt = [letter.lower() for letter in txt]
    for i, letter in enumerate(txt):
        if letter in vowels:
            print(i, letter)  # comment this entire line if running time test
            break


first_vowel("Why nymphs flyby this")


def first_vowel_recursive(txt, i=0):
    # list of vowels
    vowels = ["a", "e", "o", "i", "u"]
    txt = [letter.lower() for letter in txt]
    letter = txt.pop(0)
    if letter in vowels:
        print(i, letter)  # comment this entire line if running time test
        pass
        
    else:
        i += 1
        first_vowel_recursive(txt, i)


first_vowel_recursive('Why nymphs flyby this')

"""
You can uncomment the time test below to do a speed test on the two fucntions for doing this,
I've run it a couple of times for 100000 iterations and the loop function is on average 7 times faster
"""

# if __name__ == "__main__":
#     import timeit

#     iterations = 100000
#     recursive_time = (
#         timeit.timeit(
#             "first_vowel_recursive('Why nymphs flyby this')",
#             setup="from __main__ import first_vowel_recursive",
#             number=iterations,
#         )
#         / iterations
#     )
#     loop_time = (
#         timeit.timeit(
#             "first_vowel('Why nymphs flyby this')",
#             setup="from __main__ import first_vowel",
#             number=iterations,
#         )
#         / iterations
#     )


#     if loop_time < recursive_time:
#         print(f'loop time is {recursive_time/loop_time} times faster than recursive time ')
#     else:
#         print(f'recursive time is {recursive_time/loop_time} times faster than loop time')

