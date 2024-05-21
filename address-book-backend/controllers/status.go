package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/JJDoneAway/addressbook/models"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	// "github.com/robfig/cron"
	"github.com/carlmjohnson/versioninfo"
)

const METRIC_HEALTHINESS = "healthiness_counter"

func init() {
	metric := ginmetrics.Metric{
		Type:        ginmetrics.Counter,
		Name:        METRIC_HEALTHINESS,
		Description: "Counts the successful healthiness checks",
		Labels:      []string{"status"},
	}
	ginmetrics.GetMonitor().AddMetric(&metric)
	// c := cron.New()
	// if ok := c.AddFunc(os.Getenv("MAGIC_MONITORING_INTERVAL"), func() { happinessCheck() }); ok != nil {
	// 	siam.ErrorLogger.Printf("Wasn't able to init cron job to run the 4+1 checks. %v", ok)
	// }
	// c.Start()
}

func AddStatus(router *gin.Engine) {
	router.GET("status/up", doUp)
	router.GET("status/health", doHealth)
	router.GET("status/build", getBuild)
}

// @Summary      Get build time stamp
// @Description  Will just tell when the app was build
// @ID           getBuild
// @Tags         status
// @Produce      json
// @Success      200 {string}  string    "timestamp"
// @Router       /status/build [GET]
func getBuild(c *gin.Context) {
	c.JSON(http.StatusOK, Version{
		Version:    versioninfo.Version,
		Revision:   versioninfo.Revision,
		DirtyBuild: versioninfo.DirtyBuild,
		LastCommit: versioninfo.LastCommit,
		BuildTime:  getBuildTimeStamp(),
	})

}

func getBuildTimeStamp() string {
	b, err := os.ReadFile("timestamp.txt")
	if err != nil {
		return "unknown"
	}
	return strings.Trim(string(b), "\n")
}

type Version struct {
	Version    string
	Revision   string
	DirtyBuild bool
	LastCommit time.Time
	BuildTime  string
}

// @Summary      Tell Up status
// @Description  Will just tell you if the app is upp and running
// @ID           upStatus
// @Tags         status
// @Produce      json
// @Success      200 {string}  string    "ok"
// @Router       /status/up [GET]
func doUp(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// @Summary      Tell the health status
// @Description  Will just tell you if the app is healthy
// @ID           healthStatus
// @Tags         status
// @Produce      json
// @Success      200 {string}  string    "ok"
// @Router       /status/health [GET]
func doHealth(c *gin.Context) {
	if ok := models.CheckDB(); ok != nil {
		log.Printf("DB is not available: %v", ok)
		c.JSON(500, fmt.Sprintf("DB is broken: %v", ok))
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func happinessCheck() {
	counter := ginmetrics.GetMonitor().GetMetric(METRIC_HEALTHINESS)
	if ok := models.CheckDB(); ok != nil {
		counter.Inc([]string{"red"})
		fmt.Printf("WARN | HAPPINESS | DB | is not working proper. Error was %v\n", ok)
		return
	}

	if _, ok := models.GetAllAddresses(); ok != nil {
		counter.Inc([]string{"red"})
		fmt.Printf("WARN | HAPPINESS | SELECT_ALL | is not working proper. Error was %v\n", ok)
		return
	}

	counter.Inc([]string{"green"})

}
