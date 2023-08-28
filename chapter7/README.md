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
A pointer to a new instance can be gotten thus: `instance := &structName{value}`.

- Struct fields are referenced using the dot `.` operator, i.e. `instance.field`.

- Methods are functions that ony specific types should be able to call. They use receivers after the `func` keyword to specify what type owns them. Go methods are created thus;
```Go
func (rec *Rectangle) area() float64 {
    return rec.length * rc.breadth
}
```

- A type can be embedded in another type so that it's properties are made directly callable on the new type, demonstrated [here](./embedded_types.go).

- **[Interfaces](./interfaces.go)** allows different structs to be referred to all at once as long as they implement the set of methods that the interface describes.  
They can also be used as fields in structs allowing for waay more complex types and robust code.

- Type-assertions allows to cast an interface back to it's type struct.

- Type-switches allows for multiple tyoe assertions.