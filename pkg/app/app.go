package app

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}

func MD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func FreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
