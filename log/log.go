package log

import (
	"log"
	"os"
)

func init() {
	// 获取日志文件句柄
	// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`./go-method.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}
func Print(s string) {
	log.Println(s)
}
