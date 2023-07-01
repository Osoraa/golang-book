# Introduction to GO by Caleb Doxsey

## Chapter 5: Arrays, Slices, and Maps

### Key Takeaways
- Arrays, Slices and Maps are built-in types in Go.

- **Arrays** have to be of the same type.

- The `len()` function returns an int type. Type conversion uses the type like a function. E.g. `float64(len(num))`.

- Use and underscore in for variables you won't use. Like for unused iterators in for loops. Awesome!!!

- **Slices** are types built on top of arrays in Go and are often seen used instead of arrays.

- Unlike arrays, the length of a slice is allowed to change.

- Slices are always associated with an array, and can never be longer than than the associated array.

- Use the make function to create slices in Go. First parameter specifies type, second specifies length of the slice, and the third is the capacity of the array the slice is associated with.