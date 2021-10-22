// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  atom_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:29 下午
 */
package slab

import (
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

// NewAtomPool create a atom slab allocation memory pool.
// @Description: 创建Atom内存池
// @Date: 2021-10-19 16:29:47
// @param minSize 	最小分块大小
// @param maxSize 	最大分块大小
// @param factor 	赠长步长
// @param pageSize 	页数
// @return *AtomPool
//
func NewAtomPool(minSize, maxSize, factor, pageSize int) Pool {
	ap := &pool{
		classes: make([]Class, 0, 10),
		minSize: minSize,
		maxSize: maxSize,
	}

	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		ac := atomClass{
			size:      chunkSize,
			page:      make([]byte, pageSize),
			pageBegin: 0,
			pageEnd:   0,
			chunks:    make([]chunk, pageSize/chunkSize),
			head:      (1 << 32),
		}

		for i := 0; i < len(ac.chunks); i++ {
			chk := &ac.chunks[i]
			// lock down the capacity to protect append operation
			chk.mem = ac.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			if i < len(ac.chunks)-1 {
				chk.next = uint64(i+1+1 /* index start from 1 */) << 32
			} else {
				ac.pageBegin = uintptr(unsafe.Pointer(&ac.page[0]))
				ac.pageEnd = uintptr(unsafe.Pointer(&chk.mem[0]))
			}
		}
		ap.classes = append(ap.classes, ac)
	}

	return ap
}

var _ Class = (*atomClass)(nil)

type atomClass struct {
	size      int
	page      []byte
	pageBegin uintptr
	pageEnd   uintptr
	chunks    []chunk
	head      uint64
}

type chunk struct {
	mem  []byte
	aba  uint32 // 解决aba问题
	next uint64
}

func (c atomClass) Push(mem []byte) {
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&mem)).Data
	if c.pageBegin <= ptr && ptr <= c.pageEnd {
		i := (ptr - c.pageBegin) / uintptr(c.size)
		chk := &c.chunks[i]
		if chk.next != 0 {
			panic("slab.AtomPool: Double Free")
		}
		chk.aba++
		new := uint64(i+1)<<32 + uint64(chk.aba)
		for {
			old := atomic.LoadUint64(&c.head)
			atomic.StoreUint64(&chk.next, old)
			if atomic.CompareAndSwapUint64(&c.head, old, new) {
				break
			}
			runtime.Gosched()
		}
	}
}

func (c atomClass) Pop() []byte {
	for {
		old := atomic.LoadUint64(&c.head)
		if old == 0 {
			return nil
		}
		chk := &c.chunks[old>>32-1]
		nxt := atomic.LoadUint64(&chk.next)
		if atomic.CompareAndSwapUint64(&c.head, old, nxt) {
			atomic.StoreUint64(&chk.next, 0)
			return chk.mem
		}
		runtime.Gosched()
	}
}

func (c atomClass) Size() int {
	return c.size
}
