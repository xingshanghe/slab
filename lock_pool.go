// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  lock_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:36 下午
 */
package slab

// NewLockPool lock-free slab allocation memory pool.
// @Description: 创建chan内存池
// @Date: 2021-10-19 16:29:47
// @param minSize 	最小分块大小
// @param maxSize 	最大分块大小
// @param factor 	赠长步长
// @param pageSize 	页数
// @return *ChanPool
//
func NewLockPool(minSize, maxSize, factor, pageSize int) *LockPool {
	return nil
}

type LockPool struct {
	pool
}

func (p *LockPool) Alloc(size int) []byte {
	panic("implement me")
}

func (p *LockPool) Free(bytes []byte) {
	panic("implement me")
}

var _ Class = (*lockClass)(nil)

type lockClass struct {
}

func (*lockClass) Push(mem []byte) {
	panic("implement me")
}

func (*lockClass) Pop() []byte {
	panic("implement me")
}
