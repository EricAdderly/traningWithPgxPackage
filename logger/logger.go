package logger

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// логгер для фу-ий, ошибки для которых написаны вручную (не были пройдены какие-либо валидации)
func ErrorLog(functionName string, err string) {
	time := time.Now()
	log.WithFields(log.Fields{
		"time":  time,
		"func":  functionName,
		"error": err,
	}).Error(err)
}

// Инфологер
func InfoLog(functionName string, info string) {
	time := time.Now()
	log.WithFields(log.Fields{
		"time":  time,
		"func":  functionName,
		"error": info,
	}).Info(info)
}

// логгер для фу-ий, где ошибки были получены в процессе выполнения комплияции
func ErrorLogWithError(functionName string, err error) {
	time := time.Now()
	log.WithFields(log.Fields{
		"time":  time,
		"func":  functionName,
		"error": err,
	}).Error(err)
}
