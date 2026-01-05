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

// func ReverseProxy(targetBaseUrl string ,pathPrefix string) http.HandlerFunc{
// 	target,err:=url.Parse(targetBaseUrl)
// 	if err!=nil{
// 		fmt.Println("Error parsing target url")
// 		return nil
// 	}
// 	proxy:=httputil.NewSingleHostReverseProxy(target)
// 	orginalDirector:=proxy.Director
// 	proxy.Director=func(r *http.Request) {
// 		orginalDirector(r)
// 		orginalPath:=r.URL.Path
// 		stripedPath:=strings.TrimPrefix(orginalPath,pathPrefix)
// 		fmt.Println("stripedPath",stripedPath)

// 		r.URL.Host=target.Host
// 		r.URL.Path=target.Path + stripedPath
// 		r.Host=target.Host
		
// 		fmt.Println("Outgoing:",r.URL.Path)

// 		if userId,ok:=r.Context().Value(UserIdKey).(string);ok{
// 			r.Header.Set("X-User-Id",userId)
// 		}
// 	}
// 	return proxy.ServeHTTP
// }
func ReverseProxy(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
    target, err := url.Parse(targetBaseUrl)
    if err != nil {
        panic(err) // fail fast, not silently
    }

    proxy := httputil.NewSingleHostReverseProxy(target)

    originalDirector := proxy.Director
    proxy.Director = func(r *http.Request) {
        originalDirector(r)

        fmt.Println("INCOMING:", r.URL.Path)

        strippedPath := strings.TrimPrefix(r.URL.Path, pathPrefix)

        r.URL.Scheme = target.Scheme
        r.URL.Host = target.Host
        r.URL.Path = "/api/v1" + strippedPath
        r.Host = target.Host

        fmt.Println("FORWARDED:", r.URL.String())
    }

    return proxy.ServeHTTP
}
