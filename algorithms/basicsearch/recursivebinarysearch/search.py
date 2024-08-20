def recursive_binary_search(items: list[int], target: int) -> bool:
    # Handle the base case where the list is empty, there can be nothing to find
    if len(items) == 0:
        return False

    midpoint = len(items)//2

    if items[midpoint] == target:
        return True

    # Otherwise, make the problem set smaller
    if items[midpoint] < target:
        return recursive_binary_search(items[midpoint+1:], target)
    else:
        return recursive_binary_search(items[:midpoint], target)


def verify(result: bool):
    print(f'Target found: {result}')


numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
result = recursive_binary_search(numbers, 12)
verify(result)
result = recursive_binary_search(numbers, 6)
verify(result)
