package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func getHandler() http.Handler {
	fs := http.FileServer(http.Dir("static"))

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			rw.Write(loadHtml("static/index.html", map[string]interface{}{
				"Lang": template.JS(rawLangList),
			}))
		} else {
			fs.ServeHTTP(rw, r)
		}
	})
}

func translateHandler(rw http.ResponseWriter, r *http.Request) {
	if r.ContentLength < 3 {
		http.Error(rw, "forbidden", http.StatusForbidden)
		return
	}

	var reqParam map[string]string
	j := json.NewDecoder(r.Body)
	err := j.Decode(&reqParam)
	if err != nil {
		log.Println("translateHandler decode err:", err)
		http.Error(rw, "decode err", http.StatusForbidden)
		return
	}

	resp, err := doTranslate(reqParam["q"], reqParam["src"], reqParam["target"])
	if err != nil {
		resp = "翻译出错：" + err.Error()
		log.Println("translateHandler doTranslate err:", err)
	}

	rw.Write([]byte(resp))
}
