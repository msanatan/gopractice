/**
 * Returns the index of the target in the list, or -1 if it wasn't found
 * @param {Array<number>} items
 * @param {number} target
 */
function linearSearch(items, target) {
  for (let i = 0; i < items.length; i++) {
    if (items[i] === target) {
      return i;
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
let result = linearSearch(numbers, 12);
verify(result);
result = linearSearch(numbers, 6);
verify(result);
