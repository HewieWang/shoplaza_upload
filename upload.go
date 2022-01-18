package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/tidwall/gjson"
  "os"
)
var domain=os.Args[1]
var img=os.Args[2]
var cookie=getcookie()

func getcookie()  string{
  file, err := os.Open("shoplaza.cookie")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   content, err := ioutil.ReadAll(file)
   return string(content)
}
func upimg(task_id string){

  url := "https://"+domain+"/admin/api/image-upload/schedule"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    // return
  }
  req.Header.Add("authority", domain)
  req.Header.Add("pragma", "no-cache")
  req.Header.Add("cache-control", "no-cache")
  req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
  req.Header.Add("accept", "application/json, text/plain, */*")
  req.Header.Add("sec-ch-ua-mobile", "?0")
  req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
  req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
  req.Header.Add("sec-fetch-site", "same-origin")
  req.Header.Add("sec-fetch-mode", "cors")
  req.Header.Add("sec-fetch-dest", "empty")
  req.Header.Add("referer", "https://"+domain+"/admin/smart_apps/base/collections/_new")
  req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
  req.Header.Add("cookie", cookie)

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    // return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    // return
  }
  // return string(body)
  if  gjson.Get(string(body), "finished").String()=="1" && gjson.Get(string(body), "status").String()=="2" && gjson.Get(string(body), "task_id").String()==task_id{
      fmt.Println(gjson.Get(string(body), "success.0").String())
  }else{
      upimg(task_id)
  }
  // fmt.Println(string(body))
}

func main() {

  url := "https://"+domain+"/admin/api/images/upload"
  method := "POST"

  payload := strings.NewReader(`{"urls":["`+img+`"]}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("authority", domain)
  req.Header.Add("pragma", "no-cache")
  req.Header.Add("cache-control", "no-cache")
  req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
  req.Header.Add("accept", "application/json, text/plain, */*")
  req.Header.Add("content-type", "application/json")
  req.Header.Add("sec-ch-ua-mobile", "?0")
  req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
  req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
  req.Header.Add("origin", "https://"+domain)
  req.Header.Add("sec-fetch-site", "same-origin")
  req.Header.Add("sec-fetch-mode", "cors")
  req.Header.Add("sec-fetch-dest", "empty")
  req.Header.Add("referer", "https://"+domain+"/admin/smart_apps/base/collections/_new")
  req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
  req.Header.Add("cookie", cookie)

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  task_id:=gjson.Get(string(body), "task_id").String()

  upimg(task_id)
}
