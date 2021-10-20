// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  atom_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:29 下午
 */
package slab

// NewAtomPool create a atom slab allocation memory pool.
// @Description: 创建Atom内存池
// @Date: 2021-10-19 16:29:47
// @param minSize 	最小分块大小
// @param maxSize 	最大分块大小
// @param factor 	赠长步长
// @param pageSize 	页数
// @return *AtomPool
//
func NewAtomPool(minSize, maxSize, factor, pageSize int) *AtomPool {
	return nil
}

type AtomPool struct {
	pool
}

func (p *AtomPool) Alloc(size int) []byte {
	panic("implement me")
}

func (p *AtomPool) Free(bytes []byte) {
	panic("implement me")
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
	next uint32
}

func (*atomClass) Push(mem []byte) {
	panic("implement me")
}

func (*atomClass) Pop() []byte {
	panic("implement me")
}
