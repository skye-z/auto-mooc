package global

import (
	"net/http"
	"net/url"
)

func SendPush(msg string) {
	if !GetBool("push.enable") {
		return
	}
	url := GetString("push.url") + url.QueryEscape("AutoMoot\t"+msg)
	req, _ := http.NewRequest("GET", url, nil)
	http.DefaultClient.Do(req)
}
