package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	. "quizbot/config"
	. "quizbot/internal/repository/cache"
	. "quizbot/internal/service"
)

type Args struct {
	Model        string `form:"model" json:"model" binding:"omitempty,oneof=deepseek tongyi"`
	Language     string `form:"language" json:"language" binding:"omitempty,oneof=go javascript java python c++"`
	QuestionType uint   `form:"type" json:"type" binding:"omitempty,oneof=1 2"`
	Keyword      string `form:"keyword" json:"keyword" binding:"required"`
}

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"aiRes"`
}

type Log struct {
	StartTime string  `json:"aiStartTime"`
	EndTime   string  `json:"aiEndTime"`
	CostTime  float64 `json:"aiCostTime"`
	Req       Args    `json:"aiReq"`
	Res       Reply   `json:"aiRes"`
}

// middleware for time measurement and log storage
func TimeIt(c *gin.Context) {
	// start time
	tic := time.Now()
	// higher-order function, handling the args and prepare for the reply
	c.Next()
	toc := time.Now()
	duration := toc.Sub(tic)
	reply, ok := c.Get("REPLY")
	if !ok {
		c.JSON(http.StatusInternalServerError, Reply{http.StatusInternalServerError, "no reply", nil})
		return
	}
	args, ok := c.Get("ARGS")
	if !ok {
		c.JSON(http.StatusInternalServerError, Reply{http.StatusInternalServerError, "no args", nil})
		return
	}
	// log storage
	SaveToJSONFile(Log{
		StartTime: tic.Format("2006-01-02 15:04:05"),
		EndTime:   toc.Format("2006-01-02 15:04:05"),
		CostTime:  duration.Seconds(),
		Req:       args.(Args),
		Res:       reply.(Reply),
	})
	log.Printf("Path: %s, Duration: %v, Status: %d", c.FullPath(), duration, reply.(Reply).Code)
	// returned reply
	c.JSON(reply.(Reply).Code, reply)
}

func setDefaults(a *Args) {
	if a.Model == "" {
		a.Model = "tongyi"
	}
	if a.Language == "" {
		a.Language = "go"
	}
	if a.QuestionType == 0 {
		a.QuestionType = 1
	}
}

// do not use c.JSON() in this function, use c.Set() instead
func CreateQuestion(c *gin.Context) {
	var args Args
	// received arguments
	if err := c.ShouldBind(&args); err != nil {
		c.Set("REPLY", Reply{-1, err.Error(), ""})
		c.Abort()
		return
	}
	// set default values
	setDefaults(&args)
	c.Set("ARGS", args)
	// new a large language model and call it
	llm := NewGenerator(args.Model)
	q, err := llm.Generate(args.Keyword, args.Language, QuestionType(args.QuestionType))
	if err != nil {
		c.Set("REPLY", Reply{-1, err.Error(), ""})
		c.Abort()
		return
	}
	// returned reply
	c.Set("REPLY", Reply{0, "success", q})
}
