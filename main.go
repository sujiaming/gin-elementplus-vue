package main

import (
	"embed"
	"net/http"
	"strings"
	"wanworld/database"
	"wanworld/router"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	// 初始化数据库
	database.InitDB()
	r := gin.Default()

	// 添加 CORS 中间件（解决跨域问题）
	// r.Use(corsMiddleware())

	// 静态资源路由
	r.GET("/assets/*filepath", serveStatic)

	// 创建 API 路由组并注册路由
	apiGroup := r.Group("/api")
	router.SetupAPIRouter(apiGroup)

	// 处理前端路由（Vue Router 的 history 模式）
	// 兜底路由：优先处理未匹配的 API 请求
	r.NoRoute(func(c *gin.Context) {
		// 如果是 API 请求，返回 JSON 错误
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "API endpoint not found",
				"path":  c.Request.URL.Path,
			})
			return
		}

		// 否则交给前端处理（Vue Router）
		serveStatic(c)
	})

	r.Run(":8080")
}

func serveStatic(c *gin.Context) {
	path := c.Request.URL.Path
	if path == "/" {
		path = "static/index.html"
	} else {
		path = "static" + path
	}

	data, err := staticFiles.ReadFile(path)
	if err != nil {
		// 返回 index.html 以支持前端路由
		data, _ = staticFiles.ReadFile("static/index.html")
	}

	// 手动设置 Content-Type
	contentType := ""
	switch {
	case strings.HasSuffix(path, ".css"):
		contentType = "text/css"
	case strings.HasSuffix(path, ".js"):
		contentType = "application/javascript"
	case strings.HasSuffix(path, ".html"):
		contentType = "text/html"
	case strings.HasSuffix(path, ".png"):
		contentType = "image/png"
	default:
		contentType = http.DetectContentType(data)
	}

	c.Data(http.StatusOK, contentType, data)
}

// CORS 中间件配置
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
