package godb

import "bytes"

const (
	maxKey = 24
	maxVal = page - maxKey // 4072
)

// database record interface
type dbRecord interface {
	key() []byte // return fixed length key (  24 bytes) from record data
	val() []byte // return fixed length val (4072 bytes) from the record data, excluding nil byte remainders
}

// data record
type record struct {
	data []byte // contains a fixed length 24 byte key, leaving 4072 bytes for the value
}

// create a pointer to a new record
func newRecord(key, val []byte) *record {
	return &record{append(key, val...)}
}

// return key from data record
func (r *record) key() []byte {
	return r.data[:maxKey]
}

// return val from data record
func (r *record) val() []byte {
	if n := bytes.IndexByte(r.data[maxKey:], 0x00); n > -1 {
		return r.data[maxKey:n]
	}
	return r.data[maxKey:]
}
