package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

var (
	// MaxSizeForChan 通道大小
	MaxSizeForChan = 50000
)

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSizeForChan),
	}
	err = fl.initFile() // 按照文件路径和文件名打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启后台goroutine去写日志
	go f.writeLogBackground()
	return nil
}

// 判断是否需要记录
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// 根据文件大小，判断文件是否需要切割
func (f FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("fet file info failed, err:%v\n", err)
		return false
	}
	// 如果当前文件大小 >= 日志文件设置的最大值，就返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f FileLogger) spliteFile(file *os.File) (*os.File, error) {
	// 需要切割
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return file, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", f.fileName, nowStr)
	// 1.关闭当前日志文件
	file.Close()
	// 2.备份一下rename
	os.Rename(logName, newLogName)
	// 3.打开新的文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	// 4.将打开的新文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.spliteFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.spliteFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR级别，还再在err日志文件中在记录一边
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}

	}
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		// 先把日志发送到通道
		logTmp := &logMsg{ // 造logMsg对象
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			// 把日志丢掉，保证不出现阻塞
		}

	}
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, format, a...)
	}
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARING, format, a...)
}

// Error ...
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close ...
func (f FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
