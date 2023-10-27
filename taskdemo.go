import (
	"fmt"
	"sync"  
)
func handleTask() {
  demoList := []string{"123", "456", "abc", "ddd"}
  var wg = sync.WaitGroup{}
  
  for idx, item := range demoList {
    // 每个元素创建一个新协程去处理
    wg.Add(1) // 
    go func(idx int){
      defer wg.Done() // 协程退出前将 wg 计数器减一，否则最后计数器无法减为0，会一直卡在 wg.wait() 那一行
      
      // 内部还想新建子协程去做不同的事，减少执行时间
      var innerWg = sync.WaitGroup{} // 
      
      innerWg.Add(1)
      go func(){
        defer innerWg.Done() // 子协程退出前将 innerWg 的计数器减一
        fmt.Printf("子协程一正在执行")
      }
      
      innerWg.Add(1)
      go func(){
        defer innerWg.Done() // 子协程退出前将 innerWg 的计数器减一
        fmt.Printf("子协程二正在执行")
      }
      innerWg.Wait() // 等所有所有子协程执行完才继续往下执行
      fmt.Printf("外层协程%d即将执行完毕", idx)
    }(idx)
  }
  
  wg.Wait() // 等待 wg 的计数器减为0后才继续往下执行，等待所有的协程处理完毕
  
  fmt.Println("Execution completed. will exit")
}
func main() {
  handleTask()
}
