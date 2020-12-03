# HTTP
## １. HTTP請求
1. 請求Message

| Method(GET/POST) | URL | Version(1.1) | CRCL |
|:------:|:------:|:------:|:------:|
1. **Header標頭**
    * **表頭內容** : _[參閱HTTP協定]()_，**主要取得Client端文件類型，字元編碼，字元長度和Cookie等訊息**
    * **組成** : **鍵Key : 值Value**
    * **相關函數** : http.Request.Header型態 // 一個map[string][]string
    * **Header的方法** : .Get(), .Set(), .Value()等...
    
 2. Header與Body隔一個**空行**
 
 3. 請求方法Method
     1. **GET** : 指定表單資源**以?開頭**掛載在URL的後面，並以fidleName=value&表示並連接
        1. **相關函數-urlencoded編碼** : http.ParseForm()
            1. r.Form["formKey"] // 回傳url與表單鍵值對
            2. r.PostForm["formKey"] // 只回傳表單鍵值對
        2. **相關函數-multipart編碼** : http.ParseMultipartForm()   
            1. ~~r.Form["formKey"] // 回傳url與表單鍵值對~~
            2. r.MultipartFrom["formKey"] // 只回傳表單鍵值對
        
     2. **POST** : 指定表單資源**儲存在Body中**
        1. **相關函數-urlencoded編碼** : http.ParseForm()
             1. r.Form["formKey"] // 回傳url與表單鍵值對
             2. r.PostForm["formKey"] // 只回傳表單鍵值對
        2. **相關函數-multipart編碼** : http.ParseMultipartForm() 
             1. ~~r.Form["formKey"] // 回傳url與表單鍵值對~~
             2. r.MultipartFrom["formKey"] // 只回傳表單鍵值對
  4. 安全與冪等請求方法
  
## ２. HTTP回應
1. [一個狀態行1XX~5XX]()
    * **相關函數** : w.writeHeader(int32)
2. 0個或任意數量的回應Header
    * **w.Header()的方法** : .set(), .get(), ["Cookie"]
    * 例如: .set("location", "yahoo.com.tw")//重定向到其他網址
    * 例如: .set("Content-Type", "application/json")//回傳json資料
3. 一個空行;
4. 一個可選的Message Body