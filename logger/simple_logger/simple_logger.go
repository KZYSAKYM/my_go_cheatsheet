package simple_logger

import (
    "log"
    "fmt"
    "bytes"
)

type Logger struct {
	logger *log.Logger
	buf bytes.Buffer
}

var (
	info_logger Logger
	warn_logger Logger
	error_logger Logger
	debug_logger Logger
)

func INFO(msg ...interface{}) {
    if info_logger.logger == nil {
        info_logger.logger = log.New(
		&info_logger.buf, "[INFO]: ", log.Lshortfile)
    }

    __log(&info_logger, msg)
}

func WARN(msg ...interface{}) {
    if warn_logger.logger == nil {
        warn_logger.logger = log.New(
		&warn_logger.buf, "[WARN]: ", log.Lshortfile)
    }

    __log(&warn_logger, msg)
}

func ERROR(msg ...interface{}) {
    if error_logger.logger == nil {
        error_logger.logger = log.New(
		&error_logger.buf, "[ERROR]: ", log.Lshortfile)
    }

    __log(&error_logger, msg)
}

func DEBUG(msg ...interface{}) {
    if debug_logger.logger == nil {
        debug_logger.logger = log.New(
		&debug_logger.buf, "[DEBUG]: ", log.Lshortfile)
    }

    __log(&debug_logger, msg)
}

func __log(logger *Logger, msg ...interface{}) {
    logger.logger.Print(msg)
    fmt.Println(&logger.buf)
}
