package goid

import (
	"crypto/rand"
	"fmt"

	base58 "github.com/jbenet/go-base58"
	uuid "github.com/satori/go.uuid"
)

const base58Alphabet = base58.FlickrAlphabet

type UUID struct {
	value [16]byte
	err   error
}

func (u *UUID) String() (string, error) {
	str := fmt.Sprintf("%x-%x-%x-%x-%x", u.value[0:4], u.value[4:6], u.value[6:8], u.value[8:10], u.value[10:])
	return str, u.err
}

func (u *UUID) Base58Encode() (string, error) {
	return base58.EncodeAlphabet(u.value[:], base58Alphabet), u.err
}

func EncodeBase58(v4 string) string {
	id, _ := uuid.FromString(v4)
	return base58.EncodeAlphabet(id.Bytes(), base58Alphabet)
}

func NewUUID() *UUID {
	uuid := &UUID{}
	_, uuid.err = rand.Read(uuid.value[:])
	uuid.setVarint()
	uuid.setVersion()
	return uuid
}

func (u *UUID) setVarint() {
	const mask = 0x80
	u.value[8] = (u.value[8] & 0x3f) | mask //nolint:gomnd
}

func (u *UUID) setVersion() {
	const mask = 0x40
	u.value[6] = (u.value[6] & 0x0f) | mask //nolint:gomnd
}
