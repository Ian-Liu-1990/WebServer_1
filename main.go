package main

import (
    "net/http"
    "fmt"
    "runtime"
    "reflect"
    "encoding/base64"
    "time"
)

type handler struct{

}

func(*handler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func hello(w http.ResponseWriter, r*http.Request){
    fmt.Fprintln(w, "Hi")
}

func rootHello(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Hello Root")
}

func log(f http.HandlerFunc)( http.HandlerFunc){
    return func(w http.ResponseWriter, r* http.Request){
        name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
        fmt.Println("Handler function called -" + name)
        f.ServeHTTP(w, r)
    }
}

func loglog(f http.HandlerFunc)(http.HandlerFunc){
    return func(w http.ResponseWriter, r*http.Request){
        fmt.Println("First")
        f.ServeHTTP(w, r)
    }
}

func processForm(w http.ResponseWriter, r * http.Request){

    fmt.Fprintln(w, r.Body)
}

type funRefere func(int, int)

func impRefere(r funRefere)(funRefere){
    return func(a,b int){
        r(a, b)
    }
}

func setCookie(w http.ResponseWriter, r *http.Request){
    str := []byte("Hello Workd")
    msg := http.Cookie{
        Name: "flushCookie", //盡量不要空格
        Value: base64.URLEncoding.EncodeToString(str),
    }

    http.SetCookie(w, &msg)
}

func showCookie(w http.ResponseWriter, r *http.Request){
    c, err := r.Cookie("flushCookie")
    if err != nil{
        fmt.Fprintln(w, "Cookie Not Found")
    }else{
        rc := http.Cookie{
            Name: "flushCookie",
            MaxAge: -1,
            Expires: time.Unix(1, 0),
        }
        http.SetCookie(w, &rc)
        val, _ := base64.URLEncoding.DecodeString(c.Value)
        fmt.Fprintln(w, string(val))
    }
}

func main(){
    ha := &handler{}


    helloRoot := http.HandlerFunc(rootHello)

    Server := http.Server{
        Addr: "127.0.0.1:8080",
    }

    //mux := http.NewServeMux()


    //mux.Handle("123", http.HandlerFunc(hello)) //一個使用處理器，一個使用處理器函數 mux.HandleFunc("123", hello)

    //http.HandleFunc("/sdf", rootHello)

    http.Handle("/hello", log(loglog(helloRoot)))
    http.Handle("/ha", ha) // 我只要有實作handler介面功能，我就能傳遞實例化物件(動態連結)，http.Handle會自己判斷
    http.HandleFunc("/process", processForm)
    http.HandleFunc("/flushCookie", setCookie)
    http.HandleFunc("/showCookie", showCookie)
    Server.ListenAndServe()
}
