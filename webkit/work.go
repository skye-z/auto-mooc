package webkit

import (
	_ "embed"
	"log"
	"sync"
	"time"
)

//go:embed script.js
var script []byte

// 声明全局变量
var (
	// 任务会话
	workSession *Session
	// 任务锁
	workLock sync.Mutex
	// 任务状态
	workStatus = "stopped"
	// 退出信号
	workQuit chan bool
	// 任务错误次数
	workErrorNumber int
)

// 创建任务
func CreateWork(session *Session, status *WebKit) bool {
	// 加锁防并发
	workLock.Lock()
	defer workLock.Unlock()
	// 检查任务是否已启动
	if workStatus == "started" {
		return false
	}
	status.Running = true
	status.Error = ""
	// 存储任务会话
	workSession = session
	// 创建退出信号
	workQuit = make(chan bool)

	_, err := session.Page.Evaluate(string(script))
	if err != nil {
		log.Printf("[Work] Init script error: %v", err)
	}
	_, err = session.Page.Evaluate(string("selectClass()"))
	if err != nil {
		log.Printf("[Work] Run script error: %v", err)
	}

	// 启动协程
	go func() {
		for {
			select {
			case <-workQuit:
				log.Println("[Work] Discontinued")
				return
			default:
				// 每3秒检测1次
				time.Sleep(10 * time.Second)
				WorkContent(session, status)
			}
		}
	}()
	// 更新协程的状态
	workStatus = "started"
	return true
}

// 关闭任务
func CloseWork(status *WebKit) bool {
	// 加锁防并发
	workLock.Lock()
	defer workLock.Unlock()
	// 检查任务是否已停止
	if workStatus == "stopped" {
		return false
	}
	log.Println("[Work] Discontinuing...")
	// 发出关闭信号
	workQuit <- true
	// 更新状态
	workStatus = "stopped"
	status.Running = false
	// 等待1秒,随后关闭浏览器
	time.Sleep(1 * time.Second)
	workSession.Page.Close()
	workSession.Context.Close()
	workSession.Browser.Close()
	return true
}

// 任务内容
func WorkContent(session *Session, status *WebKit) {
	page := session.Page
	// 查找是否存在视频播放器
	visible, _ := page.Locator("#player").IsVisible()
	if !visible {
		// 错误超过3次, 结束任务
		if workErrorNumber == 3 {
			status.Error = "未检测到播放器,请检查是否需要答题"
			CloseWork(status)
			return
		}
		log.Println("[Work] Found anomaly")
		workErrorNumber++

		// 尝试选择课程
		_, err := session.Page.Evaluate(string("selectClass()"))
		if err != nil {
			log.Printf("error: %v", err)
		}
	} else {
		// 恢复播放, 清空错误计数
		workErrorNumber = 0
	}
}
