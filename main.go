package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/greenhandatsjtu/video/statik"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

var (
	wd    string // work dir
	port  int    // port
	help  bool   // show help
	files []string
)

func main() {
	index := "<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content=\"IE=edge\"><meta name=viewport content=\"width=device-width,initial-scale=1\"><link rel=icon href=/static/favicon.ico><title>video-frontend</title><link rel=stylesheet href=\"https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900\"><link rel=stylesheet href=https://cdn.jsdelivr.net/npm/@mdi/font@latest/css/materialdesignicons.min.css><link href=/static/css/chunk-vendors.a00cf1b2.css rel=preload as=style><link href=/static/js/app.2f621bf5.js rel=preload as=script><link href=/static/js/chunk-vendors.68b792e2.js rel=preload as=script><link href=/static/css/chunk-vendors.a00cf1b2.css rel=stylesheet></head><body><noscript><strong>We're sorry but video-frontend doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=/static/js/chunk-vendors.68b792e2.js></script><script src=/static/js/app.2f621bf5.js></script></body></html>"

	//static files embed in
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// get command line parameters
	currentWd, _ := os.Getwd()
	flag.BoolVar(&help, "h", false, "Show help.")
	flag.StringVar(&wd, "P", currentWd, "Specify working dir.")
	flag.IntVar(&port, "p", 8080, "Specify port.")
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

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
	log.Println("Serve at:	", wd)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFS("/video", http.Dir(wd))
	router.StaticFS("/static", statikFS)

	// index
	router.GET("/", func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.String(200, index)
	})

	// return video list
	router.GET("/list", func(context *gin.Context) {
		context.JSON(200, gin.H{"files": files})
	})

	//redirect when refresh
	router.GET("/play", func(context *gin.Context) {
		context.Redirect(302, "/")
	})

	url := "localhost:" + strconv.Itoa(port)
	log.Println("Opening " + url)

	//Open url in browser
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}

	// Listen and serve on 0.0.0.0:8080
	_ = router.Run(":" + strconv.Itoa(port))
}
