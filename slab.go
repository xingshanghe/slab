// Package slab
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  slab
 * @Version: 0.1.0
 * @Date: 2021/10/18 3:21 下午
 */
package slab

type Pool interface {
	Alloc(size int) []byte
	Free([]byte)
}

type pool struct {
	classes []Class
	minSize int
	maxSize int
}

type Class interface {
	Push(mem []byte)
	Pop() []byte
}

var _ Class = (*atomClass)(nil)

type NoPool struct{}

func (*NoPool) Alloc(size int) []byte {
	return make([]byte, size)
}

func (*NoPool) Free([]byte) {}

var _ Pool = (*NoPool)(nil)
var _ Pool = (*AtomPool)(nil)
