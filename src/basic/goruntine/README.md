# Channel
  Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
  By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
  - 性質
    - 一種資料通道 (pipe)
	- 一頭write 一頭read
	- FIFO
	- Thread-safe 可同時讀寫
	- 一個通道只能單一 data type
	- 使用完畢要用 close 關閉
  - 使用
    - 建構: make(chan int)
	- 作為 function 參數時，明定 r/w
  - Buffer
    - Buffered (緩衝)
	  - 非阻塞 (unblocking)
	  - 非同步 (async)
	  - Deadlock
	    - over buffer size
	- Unbuffered (無緩衝)
	  - 阻塞 (blocking)
	  - 同步 (sync)
	  - Deadlock
		- 直接 w/r。可透過 goruntine執行或給予 buffered 避免
		- 先r再w。給緩衝也沒用，沒值。