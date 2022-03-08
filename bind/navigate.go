package bind

import "github.com/webview/webview"

func ChangeNavigateBind(w webview.WebView) interface{} {
	return func(url string) error {
		w.Navigate(url)
		return nil
	}
}
