package main

import (
	"fmt"
	"math"
	"unicode/utf16"
)

// Function to display how an ASCII character is encoded and decoded, including memory representation.
func encodeDecodeASCII(char rune) {
	// ASCII Encoding (1 byte)
	asciiByte := byte(char)
	fmt.Printf("ASCII Encoding: '%c' -> %v \n", char, asciiByte)

	//Print the byte in different formats
	fmt.Printf("Memory Representation of '%c':\n", char)
	fmt.Printf(" Binary: %07b\n", asciiByte)      // Binary format 7 bits
	fmt.Printf(" Hexadecimal: %02x\n", asciiByte) // Hexadecimal format
	fmt.Printf(" Decimal: %d\n", asciiByte)       // Decimal fomrat

	// ASCII Decoding (1 byte)
	decodeASCII := rune(asciiByte)
	fmt.Printf("ASCII Decoding: %v -> '%c' \n", asciiByte, decodeASCII)
	fmt.Printf("Number of unique characters: %d\n", int(math.Pow(2, 7)))
}

// Other encode/decoding character unicode (UTF-8, UTF-16, UTF-23)

// Encode and decode a UTF-8 character, including memory representation.
func encodeDecodeUTF8(char rune) {

	// UTF-8 Encoding (variable bytes)
	utf8Bytes := []byte(string(char))
	fmt.Printf("UTF-8 Encoding: '%c' -> %v\n", char, utf8Bytes)

	//Print the byte in different formats
	fmt.Printf("Memory Representation of '%c':\n", char)
	fmt.Printf(" Binary: %07b\n", utf8Bytes)      // Binary format 1 - up to 4 byte
	fmt.Printf(" Hexadecimal: %02x\n", utf8Bytes) // Hexadecimal format
	fmt.Printf(" Decimal: %d\n", utf8Bytes)       // Decimal fomrat

	// UTF-8 Decoding (variable bytes)
	decodingUTF8 := string(utf8Bytes)
	fmt.Printf("UTF-8 Decoding: %v -> '%s'\n", utf8Bytes, decodingUTF8)

	fmt.Printf("Number of unique characters: %d\n", int(math.Pow(2, 4*8)))
}

// Encode and decode a UTF-16 character, including memory representation.
func encodeDecodeUTF16(char rune) {

	// UTF-16 Encoding (variable bytes)
	utf16Bytes := utf16.Encode([]rune{char})
	fmt.Printf("UTF-16 Encoding: '%c' -> %v\n", char, utf16Bytes)

	//Print the byte in different formats
	fmt.Printf("Memory Representation of '%c':\n", char)
	fmt.Printf(" Binary: %016b\n", utf16Bytes)     // Binary format  2 - up to 4 byte
	fmt.Printf(" Hexadecimal: %02x\n", utf16Bytes) // Hexadecimal format
	fmt.Printf(" Decimal: %d\n", utf16Bytes)       // Decimal fomrat

	// UTF-16 Decoding (variable bytes)
	decodingUTF16 := string(utf16.Decode(utf16Bytes))
	fmt.Printf("UTF-8 Decoding: %v -> '%s'\n", utf16Bytes, decodingUTF16)
	fmt.Printf("Number of unique characters: %d\n", int(math.Pow(2, 4*8)))

}

// Encode and decode a UTF-32 character, including memory representation.
func encodeDecodeUTF32(char rune) {

	// UTF-32 Encoding (variable bytes)
	utf32Bytes := int32(char)
	fmt.Printf("UTF-32 Encoding: '%c' -> %v\n", char, utf32Bytes)

	//Print the byte in different formats
	fmt.Printf("Memory Representation of '%c':\n", char)
	fmt.Printf(" Binary: %032b\n", utf32Bytes)     // Binary format 4 byte for all characters
	fmt.Printf(" Hexadecimal: %02x\n", utf32Bytes) // Hexadecimal format
	fmt.Printf(" Decimal: %d\n", utf32Bytes)       // Decimal fomrat

	// UTF-32 Decoding (variable bytes)
	decodingUTF32 := rune(utf32Bytes)
	fmt.Printf("UTF-8 Decoding: %v -> '%c'\n", utf32Bytes, decodingUTF32)
	fmt.Printf("Number of unique characters: %d\n", int(math.Pow(2, 4*8)))

}

func main() {

	char := 'A' // 65
	encodeDecodeASCII(char)
	fmt.Println("------------------------------------")
	encodeDecodeUTF8(char)
	fmt.Println("------------------------------------")
	encodeDecodeUTF16(char)
	fmt.Println("------------------------------------")
	encodeDecodeUTF32(char)

	cheinsChar := '世' // 65
	encodeDecodeASCII(cheinsChar)
	fmt.Println("------------------------------------")
	encodeDecodeUTF8(cheinsChar)
	fmt.Println("------------------------------------")
	encodeDecodeUTF16(cheinsChar)
	fmt.Println("------------------------------------")
	encodeDecodeUTF32(cheinsChar)

	emoji := '🙂' // 65
	encodeDecodeASCII(emoji)
	fmt.Println("------------------------------------")
	encodeDecodeUTF8(emoji)
	fmt.Println("------------------------------------")
	encodeDecodeUTF16(emoji)
	fmt.Println("------------------------------------")
	encodeDecodeUTF32(emoji)
}
