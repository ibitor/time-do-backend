package core

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"sort"
	"time"
	"timedo/config"
	"timedo/dao"
	"timedo/sql"
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

func getLastDay(item sql.Item) int {
	proc := item.ProduceDate
	now := time.Now()
	var subDay = int(now.Sub(proc).Hours() / 24)
	lastDay := item.SafeDay - subDay
	return lastDay
}

func PushAllItems() {
	items := dao.GetAllItem()
	sort.Slice(items, func(i, j int) bool {
		return getLastDay(items[i]) < getLastDay(items[j])
	})
	contents := ""
	for i := range items {
		safeDay := items[i].SafeDay
		itemName := items[i].Name
		lastDay := getLastDay(items[i])
		contents += fmt.Sprintf("%s 还剩 %d 天, 保质期共 %d 天 \n", itemName, lastDay, safeDay)
	}
	PushWX(wxConfig.CorpId, wxConfig.AgentId, wxConfig.AgentSecret, contents)
	log.Println("done=============")
}
