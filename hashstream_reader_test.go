package hashstream

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"strings"
	"testing"
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
const md5_0 = "0ed8481c815f1977d7df85cf73e1c078"

// From Linux command: sha1 (GNU coreutils) 8.22  on Fedora release 21 (Twenty One)
const sha1_0 = "063c25f97d32ed98936607d790413eb44971666f"

// From Linux command: sha256 (GNU coreutils) 8.22  on Fedora release 21 (Twenty One)
const sha256_0 = "b73e7fb51bc5b0f52aab3e61ed01bedaa3178be415e7c44398e6972ad42c7cb6"

// From Linux command: sha512 (GNU coreutils) 8.22  on Fedora release 21 (Twenty One)
const sha512_0 = "3a5c261acc31514d590e2557979ac2fc437553a9b170c1d26cdfa93ca9a6cc9d287ec76c66012142d592a81246df9304d88d4328c3973d14ed5c5993321a3a93"

var textCryptoResultMap map[string]map[hash.Hash]string

func resetCrypto() {
	cryptoResultMap := map[hash.Hash]string{
		md5.New():    md5_0,
		sha1.New():   sha1_0,
		sha256.New(): sha256_0,
		sha512.New(): sha512_0,
	}

	textCryptoResultMap = map[string]map[hash.Hash]string{
		text0: cryptoResultMap,
	}

}

func testReader(t *testing.T) {
	resetCrypto()
	for text, tcrm := range textCryptoResultMap {
		for hsh, sum := range tcrm {
			hexhash, err := applyHash(text, hsh, 32)
			if err != nil || hexhash != sum {
				t.FailNow()
			}
		}
	}

}

func Test_Reader_ShouldFailWithAlteredText(t *testing.T) {
	resetCrypto()
	for text, tcrm := range textCryptoResultMap {
		for hsh, sum := range tcrm {
			hexhash, err := applyHash(text+"f", hsh, 32)
			if hexhash == sum || err != nil {
				t.FailNow()
			}
		}
	}

}

func applyHash(text string, hs hash.Hash, bufSize int) (string, error) {
	reader := strings.NewReader(text)
	hr := NewHashReader(reader, hs)

	for {
		buf := make([]byte, bufSize)
		_, err := hr.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return hex.EncodeToString(hr.Sum()), nil

}
