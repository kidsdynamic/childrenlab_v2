package controller

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"log"
)

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	fmt.Printf("\n%x\n", h.Sum(nil))

	return fmt.Sprintf("%x", h.Sum(nil))

}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%x", b))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func checkInsertResult(result sql.Result) bool {
	updated, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return false
	}

	if updated == 0 {
		return false
	}

	return true
}
