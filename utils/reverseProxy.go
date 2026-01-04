package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type userIdkey struct {}
var UserIdKey=userIdkey{}

func ReverseProxy(targetBaseUrl string ,pathPrefix string) http.HandlerFunc{
	target,err:=url.Parse(targetBaseUrl)
	if err!=nil{
		fmt.Println("Error parsing target url")
		return nil
	}
	proxy:=httputil.NewSingleHostReverseProxy(target)
	orginalDirector:=proxy.Director
	proxy.Director=func(r *http.Request) {
		orginalDirector(r)
		orginalPath:=r.URL.Path
		stripedPath:=strings.TrimPrefix(orginalPath,pathPrefix)
		r.URL.Host=target.Host
		r.URL.Path=target.Path + stripedPath
		r.Host=target.Host

		if userId,ok:=r.Context().Value(UserIdKey).(string);ok{
			r.Header.Set("X-User-Id",userId)
		}
	}
	return proxy.ServeHTTP
}