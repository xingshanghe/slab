// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  sync_pool
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:36 下午
 */
package slab

import "sync"

func NewSyncPool(minSize, maxSize, factor, pageSize int) *SyncPool {
	return nil
}

type SyncPool struct {
	pool
}

func (p *SyncPool) Alloc(size int) []byte {
	panic("implement me")
}

func (p *SyncPool) Free(bytes []byte) {
	panic("implement me")
}

var _ Class = (*syncClass)(nil)

type syncClass struct {
	sync.Pool
}

func (syncClass) Push(mem []byte) {
	panic("implement me")
}

func (syncClass) Pop() []byte {
	panic("implement me")
}

func (syncClass) Size() int {
	panic("implement me")
}
