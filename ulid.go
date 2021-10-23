package goid

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/oklog/ulid"
)

type ULID struct {
	id  ulid.ULID
	err error
}

func (u *ULID) String() (string, error) {
	if u.err != nil {
		return "", u.err
	}
	return u.id.String(), nil
}

func NewULID() *ULID {
	return NewFromTime(time.Now().UTC())
}

func NewFromTime(t time.Time) *ULID {
	id, err := ulid.New(ulid.Timestamp(t), rand.Reader)
	ret := &ULID{
		id:  id,
		err: err,
	}
	return ret
}

func NewFromTimeAndReader(t time.Time, r io.Reader) *ULID {
	id, err := ulid.New(ulid.Timestamp(t), r)
	ret := &ULID{
		id:  id,
		err: err,
	}
	return ret
}

func NewFromTimeOnly(t time.Time) *ULID {
	id := &ulid.ULID{}
	err := id.SetTime(ulid.Timestamp(t))
	return &ULID{id: *id, err: err}
}
