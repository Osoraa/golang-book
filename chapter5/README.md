# Introduction to GO by Caleb Doxsey

## Chapter 5: Arrays, Slices, and Maps

### Key Takeaways
- Arrays, Slices and Maps are built-in types in Go.

- **Arrays** have to be of the same type.

- The `len()` function returns an int type. Type conversion uses the type like a function. E.g. `float64(len(num))`.

- Number arrays get filled with 0's when elements are not assigned.

- Use and underscore in for variables you won't use. Like for unused iterators in for loops. Awesome!!!  
  > The length of an array is part of it's type name.

- **Slices** are types built on top of arrays in Go and are often seen used instead of arrays.

- Unlike arrays, the length of a slice is allowed to change.

- Slices are always associated with an array, and can never be longer than than the associated array.

- Use the `make()` function to create slices in Go. First parameter specifies type, second specifies length of the slice, and the third is the capacity of the array the slice is associated with.

- Append to a slice using the `append()` function. If the underlying array isnt large enough, then a new array is created the former contents copied over, new contents added and then the new slice is returned.  
  > It's worth noting that the previous array and it's slice would still exist.

- Copy one array into another using the `copy(dst, src)` function. The destination slice comes before the source slice.  
If the destination slice is smaller than the src slice, then the items that can't fit in are discarded.

- Maps aka associative arrays, hash tables or dictionaries are Key-Value pairs for holding data.

- The length of a map is the number of it's key-value pairs.  
Use the builtin `delete(map, key)` function to delete map items.

- An unsuccessful lookup in a map returns athe zero value of the value type.

- [This](mapping.go) introduces mapping in Go. [This](mapping_v2.go) details a more advanced use of maps in Go.

