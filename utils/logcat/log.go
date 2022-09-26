package logcat

import (
	"fmt"
	"log"
)

// Ez wrap log

func Print(v ...any) {
	log.Println(v...)
}

func Info(v ...any) {
	log.Printf("\033[1;34m[INFO] %s\033[0m\n", fmt.Sprint(v...))
}

func Good(v ...any) {
	log.Printf("\033[1;32m[GOOD] %s\033[0m\n", fmt.Sprint(v...))
}

func Error(v ...any) {
	log.Printf("\033[1;31m[ERROR] %s\033[0m\n", fmt.Sprint(v...))
}

func ErrorEnd(v ...any) {
	log.Fatalf("\033[1;31m[ERROR] %s\033[0m\n", fmt.Sprint(v...))
}

func Warn(v ...any) {
	log.Printf("\033[1;33m[WARN] %s\033[0m\n", fmt.Sprint(v...))
}

func Debug(v ...any) {
	log.Printf("\033[1;35m[DEBUG] %s\033[0m\n", fmt.Sprint(v...))
}
