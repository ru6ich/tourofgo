# Go basic types:

bool: `1 byte`

string: `n byte`

int: `machine word size - 4 byte on 32-bit systems, 8 byte on 64-bit systems`
int8: `1 byte | from -128 to 127`
int16: `2 byte | from -32768 to 32767` 
int32: `4 byte | from -2147483648 to 2147483647` 
int64: `8 byte |from –9 223 372 036 854 775 808 to 9 223 372 036 854 775 807`

uint: `machine word size - 4 byte on 32-bit systems, 8 byte on 64-bit systems`
uint8: `1 byte | from 0 to 255`
uint16: `2 byte | from 0 to 65535`
uint32: `4 byte | from 0 to 4294967295` 
uint64: `8 byte | from 0 to 18 446 744 073 709 551 615` 
uintptr: `4 byte on 32-bit systems, 8 byte on 64-bit systems. Designed to perform math operations on memory adresses and should not be used to store random integers ` 

byte - `alias for uint8`

rune - `alias for int32, represents a unicode code point`

float32: `4 byte | 6 decimal digit accuracy`
float64: `8 byte | 15 decimal digit accuracy`
complex64: `8 byte` 
complex128: `16 byte`

int, uint, uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit sistems. When you need an integer value you should use int unless you have specific reason to use a sized or unsigned  integer type. 