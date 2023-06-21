# Map

golang的map底层是一个hash表，通过键值对进行映射

1. 数据结构

2. 扩容机制

3. 查找&插入

4. hash冲突

数据结构

```
type hmap struct {
    count int // map大小，len函数取得这个值
    flags uint8 // map状态标识
    B uint8 // 最多容纳 6.5 * 2^8 个元素
    noverflow uint16 // 溢出桶buckets数量
    hash0 uint32 //hash种子
    
    buckets unsafe.Pointer //指向最大2^B 个Buckets组的指针
    oldbuckets unsafe.Pointer // 指向扩容之前的Buckets数组，不增长九尾nil
    nevacuate uintptr // 搬迁精度
    extra *mapextra //额外信息
}

typ mapextra struct {
    overflow *[]*bmap
    oldoverflow *[]*bmap
    nextOverflow *bmap
}


 //在编译期间会产生新的结构体，bucket
 type bmap struct {
     tophash [8]uint8 //存储哈希值的高8位
     data    byte[1]  //key value数据:key/key/key/.../value/value/value...
     overflow *bmap   //溢出bucket的地址
 }
```

![](https://store.babyzoo.club/study/hmap.png)
