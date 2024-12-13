package filter

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type BloomFilter struct {
	bitArray  []byte       // 位数组
	size      int          // 位数组大小
	hashCount int          // 哈希函数数量
	mu        sync.RWMutex // 保护并发访问
}

// NewDefaultBloomFilter  创建一个默认大小的布隆过滤器
func NewDefaultBloomFilter() *BloomFilter {
	return NewBloomFilterBySizeAndHashCount(2^24, 3)
}

func NewBloomFilterBySizeAndHashCount(size, hashCount int) *BloomFilter {
	return &BloomFilter{
		bitArray:  make([]byte, size),
		size:      size,
		hashCount: hashCount,
	}
}

// hash 函数：使用 FNV-1a 哈希算法
func (bf *BloomFilter) hash(data string, i int) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(fmt.Sprintf("%d-%s", i, data)))
	if err != nil {
		return 0
	}
	return int(h.Sum32()) % bf.size
}

// Add 向布隆过滤器添加元素
func (bf *BloomFilter) Add(data string) {
	bf.mu.Lock()
	defer bf.mu.Unlock()
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(data, i)
		bf.bitArray[index] = 1
	}
}

// Exists 判断元素是否存在于布隆过滤器中
func (bf *BloomFilter) Exists(data string) bool {
	bf.mu.RLock() // 读锁，允许多个读取
	defer bf.mu.RUnlock()
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(data, i)
		if bf.bitArray[index] == 0 {
			return false
		}
	}
	return true
}

// Load 导出，用于存储第三方
func (bf *BloomFilter) Load() []byte {
	return bf.bitArray
}

// Store 用于第三方导入
func (bf *BloomFilter) Store(src []byte) {
	bf.size = max(bf.size, len(src))
	bf.bitArray = make([]byte, bf.size)
	copy(bf.bitArray, src)
}
