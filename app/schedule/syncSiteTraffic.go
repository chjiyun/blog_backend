package schedule

import (
	"blog_backend/app/service/siteTrafficService"
	"blog_backend/config"
)

func (s Schedule) SyncSiteTraffic() MySchedule {
	task := MySchedule{
		Cron:      "0 0 0 * * ?",
		Immediate: true,
	}
	task.Task = func() {
		err := siteTrafficService.SyncSiteTraffic(false)
		if err != nil {
			config.Logger.Error(err)
			return
		}
		config.Logger.Info("siteTraffic统计数据已同步至redis")
	}
	return task
}
