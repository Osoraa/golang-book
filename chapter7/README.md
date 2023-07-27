# Introduction to GO by Caleb Doxsey

## Chapter 7: Structs and Interfaces

### Key Takeaways
- The `struct` keyword along with `type` is used to create collections or new types, akin to classes.  
``` Go
type newStructName struct {
    fieldName fieldType
}
```

- An instance of a new type can be created in different ways;  
`var instanceName structName` or `instanceName := new(structName)`.  
Using the `var` keyword creates the instance but with it's fields set to their corresponding zero value.  
Using the `:=` operator creates the instance in memory and returns a pointer to the space in memory.

- We can give the instance fields initial values as the case may be;  
`instance := structName{field: value}` or `instance := structName{value, value, ...}` if the order of the fields are known.  
A pointer to a new instance can be gotten thus: `instance := &structName{value}`

- Struct fields are referenced using the dot `.` operator, i.e. `instance.field`