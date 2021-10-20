Slab
============
**the project is forked from [funny/slab](https://github.com/funny/slab)** 

Introduction
============
Slab allocation memory pools for Go.

Usage
=====

Use lock-free memory pool:

```go
pool := slab.NewAtomPool(
	64,          // The smallest chunk size is 64B.
	64 * 1024,   // The largest chunk size is 64KB.
	2,           // Power of 2 growth in chunk size.
	1024 * 1024, // Each slab will be 1MB in size.
)

buf1 := pool.Alloc(64)

    ... use the buf ...
	
pool.Free(buf)
```

Use `chan` based memory pool:

```go
pool := slab.NewChanPool(
	64,          // The smallest chunk size is 64B.
	64 * 1024,   // The largest chunk size is 64KB.
	2,           // Power of 2 growth in chunk size.
	1024 * 1024, // Each slab will be 1MB in size.
)

buf1 := pool.Alloc(64)

    ... use the buf ...
	
pool.Free(buf)
```

Use `sync.Pool` based memory pool:

```go
pool := slab.NewSyncPool(
	64,          // The smallest chunk size is 64B.
	64 * 1024,   // The largest chunk size is 64KB.
	2,           // Power of 2 growth in chunk size.
)

buf := pool.Alloc(64)

    ... use the buf ...
	
pool.Free(buf)
```