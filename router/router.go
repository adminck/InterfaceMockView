package router

import (
	API "InterfaceMockView/api"
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.Use(LoadTls())  // 打开就能玩https了
	// 跨域
	Router.Use(Cors())
	Router.Static("/static", "./dist/static")
	//Router.LoadHTMLGlob("./dist/index.html") //这是前台的index

	Router.NoMethod(NoRouteFunc())
	Router.NoRoute(NoRouteFunc())
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")

	//登陆以及鉴权路由
	InitHomeRouter(ApiGroup)
	InitAuthorityLoginRouter(ApiGroup) // 注册基础功能路由 不做鉴权
	InitDomainRouter(ApiGroup)
	InitApiInfoRouter(ApiGroup)
	InitApiJsonInfoRouter(ApiGroup)
	return Router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func ApiProxy(w http.ResponseWriter, r *http.Request,domain models.Domain) {
	var Scheme string
	if r.TLS != nil {
		Scheme = "https"
	} else {
		Scheme = "http"
	}
	var proxyurl = &url.URL{
		Scheme: Scheme,
		Host:   r.Host,
	}
	if domain.HostAgent != "" {
		proxyurl.Host = domain.HostAgent
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyurl)
	proxy.ServeHTTP(w, r)
}

func NoRouteFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiPath := c.Request.URL.Path
		ApiHost := c.Request.Host
		if ApiInfo, err := API.QueryApi(ApiPath, ApiHost); err != nil {
			if domain, err := API.QueryApiProxy(ApiHost); err != nil && domain.IsHostAgent {
				ApiProxy(c.Writer, c.Request,domain)
			} else {
				common.Result(common.ERROR, gin.H{}, fmt.Sprintf("接口未定义%v", err.Error()), c)
			}
			c.Abort()
			return
		} else {
			if JsonInfo, err := API.QueryApiJsonInfo(c, ApiInfo.ID);err != nil {
				if domain, _ := API.QueryApiProxy(ApiHost); domain.IsHostAgent {
					ApiProxy(c.Writer, c.Request,domain)
					c.Abort()
					return
				}
				common.Result(common.ERROR, gin.H{}, fmt.Sprintf("接口json获取失败%v", err), c)
				c.Abort()
				return
			}else {
				API.GetJsonData(c,JsonInfo.JsonContent)
			}
		}
		c.Abort()
		return
	}
}
