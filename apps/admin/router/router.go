package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Router is a app level router.
type Router struct {
	*gin.Engine
}

// New initializes new gin router.
func New() *Router {
	gin.ForceConsoleColor()
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/info", func(c *gin.Context) {
		password, err := ioutil.ReadFile("/data/PASSWORD")
		if err != nil {
			fmt.Println(err)
		}

		userName, err := ioutil.ReadFile("/data/USER_NAME")
		if err != nil {
			fmt.Println(err)
		}

		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":          "admin",
				"HOST_IP_ADDR":  string(value),
				"host":          c.Request.Host,
				"USER_NAME":     string(userName),
				"PASSWORD":      string(password),
				"url":           c.Request.RequestURI,
				"USER_NAME_ENV": os.Getenv("USER_NAME_ENV"),
				"PASSWORD_ENV":  os.Getenv("PASSWORD_ENV"),
			},
		})
	})

	r.GET("/admin", func(c *gin.Context) {
		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":         "admin",
				"host":         c.Request.Host,
				"url":          c.Request.RequestURI,
				"HOST_IP_ADDR": string(value),
			},
		})

	})

	r.GET("/manager", func(c *gin.Context) {
		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		resp, err := http.Get("http://localhost:5000/info")
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var target interface{}

		err = json.Unmarshal(body, &target)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":         "manager",
				"host":         c.Request.Host,
				"url":          c.Request.RequestURI,
				"HOST_IP_ADDR": string(value),
			},
			"hostInfo": target,
		})
	})

	r.GET("/auth", func(c *gin.Context) {
		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		resp, err := http.Get("http://localhost:4000/info")
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var target interface{}

		err = json.Unmarshal(body, &target)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":         "manager",
				"host":         c.Request.Host,
				"url":          c.Request.RequestURI,
				"HOST_IP_ADDR": string(value),
			},
			"hostInfo": target,
		})
	})

	return &Router{r}
}
