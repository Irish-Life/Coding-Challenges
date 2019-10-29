num_list = [1, 6, 5, 3, 5, 6, 10, 11, 3, 106, 8]
k = 16


def list_check_sum(n: list, k: int) -> bool:
    i = 0
    compare_num = n.pop(i)  # pops first element from list and assigns to compare_num
    while i < len(n):
        if compare_num + n[i] == k:  # if we find a match return True
            print(f"{compare_num} + {n[i]} = {k}")
            return True
        i += 1  # otherwise try next index

    if not n:  # if list is empty
        return False
    else:  # recursively returns the function called on itself
        return list_check_sum(n, k)

print(list_check_sum(num_list, k))

