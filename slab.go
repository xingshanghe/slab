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

func (p *pool) Alloc(size int) []byte {
	if size <= p.maxSize {
		for i := 0; i < len(p.classes); i++ {
			if p.classes[i].Size() >= size {
				mem := p.classes[i].Pop()
				if mem != nil {
					return mem[:size]
				}
				break
			}
		}
	}
	return make([]byte, size)
}

func (p *pool) Free(mem []byte) {
	size := cap(mem)
	for i := 0; i < len(p.classes); i++ {
		if p.classes[i].Size() == size {
			p.classes[i].Push(mem)
			break
		}
	}
}

type Class interface {
	Push(mem []byte) // 入栈
	Pop() []byte     // 出栈

	Size() int
}

type NoPool struct{}

func (*NoPool) Alloc(size int) []byte {
	return make([]byte, size)
}

func (*NoPool) Free([]byte) {}

var _ Pool = (*NoPool)(nil)
var _ Pool = (*pool)(nil)
