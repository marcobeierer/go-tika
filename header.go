package tika

import "net/http"

type Header struct {
	http.Header
}

func NewHeader() Header {
	return Header{}
}

func (qq Header) AcceptText() Header {
	qq.Add("Accept", "text/plain")
	return qq
}

// OCRLanguage accepts string in format "eng+deu+spa"
// language codes are in ISO-639-2 format
func (qq Header) SetOCRLanguage(language string) Header {
	qq.Add("X-Tika-OCRLanguage", language)
	return qq
}
