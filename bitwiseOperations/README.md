# Bitwise Operations

& - `sets 1 if both digits are 1` : 0b1100 & 0b1000 = 0b1000

| - `sets 1 if at least one of digits are 1` : 0b1011 | 0b1100 = 0b1111

a^b - `sets 1 if both digits are different` : 0b1100 ^ 0b1010 = 0b0110 

a^ - `sets bit to opposite` : ^0b1001 = 0b0110

&^ - `clears the bits included in the second operand` : 0b1010 &^ 0b1101 = 0b0010

<< n - `left bit shift, multiplication by 2ᴺ(complemented by 0 on the right)`: 1 << 3 = 8
001 << 3 = 1000

>> n  - `right bit shift, division by 2ᴺ (high bits are discarded)` : 16 >> 2 = 4
10000 >> 2 = 100