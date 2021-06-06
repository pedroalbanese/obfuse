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
        "fmt"
        "unsafe"
)

const (
        EAX = uint8(unsafe.Sizeof(true))
)

func main() {
var uid []byte

uid = append(uid, ((EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX<<EAX^EAX))
uid = append(uid, ((((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX<<EAX^EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)
uid = append(uid, (((EAX<<EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX)<<EAX|EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX)
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX)<<EAX|EAX))
uid = append(uid, (((EAX<<EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX)<<EAX)
uid = append(uid, (((EAX<<EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX<<EAX)
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX<<EAX)
uid = append(uid, (((EAX<<EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX<<EAX)
uid = append(uid, ((((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)
uid = append(uid, (((EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX<<EAX|EAX))
uid = append(uid, ((((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX|EAX))
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)
uid = append(uid, (((EAX<<EAX^EAX)<<EAX^EAX)<<EAX<<EAX<<EAX|EAX))
uid = append(uid, ((((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX|EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX^EAX)<<EAX|EAX))
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX)
uid = append(uid, (((EAX<<EAX^EAX)<<EAX<<EAX^EAX)<<EAX|EAX)<<EAX)
uid = append(uid, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX<<EAX^EAX))
fmt.Println(string(uid))
}
```
## License

This project is licensed under the ISC License.

#### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Lab.
