package core

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
	"timedo/config"
	"timedo/dao"
)

var wxConfig = config.ReadConfig()

func CronPush() {
	wxCron := cron.New()
	log.Println("======== start cron ========")
	err := wxCron.AddFunc(wxConfig.CronExp, PushAllItems)
	if err != nil {
		log.Panicln("error: wx cron")
	}
	wxCron.Start()
}

func PushAllItems() {
	items := dao.GetAllItem()
	contents := ""
	for i := range items {
		safeDay := items[i].SafeDay
		itemName := items[i].Name
		proc := items[i].ProduceDate
		now := time.Now()
		var subDay = int(now.Sub(proc).Hours() / 24)
		lastDay := items[i].SafeDay - subDay
		contents += fmt.Sprintf("%s 还剩 %d 天, 保质期共 %d 天 \n", itemName, lastDay, safeDay)
	}
	PushWX(wxConfig.CorpId, wxConfig.AgentId, wxConfig.AgentSecret, contents)
	log.Println("done=============")
}
