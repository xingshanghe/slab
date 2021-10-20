// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  chan_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:35 下午
 */
package slab

// NewChanPool reate a chan based slab allocation memory pool.
// @Description: 创建chan内存池
// @Date: 2021-10-19 16:29:47
// @param minSize 	最小分块大小
// @param maxSize 	最大分块大小
// @param factor 	赠长步长
// @param pageSize 	页数
// @return *ChanPool
//
func NewChanPool(minSize, maxSize, factor, pageSize int) *ChanPool {
	return nil
}

type ChanPool struct {
	pool
}

func (p *ChanPool) Alloc(size int) []byte {
	panic("implement me")
}

func (p *ChanPool) Free(bytes []byte) {
	panic("implement me")
}

var _ Class = (*chanClass)(nil)

type chanClass struct {
	size      int
	page      []byte
	pageBegin uintptr
	pageEnd   uintptr
	chunks    chan []byte
}

func (*chanClass) Push(mem []byte) {
	panic("implement me")
}

func (*chanClass) Pop() []byte {
	panic("implement me")
}
