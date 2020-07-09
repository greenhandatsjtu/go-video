package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	wd    string // work dir
	help  bool   // show help
	files []string
)

func main() {

	// get command line parameters
	currentWd, _ := os.Getwd()
	flag.BoolVar(&help, "h", false, "Show help.")
	flag.StringVar(&wd, "P", currentWd, "Choose working dir.")
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	projectPath := os.Getenv("GOPATH") + "/src/github.com/greenhandatsjtu/video/video-frontend/dist/"
	index := "index.html"
	fileInfoList, _ := ioutil.ReadDir(wd)
	for _, file := range fileInfoList {
		if file.IsDir() {
			continue
		}
		// 过滤指定格式
		ext, _ := regexp.Compile("(.mp4|.mkv|.avi|.flv)$")
		if ext.MatchString(file.Name()) {
			files = append(files, file.Name())
		}
	}
	log.Println("Serve at  ", wd)
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFS("/video/", http.Dir(wd))
	router.LoadHTMLFiles(projectPath + "index.html")
	router.Static("/js", projectPath+"js")
	router.Static("/css", projectPath+"css")
	router.Static("/img", projectPath+"img")
	router.StaticFile("/favicon.ico", projectPath+"favicon.ico")

	router.GET("/", func(context *gin.Context) {
		context.HTML(200, index, nil)
	})

	router.GET("/list", func(context *gin.Context) {
		context.JSON(200, gin.H{"files": files})
	})

	// Listen and serve on 0.0.0.0:8080
	_ = router.Run()
}
