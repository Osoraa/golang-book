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

- Variadic Functions
