def linear_search(items: list[int], target: int) -> int | None:
    '''
    Returns the index of the target in the list, or None if it wasn't found
    '''
    for i in range(0, len(items)):
        if items[i] == target:
            return i
    return None


def verify(index: int | None):
    if index is not None:
        print(f'Target found at index {index}')
    else:
        print('Target not found in list')


numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
result = linear_search(numbers, 12)
verify(result)
result = linear_search(numbers, 6)
verify(result)
