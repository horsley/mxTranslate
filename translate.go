package main

import (
	"net/http"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/translate/v2"
)

func ApiKey(u string) googleapi.CallOption { return apiKey(u) }

type apiKey string

func (q apiKey) Get() (string, string) { return "key", string(q) }

func doTranslate(q, src, target string) (string, error) {
	svc, err := translate.New(http.DefaultClient)
	if err != nil {
		return "", err
	}

	call := svc.Translations.List([]string{q}, target)

	if src != "" {
		call.Source(src)
	}

	resp, err := call.Do(ApiKey(API_KEY))
	if err != nil {
		return "", err
	}

	if len(resp.Translations) < 1 {
		return "翻译结果为空！", nil
	}

	return resp.Translations[0].TranslatedText, nil
}
