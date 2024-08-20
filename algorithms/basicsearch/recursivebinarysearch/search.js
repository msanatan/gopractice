/**
 * Returns the index of the target in the list, or -1 if it wasn't found
 * @param {Array<number>} items
 * @param {number} target
 * @returns {boolean}
 */
function recursiveBinarySearch(items, target) {
  if (items.length === 0) {
    return false;
  }

  const midpoint = items.length / 2;
  if (items[midpoint] === target) {
    return true;
  }

  if (items[midpoint] < target) {
    return recursiveBinarySearch(items.slice(midpoint + 1), target);
  } else {
    return recursiveBinarySearch(items.slice(0, midpoint), target);
  }
}

/**
 * Prints for human readable debugging
 * @param {boolean} index
 */
function verify(result) {
  console.log(`Target found: ${result}`);
}

const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
let result = recursiveBinarySearch(numbers, 12);
verify(result);
result = recursiveBinarySearch(numbers, 6);
verify(result);
