package bind

import "github.com/webview/webview"

func AllBindCollection(w webview.WebView) {
	w.Bind("goLog", GoLogBind())
	w.Bind("getString", GetStringBind())
	w.Bind("getFileList", FileBind())
	w.Bind("getComplexData", ComplexBind())
	w.Bind("getAllName", GetAllNameBind())
	w.Bind("getNameById", GetNameByIdBind())
	w.Bind("insertName", InsertBind())
	w.Bind("getUrlImage", UrlImageBind())
	w.Bind("getLocalImage", LocalImageBind())
	w.Bind("changeNavigate", ChangeNavigateBind(w))
	w.Bind("quit", QuitBind(w))
}
