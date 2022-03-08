package bind

import "github.com/webview/webview"

func ImageViewer() interface{} {
	return func(src, alt string) error {
		w := webview.NewWindow(false, nil)
		w.SetSize(400, 580, webview.HintNone)
		w.Navigate("http://localhost:8080/image?src=" + src + "&alt=" + alt)
		w.Run()
		return nil
	}
}
