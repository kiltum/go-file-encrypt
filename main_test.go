package go_file_encrypt

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func testEq(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestEncryptFile(t *testing.T) {
	d1 := []byte("1234567890")
	file := "/tmp/test_encrypt_file"
	file_pass := "123456"
	err := ioutil.WriteFile(file, d1, 0644)
	if err != nil {
		t.Errorf("%s \n", err)
	}

	d2, err := ioutil.ReadFile(file)

	if err != nil {
		t.Errorf("%s \n", err)
	}

	if testEq(d1, d2) != true {
		t.Errorf("Write: %s  Read: %s\n", d1, d2)
	}

	err = EncryptFile(file, file_pass)

	if err != nil {
		t.Errorf("%s \n", err)
	}

	err = DecryptFile(file+".encrypt", file_pass)

	if err != nil {
		t.Errorf("%s \n", err)
	}

	d2, err = ioutil.ReadFile(file)

	if err != nil {
		t.Errorf("%s \n", err)
	}

	if testEq(d1, d2) != true {
		t.Errorf("Encrypted failed Write: %s  Read: %s\n", d1, d2)
	}

}

func TestDeriveKey(t *testing.T) {
	salt := []byte("salt")
	key, _ := deriveKey("1234567890", salt)
	want := []byte("3f8b7ee77ca57d362e54efbac78f02d421e198c21fd1ae664e196c959b88a5f9")
	dst := make([]byte, hex.DecodedLen(len(want)))
	hex.Decode(dst, want)
	//fmt.Printf("%x %s", key, salt)
	if testEq(key, dst) != true {
		t.Errorf("keys not equal. want %x got %x", dst, key)
	}
}
