# hashstream
Originally I wrote this to providce crypto hash Writer/Reader support for files too large to fit in memory or streaming. Inspired by Java's [DigestInputStream](https://docs.oracle.com/javase/7/docs/api/java/security/DigestInputStream.html) and [DigestInputStream](https://docs.oracle.com/javase/7/docs/api/java/security/DigestInputStream.html)

However, it was [pointed out](https://groups.google.com/d/msg/golang-nuts/NOtenKfslLg/6GcXf6nx-TUJ) to me that (io.MultWriter)[http://golang.org/pkg/io/#MultiWriter] would be a better way to do this.

Yes, it was!  :-)

So I have replaced my original code with an example of doing the same this with io.MultiWriter

# Example

Can be found in main.go:


```
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

const text0 = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit  
amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante 
hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet 
vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur. Donec ut 
libero sed arcu vehicula ultricies a non tortor. Lorem ipsum dolor sit amet, 
consectetur adipiscing elit. Aenean ut gravida lorem. Ut turpis felis, pulvinar a 
semper sed, adipiscing id dolor. Pellentesque auctor nisi id magna consequat sagittis. 
Curabitur dapibus enim sit amet elit pharetra tincidunt feugiat nisl imperdiet. Ut 
convallis libero in urna ultrices accumsan. Donec sed odio eros. Donec viverra mi quis 
quam pulvinar at malesuada arcu rhoncus. Cum sociis natoque penatibus et magnis dis 
parturient montes, nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae 
nisi at sem facilisis semper ac in est.
`

// From Linux command: md5sum (GNU coreutils) 8.22  on Fedora release 21 (Twenty One)
const md5Sum = "0ed8481c815f1977d7df85cf73e1c078"

func main() {
	reader := strings.NewReader(text0)

	fo, err := os.Create("out.bytes")
	if err != nil {
		panic(err)
	}

	hash := md5.New()
	writer := io.MultiWriter(fo, hash)

	buf := make([]byte, 32)
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := writer.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	fmt.Println(md5Sum)
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}

```

Output:
```
0ed8481c815f1977d7df85cf73e1c078
0ed8481c815f1977d7df85cf73e1c078
```

#Variations
Note that multiple hashes (should you need this) could be calculated at once, with something like:
```
	md5 := md5.New()
	sha256 := sha256.New()
	writer := io.MultiWriter(fo, md5, sha256)
	.
	.
	.
	fmt.Println(hex.EncodeToString(md5.Sum(nil)))
	fmt.Println(hex.EncodeToString(sha256.Sum(nil)))	

```

#License
MIT license. See LICENSE file

Copyright 2015 Glen Newton



