package main

import (
	"encoding/json"
	"log"
)

type LangItem struct {
	Name string
	Val  string
}

var rawLangList = `
[
{"Name": "Chinese Simplified", "Val": "zh-CN"},
{"Name": "Afrikaans", "Val": "af"},
{"Name": "Albanian", "Val": "sq"},
{"Name": "Arabic", "Val": "ar"},
{"Name": "Armenian", "Val": "hy"},
{"Name": "Azerbaijani", "Val": "az"},
{"Name": "Basque", "Val": "eu"},
{"Name": "Belarusian", "Val": "be"},
{"Name": "Bengali", "Val": "bn"},
{"Name": "Bosnian", "Val": "bs"},
{"Name": "Bulgarian", "Val": "bg"},
{"Name": "Catalan", "Val": "ca"},
{"Name": "Cebuano", "Val": "ceb"},
{"Name": "Chichewa", "Val": "ny"},
{"Name": "Chinese Traditional", "Val": "zh-TW"},
{"Name": "Croatian", "Val": "hr"},
{"Name": "Czech", "Val": "cs"},
{"Name": "Danish", "Val": "da"},
{"Name": "Dutch", "Val": "nl"},
{"Name": "English", "Val": "en"},
{"Name": "Esperanto", "Val": "eo"},
{"Name": "Estonian", "Val": "et"},
{"Name": "Filipino", "Val": "tl"},
{"Name": "Finnish", "Val": "fi"},
{"Name": "French", "Val": "fr"},
{"Name": "Galician", "Val": "gl"},
{"Name": "Georgian", "Val": "ka"},
{"Name": "German", "Val": "de"},
{"Name": "Greek", "Val": "el"},
{"Name": "Gujarati", "Val": "gu"},
{"Name": "Haitian Creole", "Val": "ht"},
{"Name": "Hausa", "Val": "ha"},
{"Name": "Hebrew", "Val": "iw"},
{"Name": "Hindi", "Val": "hi"},
{"Name": "Hmong", "Val": "hmn"},
{"Name": "Hungarian", "Val": "hu"},
{"Name": "Icelandic", "Val": "is"},
{"Name": "Igbo", "Val": "ig"},
{"Name": "Indonesian", "Val": "id"},
{"Name": "Irish", "Val": "ga"},
{"Name": "Italian", "Val": "it"},
{"Name": "Japanese", "Val": "ja"},
{"Name": "Javanese", "Val": "jw"},
{"Name": "Kannada", "Val": "kn"},
{"Name": "Kazakh", "Val": "kk"},
{"Name": "Khmer", "Val": "km"},
{"Name": "Korean", "Val": "ko"},
{"Name": "Lao", "Val": "lo"},
{"Name": "Latin", "Val": "la"},
{"Name": "Latvian", "Val": "lv"},
{"Name": "Lithuanian", "Val": "lt"},
{"Name": "Macedonian", "Val": "mk"},
{"Name": "Malagasy", "Val": "mg"},
{"Name": "Malay", "Val": "ms"},
{"Name": "Malayalam", "Val": "ml"},
{"Name": "Maltese", "Val": "mt"},
{"Name": "Maori", "Val": "mi"},
{"Name": "Marathi", "Val": "mr"},
{"Name": "Mongolian", "Val": "mn"},
{"Name": "Myanmar (Burmese)", "Val": "my"},
{"Name": "Nepali", "Val": "ne"},
{"Name": "Norwegian", "Val": "no"},
{"Name": "Persian", "Val": "fa"},
{"Name": "Polish", "Val": "pl"},
{"Name": "Portuguese", "Val": "pt"},
{"Name": "Punjabi", "Val": "ma"},
{"Name": "Romanian", "Val": "ro"},
{"Name": "Russian", "Val": "ru"},
{"Name": "Serbian", "Val": "sr"},
{"Name": "Sesotho", "Val": "st"},
{"Name": "Sinhala", "Val": "si"},
{"Name": "Slovak", "Val": "sk"},
{"Name": "Slovenian", "Val": "sl"},
{"Name": "Somali", "Val": "so"},
{"Name": "Spanish", "Val": "es"},
{"Name": "Sudanese", "Val": "su"},
{"Name": "Swahili", "Val": "sw"},
{"Name": "Swedish", "Val": "sv"},
{"Name": "Tajik", "Val": "tg"},
{"Name": "Tamil", "Val": "ta"},
{"Name": "Telugu", "Val": "te"},
{"Name": "Thai", "Val": "th"},
{"Name": "Turkish", "Val": "tr"},
{"Name": "Ukrainian", "Val": "uk"},
{"Name": "Urdu", "Val": "ur"},
{"Name": "Uzbek", "Val": "uz"},
{"Name": "Vietnamese", "Val": "vi"},
{"Name": "Welsh", "Val": "cy"},
{"Name": "Yiddish", "Val": "yi"},
{"Name": "Yoruba", "Val": "yo"},
{"Name": "Zulu", "Val": "zu"}
]
`

var langList []LangItem

func init() {
	err := json.Unmarshal([]byte(rawLangList), &langList)
	if err != nil {
		log.Println("decode lang list err:", err)
	} else {
		log.Println("decoded lang items count:", len(langList))
	}
}
