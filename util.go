package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

//改变工作目录到可执行文件所在目录
//否则相对路径读写的文件可能会有问题
func switchPwd() {
	var wd string

	if runtime.GOOS == "linux" {
		if exe, err := os.Readlink("/proc/self/exe"); err != nil {
			log.Println("switchPwd:read exe path err:", err)
			os.Exit(1)
		} else {
			wd = filepath.Dir(exe)

		}
	} else {
		wd = filepath.Dir(os.Args[0])
	}

	if err := os.Chdir(wd); err != nil {
		log.Println("switchPwd:chdir to path:", wd, " err:", err)
		os.Exit(1)
	}
}

func loggerInit() {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("open log file error!", err.Error())
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile)) //同时写到stdout和日志文件
}

func mustInt(src string) int {
	i, _ := strconv.Atoi(src)
	return i
}

//获取ClientIp
func GetClientIp(r *http.Request) string {
	var resultIp string
	realIp := r.Header.Get("X-Real-Ip")
	fwFor := r.Header.Get("X-Forwarded-For")

	if realIp == "" && fwFor == "" { //无代理
		spIdx := strings.LastIndex(r.RemoteAddr, ":") //从尾部找冒号，因为IPv6在地址中部有冒号
		if spIdx == -1 {
			return r.RemoteAddr
		}
		return r.RemoteAddr[:spIdx]
	} else if fwFor != "" { //优先看forward for
		ipLst := strings.Split(fwFor, ",")
		resultIp = strings.TrimSpace(ipLst[0])
	} else {
		resultIp = realIp
	}
	//由于fw头可能是伪造的，这里进行一次校验
	if net.ParseIP(resultIp) != nil {
		return resultIp
	} else {
		return ""
	}
}

//加载指定html文件，允许传数据
func loadHtml(file string, data interface{}) []byte {
	t, err := template.ParseFiles(file)
	if err != nil {
		return []byte("tmpl parse error:" + err.Error())
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)

	if err != nil {
		return []byte("tmpl exec error:" + err.Error())
	}

	return buf.Bytes()
}
