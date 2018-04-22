package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// 未承認だった場合
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		log.Info("authHandler.serveHTTP: 未認証です。")
	} else if err != nil {
		// 別の何らかのエラーが発生
		panic(err.Error())
	} else {
		// 成功した場合、ラップされたハンドラを呼び出す。
		h.next.ServeHTTP(w, r)
		log.Info("認証に成功しました。")
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}