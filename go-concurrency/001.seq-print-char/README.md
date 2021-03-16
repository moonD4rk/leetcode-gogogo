# 多线程顺序输出字符

1. 开启三个线程顺序输出 `ABC` 三个字符各 10 次
2. 开启 n 个线程顺序输出 1-n 个数字

## FanIn-Fanout

```go

// generate 生成 chan
func generate(msg string, times int) chan message {
	m := make(chan message)
	go func() {
		defer close(m)
		for i := 0; i < times; i++ {
			m <- message{
				str: fmt.Sprintf("%s %d", msg, i),
			}
		}
	}()
	return m
}

// fanIN 从多个 chan 中读取数据，并将起发送给结果 chan
func fanIn(times int, inputs ...chan message) chan message {
	result := make(chan message)
	go func() {
		defer close(result)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				result <- <-input
			}
		}
	}()
	return result
}

```