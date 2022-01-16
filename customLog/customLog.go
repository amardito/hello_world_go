package customLog

import (
	"fmt"
	"time"
	"log"
)

func Error(l *log.Logger, forModule string, msg error) {
	moduleName := fmt.Sprintf(" [%v] ", forModule)
	//the l params bring exported types of Logger on log.New, for more Informations https://pkg.go.dev/log#Logger
	l.SetPrefix(" \n[Error] " + time.Now().Format("2006-01-02 15:04:05") + moduleName)
	l.Print(msg)
}

func Success(l *log.Logger, forModule string, msg interface{}) {
	moduleName := fmt.Sprintf(" [%v] ", forModule)
	l.SetPrefix(" \n[Success] " + time.Now().Format("2006-01-02 15:04:05") + moduleName)
	l.Print(msg)
}