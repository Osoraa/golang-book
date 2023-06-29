# Introduction to GO by Caleb Doxsey

## Chapter 2: Variables

### Key Takeaways
- The **var** keyword is used to create variables. Variables can be created or assigned upon creation i.e. `var name string` or `var name string = "Osoraa"`

- Th Go compiler can infer types so `var name = "Bookie"` is also fine, so long as the variable is instantiated upon creation.

- Idiomatic Go favours the walrus operator over the `var` statement for variable instantiation. i.e. `name := "Esobe"`.