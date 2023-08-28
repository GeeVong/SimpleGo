package common

import (
	"fmt"
	logs "github.com/liangdas/mqant/log/beego"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var HttpServers = make(map[string]*HttpServer)
var httpMutex sync.RWMutex //

type HttpServerConfig struct {
	group   string
	minPort int
	maxPort int
}

var httpConfig = make(map[string]*HttpServerConfig)

type HttpServer struct {
	server          *http.Server
	mux             *http.ServeMux
	port            int
	group           string
	callBackDealFun map[string]string
}

func InitHttpServerConfig(group string, minPort, maxPort int) {
	_httpConfig := &HttpServerConfig{
		minPort: minPort,
		maxPort: maxPort,
	}
	httpConfig[group] = _httpConfig
}

func startServer(group string) *HttpServer {
	defer httpMutex.RUnlock()
	httpMutex.RLock()
	_httpConfig, ok := httpConfig[group]
	if !ok || _httpConfig == nil {
		lgames.LogError("HTTPSERVER", "http server fail !  httpConfig length : %d , group:%s", len(httpConfig), group)
		return nil
	}
	if _, ok := HttpServers[group]; !ok {
		for i := _httpConfig.minPort; i < _httpConfig.maxPort; i++ {
			if !isCanListen(i) {
				continue
			}
			mux := http.NewServeMux()
			mux.HandleFunc("/default", httpHandle)
			server := &http.Server{Addr: ":" + strconv.Itoa(i), Handler: mux}
			go server.ListenAndServe()
			HttpServers[group] = &HttpServer{server: server, mux: mux, port: i, group: group, callBackDealFun: make(map[string]string)}
			lgames.LogInfo("HTTPSERVER", "http server success !  listen port : %d", i)
			break
		}
	}
	return HttpServers[group]
}

// 注册回调地址时记得解注册 UnRegisterCallBackUrl
func StartServerAndRegisterCallBack(group string, urls string, callBack func(*HttpBackParams) *HttpBackParams) bool {
	httpServers := startServer(group)
	if httpServers == nil {
		return false
	}
	return httpServers.registerCallBackUrl(urls, callBack)
}

func isCanListen(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		l.Close()
		return true
	}
	return false
}

func httpHandle(resp http.ResponseWriter, req *http.Request) {
	logs.Info("http server start .....")
}

func (self *HttpServer) registerCallBackUrl(urls string, callBack func(*HttpBackParams) *HttpBackParams) bool {
	if callBack == nil {
		lgames.LogError("HTTPSERVER", "http 回调方法为空")
		return false
	}
	_, ok := self.callBackDealFun[urls]
	if ok {
		lgames.LogError("HTTPSERVER", "http 回调方法已经被注册")
		return false
	}
	f := func(resp http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		body, _ := ioutil.ReadAll(req.Body)
		date := make(map[string]string)
		for k, arr := range req.PostForm {
			if len(arr) > 0 {
				date[k] = arr[0]
			}
		}
		params := &HttpBackParams{
			Body:     body,
			PostDate: date,
		}
		httpBackParams := callBack(params)
		resp.Write(httpBackParams.Response)
	}
	self.callBackDealFun[urls] = urls
	self.mux.HandleFunc("/"+urls, f)

	lgames.LogInfo("HTTPSERVER", "http 注册回调方法: %s", urls)
	return true
}

func UnRegisterCallBackUrl(group string, urls string) {
	defer httpMutex.RUnlock()
	httpMutex.RLock()
	if httpSrv, ok := HttpServers[group]; ok {
		if _, ok := httpSrv.callBackDealFun[urls]; ok {
			delete(httpSrv.callBackDealFun, urls)
			lgames.LogInfo("HTTPSERVER", "http 解除回调方法: %s", urls)
		}
		if len(httpSrv.callBackDealFun) <= 0 {
			port := httpSrv.port
			httpSrv.server.Close()
			//httpSrv.server.Shutdown(context.Background())
			delete(HttpServers, group)
			lgames.LogInfo("HTTPSERVER", "http 端口释放: %d", port)
		}
	}
}

type HttpBackParams struct {
	Body     []byte
	PostDate map[string]string
	Response []byte
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	p := &HttpClient{
		client: &http.Client{},
	}
	return p
}

func (self *HttpClient) AsyncHttpPost(targetHostUrl string, date url.Values) {
	if targetHostUrl == "" {
		return
	}

	_, err := self.client.PostForm(targetHostUrl, date)
	if err != nil {
		lgames.LogError("http post", "post 请求失败", err)
	}
}

func (self *HttpClient) SyncHttpPost(targetHostUrl string, date url.Values) []byte {
	if targetHostUrl == "" {
		return nil
	}

	resp, err := self.client.PostForm(targetHostUrl, date)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	return body
}

func (self *HttpClient) SyncHttpGet(urls string) []byte {
	//提交请求
	request, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		lgames.LogError("http get", "get 请求失败", err)
	}

	//处理返回结果
	response, err2 := self.client.Do(request)
	if err2 != nil {
		lgames.LogError("http get", "get 请求失败", err2)
	}
	defer response.Body.Close()
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		lgames.LogError("http get", "ioutil.ReadAll failed ,err:%v", err1)
	}
	return body
}

func (self *HttpClient) AsyncHttpGet(urls string) {

}
