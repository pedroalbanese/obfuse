package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// ===== flags =====

var (
	str = flag.String("s", "", "String to obfuscate")
	varn = flag.String("v", "str", "Variable name")
	raw  = flag.Bool("r", false, "Print summarized code")
)

// ===== helpers =====

func rotateLeftByte(b byte, n uint8) byte {
	n &= 7
	return (b<<n | b>>(8-n)) & 0xFF
}

func rotateRightByte(b byte, n uint8) byte {
	n &= 7
	return (b>>n | b<<(8-n)) & 0xFF
}

func obfuscateChunkAddRotateXor(chunk []byte, add byte, rot uint8, xor byte) string {
	obf := make([]byte, len(chunk))
	for i, b := range chunk {
		tmp := byte((uint16(b) + uint16(add)) & 0xFF)
		tmp = rotateLeftByte(tmp, rot)
		obf[i] = tmp ^ xor
	}
	return base64.StdEncoding.EncodeToString(obf)
}

func randByteInRange(min, max byte) byte {
	var b [1]byte
	rng := int(max-min) + 1
	for {
		io.ReadFull(rand.Reader, b[:])
		if int(b[0]) < 256-(256%rng) {
			return byte(int(b[0])%rng) + min
		}
	}
}

// ===== generator =====

func generateGoCode(secret, varname string, summarized bool) string {
	secretBytes := []byte(secret)
	chunkSize := 8

	type frag struct {
		encoded string
		add     byte
		rot     uint8
		xor     byte
	}

	var fragments []frag

	for i := 0; i < len(secretBytes); i += chunkSize {
		end := i + chunkSize
		if end > len(secretBytes) {
			end = len(secretBytes)
		}

		add := randByteInRange(1, 254)
		rot := randByteInRange(1, 7)
		xor := randByteInRange(1, 255)

		enc := obfuscateChunkAddRotateXor(secretBytes[i:end], add, uint8(rot), xor)
		fragments = append(fragments, frag{enc, add, uint8(rot), xor})
	}

	var b strings.Builder

	// ===== FULL MODE =====
	if !summarized {
		b.WriteString("package main\n\n")
		b.WriteString("import (\n")
		b.WriteString("\t\"encoding/base64\"\n")
		b.WriteString("\t\"fmt\"\n")
		b.WriteString("\t\"log\"\n")
		b.WriteString(")\n\n")

		b.WriteString("func rotateRightByte(b byte, n uint8) byte {\n")
		b.WriteString("\tn &= 7\n")
		b.WriteString("\treturn (b>>n | b<<(8-n)) & 0xFF\n")
		b.WriteString("}\n\n")

		b.WriteString("func decryptFragment(encoded string, add byte, rot uint8, xor byte) []byte {\n")
		b.WriteString("\tdata, err := base64.StdEncoding.DecodeString(encoded)\n")
		b.WriteString("\tif err != nil { log.Fatal(err) }\n")
		b.WriteString("\tfor i := range data {\n")
		b.WriteString("\t\ttmp := data[i] ^ xor\n")
		b.WriteString("\t\ttmp = rotateRightByte(tmp, rot)\n")
		b.WriteString("\t\tdata[i] = byte((uint16(tmp) + 256 - uint16(add)) & 0xFF)\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn data\n")
		b.WriteString("}\n\n")

		b.WriteString("func main() {\n")
	}

	// ===== COMMON OUTPUT =====

	fmt.Fprintf(&b, "\tvar %s []byte\n\n", varname)

	for _, f := range fragments {
		fmt.Fprintf(
			&b,
			"\t%s = append(%s, decryptFragment(%q, 0x%X, %d, 0x%X)...)\n",
			varname, varname,
			f.encoded, f.add, f.rot, f.xor,
		)
	}

	fmt.Fprintf(&b, "\n\tfmt.Println(string(%s))\n", varname)
		
	if !summarized {
		b.WriteString("}\n")
	}

	return b.String()
}

// ===== main =====

func main() {
	flag.Parse()

	if *str == "" {
		fmt.Printf("Usage: %s -s <string> [-v name] [-r]\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	fmt.Println(generateGoCode(*str, *varn, *raw))
}
