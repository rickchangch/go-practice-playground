package contextDemo

/*
 Context、WaitGroup、Select Case 皆是處理多個 goruntine 的方法
 而使用情境也有所不同
*/

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type contextDemo struct{}

var Handler = &contextDemo{}

func (c *contextDemo) Run() {

	waitGroupDemo()

	selectCaseDemo()

	ContextTry1()

	ContextTry2()
}

/*
 當您需要將同一件事情拆成不同的 Job 下去執行，
 最後需要等到全部的 Job 都執行完畢才繼續執行主程式，這時候就需要用到 WaitGroup。
 但這種方式卻無法終止執行中的 Job，因此可以使用 Select-case
*/
func waitGroupDemo() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 1 done.")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("job 2 done.")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All Done.")
}

/*
 只要在任何地方，將 true 丟進 stop channel 就可以終止 Job。
 然而，當有複數個 goruntine 在執行，又想控制 goruntine 中的 goruntine 呢？
*/
func selectCaseDemo() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	stop <- true
	time.Sleep(5 * time.Second)
}

/*
 透過 context.Background() 去註冊 Parent context，且每個 Context 都會附帶一個 Cancel Event，
 當 goruntine 透過 cancel 停止後， <-ctx.Done 就取得到值 藉此控制所有 ctx 下的 goruntine。
*/
func ContextTry1() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

/*
 測試 Timeout, Deadline
*/
type contextKey string

func contextWork(ctx context.Context) {
	dealine, ok := ctx.Deadline()
	name := ctx.Value(contextKey("name"))

	if ok {
		log.Println(name, "has deadline:", dealine.Format("2006-01-02 15:04:05"))
	} else {
		log.Println(name, "does not have deadline")
	}
}

func ContextTry2() {

	timeout := 3 * time.Second
	deadline := time.Now().Add(10 * time.Second)

	timeoutContext, timeoutCancelFunc := context.WithTimeout(context.Background(), timeout)
	defer timeoutCancelFunc()

	cancelContext, cancelFunc := context.WithCancel(context.Background())

	deadlineContext, deadlineCancelFunc := context.WithDeadline(context.Background(), deadline)
	defer deadlineCancelFunc()

	contextWork(context.WithValue(timeoutContext, contextKey("name"), "[Timeout Context]"))
	contextWork(context.WithValue(cancelContext, contextKey("name"), "[Canncel Context]"))
	contextWork(context.WithValue(deadlineContext, contextKey("name"), "[Deadline Context]"))

	<-timeoutContext.Done()
	log.Println("timeout ...")

	log.Println("cancel error:", cancelContext.Err())
	log.Println("canncel...")
	cancelFunc()
	log.Println("cancel error:", cancelContext.Err())

	<-cancelContext.Done()
	log.Println("The cancel context has been cancelled...")

	<-deadlineContext.Done()
	log.Println("The deadline context has been cancelled...")
}
