def binary_search(items: list[int], target: int) -> int | None:
    first = 0
    last = len(items) - 1

    while first <= last:
        midpoint = (first + last) // 2

        if items[midpoint] == target:
            return midpoint
        elif items[midpoint] < target:
            first = midpoint + 1
        else:
            last = midpoint - 1

    return None


def verify(index: int | None):
    if index is not None:
        print(f'Target found at index {index}')
    else:
        print('Target not found in list')


numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
result = binary_search(numbers, 12)
verify(result)
result = binary_search(numbers, 6)
verify(result)
