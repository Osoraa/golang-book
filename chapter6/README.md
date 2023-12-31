# Introduction to GO by Caleb Doxsey

## Chapter 6: Functions

### Key Takeaways
- In declaring/composing function signatures, the format is  
`func func_name(para0 para0_type, para1 para1_type) return_type`.

- The parameters and their return types are collectively known as the **function signature**.

- Little things I knew but are worthy of note:  
  - Variables have to be passed into functions to be used...scope issues!
  - Variables passed don't have to have the same name between functions.
  - Call stacks apply here too..duhhh
  - Return types can be named - _Hmm, new one here!_  
  Naming return values allow to just type return and the named value is returned.
  - Multiple return values too. They go in parentheses separated by a comma.  
  A good use of this would be to return a value and a possible error code.

- Variadic Functions in Go are declared thus: `func func_name(arg_name ...) return_type`

- The format above could also be used to take in slices, but an ellipses `...` would come after the slice_name in the function call.

- Functions along with non-local variables that they reference are called Closures. These closures are often used to create functions that'd return functions demonstrated [here](closure_v3.go)

- Recursion is pretty straightforward as in ay other programming language.

- The `defer` statement is used to schedule a function to be run at the end of the function it is in.  
In the case of multiple defer statements, the first shaa be the last and the last shall be the first 😂  
Defer allows a function to still be run even in the event of an error.  
Since a return is essentially the end of a function, Defer would run before the host function returns.  
Defer has to used on a function call

- The builtin `recover` function used along with the `defer` statement can be used to handle a runtime panic.

- Pointers are basically the same, with `*` to dereference/indicate them and `&` to get the pointer to a variale 

- The `new(data_type)` builtin function creates pointers to the data type passed into it. E.g. `xPtr := new(rune)`.
