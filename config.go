package godb

/*
 *	global (constants and variables)
 */

/*
 *	mmap / engine (constants and variables)
 */
const (
	slab = 1 << 26 // 64MB
	page = 1 << 12 //  4KB
)

var temp = make([]byte, page, page)

/*
 *	tree / index (constants and variables)
 */
