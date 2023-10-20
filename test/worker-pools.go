package main

import (
	"fmt"
	"time"
)

/**
 * <-chan 数据输出的管道
 * chan<- 数据放入的管道
 */
func worker(id int, jobs <-chan int, results chan<- int) {
	for v := range jobs {
		fmt.Println("worker id:", id, ",started job:", v)
		time.Sleep(time.Second)
		fmt.Println("worker id:", id, ",finished job:", v)
		results <- v * 2
	}
	close(results)
}

// 这是 worker 程序，我们会并发的运行多个 worker。
// worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。
// 每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。
func main() {
	const num = 5

	results := make(chan int, num)
	jobs := make(chan int, num)

	//for i := 0; i < 3; i++ {
	//go协程启动了work函数，jobs内没有数据，暂时阻塞了
	go worker(1, jobs, results)
	//}

	for i := 0; i < num; i++ {
		//放入数据进入jobs
		jobs <- i
	}
	//关闭了jobs管道
	close(jobs)

	for i := 0; i < num; i++ {
		//results 将结果输出出来
		//<-results
	}

}
