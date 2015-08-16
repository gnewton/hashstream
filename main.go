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
	if _, err := io.Copy(writer, reader) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(md5Sum)
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
