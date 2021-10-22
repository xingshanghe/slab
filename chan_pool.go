// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  chan_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:35 下午
 */
package slab

import "unsafe"

// NewChanPool reate a chan based slab allocation memory pool.
// @Description: 创建chan内存池
// @Date: 2021-10-19 16:29:47
// @param minSize 	最小分块大小
// @param maxSize 	最大分块大小
// @param factor 	赠长步长
// @param pageSize 	页数
// @return *ChanPool
//
func NewChanPool(minSize, maxSize, factor, pageSize int) Pool {
	cp := &pool{
		classes: make([]Class, 0, 10),
		minSize: minSize,
		maxSize: maxSize,
	}
	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		cc := chanClass{
			size:   chunkSize,
			page:   make([]byte, pageSize),
			chunks: make(chan []byte, pageSize/chunkSize),
		}
		cc.pageBegin = uintptr(unsafe.Pointer(&cc.page[0]))
		for i := 0; i < pageSize/chunkSize; i++ {
			// lock down the capacity to protect append operation
			mem := cc.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			cc.chunks <- mem
			if i == len(cc.chunks)-1 {
				cc.pageEnd = uintptr(unsafe.Pointer(&mem[0]))
			}
		}
		cp.classes = append(cp.classes, cc)
	}
	return cp
}

var _ Class = (*chanClass)(nil)

type chanClass struct {
	size      int
	page      []byte
	pageBegin uintptr
	pageEnd   uintptr
	chunks    chan []byte
}

func (c chanClass) Push(mem []byte) {
	select {
	case c.chunks <- mem:
	default:
		mem = nil
	}
}

func (c chanClass) Pop() []byte {
	select {
	case mem := <-c.chunks:
		return mem
	default:
		return nil
	}
}

func (c chanClass) Size() int {
	return c.size
}
