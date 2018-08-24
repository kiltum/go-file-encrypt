# go-file-encrypt

Here is pretty simple go functions for encrypt/decrypt files. Suitable for very lange files, because not reading them to the memory

Using AES-256 as crypt algorithms.

## Usage

```go
package main

import (
	gfe "github.com/kiltum/go-file-encrypt"
	"fmt"
	"io/ioutil"
	"github.com/prometheus/common/log"
)


func main () {
	
	d1 := []byte("1234567890")
	file := "/tmp/test_encrypt_file"
	file_pass := "123456"
	err := ioutil.WriteFile(file, d1, 0644)
	if err != nil {
		log.Fatal("%s \n", err)
	}

	d2, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("%s \n", err)
	}

	fmt.Printf("Write %s\n", d1)


	err = gfe.EncryptFile(file, file_pass)

	if err != nil {
		log.Fatal("%s \n", err)
	}

	err = gfe.DecryptFile(file+".encrypt", file_pass)

	if err != nil {
		log.Fatal("%s \n", err)
	}

	d2, err = ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("%s \n", err)
	}
	fmt.Printf("Read: %s\n", d2)

	err = gfe.DecryptFile(file+".encrypt", file_pass+"Wrong")

	if err != nil {
		log.Fatal("%s \n", err)
	}

	d2, err = ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("%s \n", err)
	}
	fmt.Printf("Read with wrong pass: %s\n", d2)
}

```

You should get

```
Write 1234567890
Read: 1234567890
Read with wrong pass: [��5�Q�h
```
