/*
channel的基本概念:
	1. Go语言的并发模型是CSP,提倡通过通信共享内存而不是通过共享内存实现通信
	2. channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
	3. Go语言中的通道是一种特殊的类型,总是遵循先入先出的规则,保证收发数据的顺序
	4. 每一个通道都是一个具体类型的导管,声明channel的时候需要为其指定元素类型
	5. 通道是引用类型,其零值是nil
	6. 声明的通道需要使用make函数初始化后才能使用
	7. 根据make创建通道时,是否分配缓冲大小分为有缓冲通道和无缓冲通道
channel的基本操作:
	1. 发送: chan <- 10
	2. 接收: val,ok := <-chan //ok为布尔值,当通道关闭后返回false
	3. 关闭: close(chan) // 关闭通道不是必须的
	4. 创建: make(chan, type, cap)
	5. 遍历: for range // 遍历之前需要先关闭通道
关闭后的通道有以下特点:
	1. 对一个关闭的通道再发送值就会导致panic
	2. 对一个关闭的通道进行接收会一直获取值,直到通道为空
	3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值
	4. 关闭一个已经关闭的通道会导致panic
无缓冲的通道:
	1. 无缓冲通道又称为阻塞通道,通过make(chan type)创建
	2. 无缓冲通道只有在有人接收值的时候才能发送值，否则就会阻塞
	3. 无缓冲通道上的发送操作会阻塞,直到另一个goroutine在该通道上执行接收操作,
       如果接收操作先执行,接收方的goroutine会阻塞,直到另一个goroutine在该通道上发送一个值
	4. 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此,无缓冲通道也称为同步通道
有缓冲的通道:
	1. 通过make(chan type,cap)创建有缓冲的通道,要求cap大于0表示通道中能存放元素的数量
	2. 一旦通道中的元素数量超过cap值,通道就会阻塞
	3. 可以通过len()和cap()内建函数获取通道中元素的数量和通道的容量
单向通道:
	1. 单向通道是只能用于发送或者只能用于接收的通道
	2. 只写单向通道: chan<- type,只能往通道中写入type类型的数据,只能执行发送操作
	3. 只读单向通道: <-chan type,只能从通道中读取type类型的数据,只能执行接收操作
	4. 在函数传参和任何赋值操作中可以将双向通道转化为单向通道,但反过来不行
select多路复用:
	1. 用于同时从多个通道接收/发送数据,有一系列的case分支和一个默认的分支,每个case会对应一个通道的通信(接收或发送)过程
	2. select会一直等待,直到某个case的通信操作完成时,就会执行case分支对应的语句
	3. 如果多个case同时满足,select会随机选择一个执行
	4. 对于没有case的select{}会一直等待,可用于阻塞main函数
互斥锁:
	1. Go语言内置sync包中的Mutex实现了互斥锁的功能,保证同一时间有且只有一个goroutine进入临界区
	2. sync.Mutex.Lock(): 加锁
	3. sync.Mutex.Unlock(): 解锁
	4. 多个goroutine同时等待一个锁时,唤醒的策略是随机的
读写互斥锁:
	1. 当我们并发的去读取一个资源但不涉及资源的修改时是没必要加锁的,这种场景下使用读写锁更适合
	2. Go语言中使用sync包中的RWMutex实现读写锁
	3. 读写锁分为读锁和写锁。当一个goroutine获取读锁之后,其他的goroutine如果是获取读锁会继续得锁,
	   如果是获取写锁就会等待;当一个goroutine获取写锁之后,其他的goroutine无论是获取读锁还是写锁都会等待
	4. sync.RWMutex.Lock、sync.RWMutex.Unlock(): 加写锁、解写锁
	5. sync.RWMutex.RLock()、sync.RWMutex.RUnlock(): 加读锁、解读锁
sync.WaitGroup:
	1. Go语言中用于实现并发同步的内置包,是一个结构体,传递的时候要传递指针
	2. sync.WaitGroup内部维护着一个计数器,通过对计数器进行增减来统计当前并发任务个数,当计数器为0时表示并发任务完成
	3. sync.WaitGroup.Add(n int): 计数器+n
	4. sync.WaitGroup.Done(): 计数器-1
	5. sync.WaitGroup.Wait(): 阻塞直到计数器为0
sync.Map:
	1. Go语言内置的map不是并发安全的,所以提供了一个开箱即用并发安全的sync.Map
	2. 开箱即用表示sync.Map不需要make函数初始化就能使用,并且内置了Store、Load、LoadOrStore、Delete、Range等方法
sync/atomic原子操作:
	1. Go内置的原子操作标准库,提供了底层的原子级内存操作
*/
package main
