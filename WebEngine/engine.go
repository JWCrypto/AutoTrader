package WebEngine

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/gin-contrib/static"
	"github.com/tinyhui/CryptoTrader/AnalyseEngine"
)

type WebEngine interface {
	init()
	setupBasicEndpoints()
	Run()
}

type webEngine struct {
	router       *gin.Engine
	serverConfig ServerConfig
	staticConfig TemplateConfig
}

type webEngineBuilder struct {
	webEngine *webEngine
}

func NewWebEngineBuilder() *webEngineBuilder {
	engine := &webEngine{}
	return &webEngineBuilder{
		webEngine: engine,
	}
}

func (builder *webEngineBuilder) WithServerConfig(config ServerConfig) *webEngineBuilder {
	builder.webEngine.serverConfig = config
	return builder
}

func (builder *webEngineBuilder) WithStaticConfig(config TemplateConfig) *webEngineBuilder {
	builder.webEngine.staticConfig = config
	return builder
}

func (builder *webEngineBuilder) Build() *webEngine {
	builder.webEngine.init()
	builder.webEngine.setupBasicEndpoints()
	return builder.webEngine
}

func (e *webEngine) init() {
	e.router = gin.Default()
}

func (e *webEngine) setupBasicEndpoints() {
	engine := e.router

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// load static files under /public
	engine.Use(static.ServeRoot("/public", e.staticConfig.StaticRoot))

	// load html templates
	engine.LoadHTMLGlob(e.staticConfig.HtmlRoot)
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Auto Crypto Trading Platform",
		})
	})
}

func (e *webEngine) Run() {
	if e.serverConfig.Port == 0 {
		e.serverConfig.Port = 8000
	}
	addr := fmt.Sprintf("%s:%d", e.serverConfig.Location, e.serverConfig.Port)
	e.router.Run(addr)
}
