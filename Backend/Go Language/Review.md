# Variables
 - declarations, walrus operator

# Comments
 - single line, multi line

# Compilation Process
 - code -> compiler -> exe
 - Errors - compilation errors, runtime errors

# Type Sizes
 - uint uint8 uint16 uint32 uint64 uintptr

 - float32 float64
 
 - complex64 complex128 

 - When should I use a more specific type? - When super concerned about performance and memory usage.

# Type Inference
 - auto typing based on value

# Same Line Directions
 - age, name  := 10, "Smith"

# Go runtime
 - purpose is to clean up unused memory at runtime
 - includes a garbage collector that automatically frees up memory that's no longer in use.

# Constants
 - const pi = 3.14

# String Formattings
 - fmt.Printf -> prints
 - fmt.Sprintf -> returns


# Runes and String Encoding
 - a "character" is a single byte (8 bits) 
 - using ASCII encoding we can represent 128 characters with 7bits
 - Go strings are sequences of bytes, but Go has special type 'rune' which is an alias for int32. Means that a rune is a 32-bit integer (4 bytes), which is a large enought o hold any Unicode point.
 - when working with strings you need to be aware of the encoding (bytes -> representation). Go uses UTF-8 encoding which is a variable-length encoding.

# What does this mean?
 - There are 2 main reason:
 - 1. When you need to work with every characters in a string, you should use a "rune" type. It breaks strings up into their individual characters, which can be more than one byte long.
 - 2. We can include a wide variety of Unicode characters in our strings like emojis and Chinese characters and it will be fine.


<br>
<hr>

# Conditionals
 
<br>
<hr>

# Functions

<br>
<hr>

# Structs
 - represents structured data


 ```go
   type car struct {
      brand string
      color string  

   } 

 ```


<br>
<hr>
# Interfaces

<br>
<hr>
# Errors


<br>
<hr>

# Loops

<br>
<hr>
# Slices

- Arrays are:
- fixed-size groups of variables of the same type
```
 nums := [3]int{1,2,3}	
```

<br>
<hr>

# Maps
# Pointers
# Packages and Modules
# Channels
# Mutexes
# Generics 
# Enums


