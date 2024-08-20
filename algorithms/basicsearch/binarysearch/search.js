/**
 * Returns the index of the target in the list, or -1 if it wasn't found
 * @param {Array<number>} items
 * @param {number} target
 * @returns {number}
 */
function binarySearch(items, target) {
  let first = 0;
  let last = items.length;

  while (first <= last) {
    let median = Math.floor((first + last) / 2);
    if (items[median] === target) {
      return median;
    } else if (items[median] < target) {
      first = median + 1;
    } else {
      last = median - 1;
    }
  }

  return -1;
}

/**
 * Prints for human readable debugging
 * @param {number} index
 */
function verify(index) {
  if (index > -1) {
    console.log(`Target was found at position ${index}`);
  } else {
    console.log("Target was not found");
  }
}

const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
let result = binarySearch(numbers, 12);
verify(result);
result = binarySearch(numbers, 6);
verify(result);
