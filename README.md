# obfuse
## Minimalist golang string obfuscator 

### Description:
When you want to store specific strings in your go code they are automatically
searchable in the go binary.

For example take the following code (`x.go`):

```go
    package main

    import (
        "fmt"
    )

    const (
        C = "hello"
    )

    func main() {
        fmt.Println(C)
    }
```

Now run `go build ./x.go` to build it and the run `strings ./x | grep hello`.

Surprise ! `hello` is visible to anyone who has access the go binary.

[obfuse](https://github.com/pedroalbanese/obfuse) tries to mitigate this
problem by obfuscating the string so it is no longer searchable with
[strings](https://linux.die.net/man/1/strings).

### Usage:
<pre> -r    Print summarized code.
 -s string
       String to obfuscate.
 -v string
       Variable name. (default "str")</pre>

### Example:
<pre> obfuse -s "35495c82-c65d-11eb-a48d-f30d71e4e8ad" -v uid</pre>
#### Will return:
```go
package main

import (
        "encoding/base64"
        "fmt"
        "log"
)

func rotateRightByte(b byte, n uint8) byte {
        n &= 7
        return (b>>n | b<<(8-n)) & 0xFF
}

func decryptFragment(encoded string, add byte, rot uint8, xor byte) []byte {
        data, err := base64.StdEncoding.DecodeString(encoded)
        if err != nil { log.Fatal(err) }
        for i := range data {
                tmp := data[i] ^ xor
                tmp = rotateRightByte(tmp, rot)
                data[i] = byte((uint16(tmp) + 256 - uint16(add)) & 0xFF)
        }
        return data
}

func main() {
        var uid []byte

        uid = append(uid, decryptFragment("L09fj08onz8=", 0xF0, 4, 0x1D)...)
        uid = append(uid, decryptFragment("jtdoibeOCQk=", 0xB2, 5, 0x75)...)
        uid = append(uid, decryptFragment("Tkj/Ru2VTP8=", 0x2A, 1, 0x51)...)
        uid = append(uid, decryptFragment("3pGl1oGp2pU=", 0xC5, 2, 0x72)...)
        uid = append(uid, decryptFragment("C9lLGw==", 0xCA, 4, 0xF9)...)

        fmt.Println(string(uid))
}
```
## License

This project is licensed under the ISC License.

#### Copyright (c) 2020-2026 Pedro F. Albanese - ALBANESE Research Lab.
