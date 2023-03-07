package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recover() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.next()
	}
}

func trace(msg string) string {
	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	var bulider strings.Builder
	bulider.WriteString(msg + "\n trace stack")
	for _, pc := range pcs[:n] {
		fun := runtime.FuncForPC(pc)
		f, l := fun.FileLine(pc)
		bulider.WriteString(fmt.Sprintf("\n\t%s:%d", f, l))
	}
	return bulider.String()
}
