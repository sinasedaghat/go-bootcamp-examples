# fmt.Printf()
## Verbs
All of verbs are type-safe in `fmt.Printf()`
- `%v`: verb; print all type of value
```go
fmt.Printf("%v", "string")     // string
fmt.Printf("%v", 32)           // 32
fmt.Printf("%#v", [4]string{}) // [4]string{"", "", "", ""} (show array type and values)
fmt.Printf("%+v", struct{field string}{field: "test"}) // {field:test} (for print structs, it adds field names before the values)
```
- `%q`: quote 
```go
fmt.Printf("%q", "Quote") // "Quote"
```
- `%s`: string
```go
fmt.Printf("%s", "String") // String
fmt.Printf("(%1s)", "S")   // (S)
fmt.Printf("(%2s)", "S")   // ( S)
fmt.Printf("(%3s)", "S")   // (  S)
fmt.Printf("(%4s)", "S")   // (   S)
fmt.Printf("(%5s)", "S")   // (    S) // right aligned (5 character space)
fmt.Printf("(%-5s)", "S")  // (S    ) // left aligned (5 character space)
```
- `%d`: integer number
```go
fmt.Printf("%d", 20) // 20
```
- `%f`: float number
```go
fmt.Printf("%f", 32.604)   // 32.604000
fmt.Printf("%.0f", 32.604) // 33
fmt.Printf("%.1f", 32.604) // 32.6
fmt.Printf("%.2f", 32.604) // 32.60
fmt.Printf("%.3f", 32.604) // 32.604
fmt.Printf("%.4f", 32.604) // 32.6040
fmt.Printf("%.5f", 32.604) // 32.60400
fmt.Printf("%.6f", 32.604) // 32.604000
```
- `%g`: versatile format specifier for floating-point numbers.
```go
fmt.Printf("%g\n", 123.456)       // 123.456
fmt.Printf("%g\n", 0.0000123456)  // 1.23456e-005
fmt.Printf("%g\n", 123456789.123) // 1.23457e+008
fmt.Printf("%g\n", 5.0)           // 5
```
- `%e`: format a floating-point number in scientific notation (also known as E notation)
```go
fmt.Printf("%e\n", 123.456)       // 1.234560e+002
fmt.Printf("%e\n", 0.0000123456)  // 1.234560e-005
fmt.Printf("%e\n", 123456789.123) // 1.234568e+008
fmt.Printf("%e\n", 5.0)           // 5.000000e+000
```
- `%t`: boolean (bool)
```go
fmt.Printf("%t", true) // true
```
- `%T`: type of variable
```go
var speed int
fmt.Printf("%T", speed) // int
```
- `%p`: pointer
```go
x := 10
fmt.Printf("%p", &x) // 0xc000018090
```
- `%c`: character depending on code point (rune)
```go
fmt.Printf("%c", 65)  // A
fmt.Printf("%c", 'A') // A
```
- `%x`:  code point (rune) as  hexadecimal
```go
fmt.Printf("%x", 65)  // 41
fmt.Printf("%x", 'A') // 41
fmt.Printf("%x", "✅") // e29c85
fmt.Printf("% x", "✅") // e2 9c 85
```
- `%[x]anyVerb`: xth value (Argument index)
```go
fmt.Printf(
  "%[1]v is %d away. Think! %[2]d kms! %[1]s OMG!",
  "Venus", 261,
) // Venus is 261 away. Think! 261 kms! Venus OMG!
```
- `%%`: show % sign
```go
fmt.Printf("%v%%", 24) // 24%
```

## Escape Sequences
- `\n`: use for printing new line
```go
fmt.Printf("Hi\nHi") 
/*
  Hi
  Hi
*/
```
- `\t`: use for printing horizontal tab
```go
fmt.Print("Age:\t30") // Age:	  30
```
