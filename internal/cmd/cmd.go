package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	Main = gcmd.Command{
		Name:  "config-deliver-client",
		Usage: "config-deliver-client",
		Brief: "start config deliver client",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			tasks, err := g.Cfg().Get(ctx, "tasks")
			if err != nil {
				return
			}
			gutil.Dump(tasks)

			// 等待
			select {}

		},
	}
	// Task = gcmd.Command{
	// 	Name:  "task",
	// 	Usage: "task",
	// 	Brief: "start http server",
	// 	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
	// 		g.Log().Info(ctx, "config-deliver-client start!")
	// 		// gcron.SetLogger(glog.New())

	// 		tasks, err := g.Cfg().Get(ctx, "tasks")
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// 遍历tasks数组
	// 		for _, task := range tasks.Slice() {
	// 			var taskName = task.(map[string]interface{})["name"].(string)
	// 			var taskFile = task.(map[string]interface{})["file"].(string)
	// 			var taskCorn = task.(map[string]interface{})["corn"].(string)
	// 			var localDir = task.(map[string]interface{})["localDir"].(string)
	// 			var isRunOnStart = task.(map[string]interface{})["isRunOnStart"].(bool)

	// 			// 增加任务
	// 			_, err = gcron.AddSingleton(ctx, taskCorn, func(ctx context.Context) {
	// 				g.Log().Debug(ctx, "task", taskName, taskFile, taskCorn, localDir, isRunOnStart)
	// 				service.GetConfig(ctx, taskName, taskFile, localDir)
	// 			}, taskName)
	// 			if err != nil {
	// 				g.Log().Error(ctx, "task", err)
	// 				return err
	// 			}
	// 			g.Log().Info(ctx, "add new task", taskName, taskFile, taskCorn)
	// 			// 如果是启动时执行
	// 			if isRunOnStart {
	// 				g.Log().Info(ctx, "run task on start", taskName, taskFile, taskCorn)
	// 				service.GetConfig(ctx, taskName, taskFile, localDir)
	// 			}
	// 		}
	// 		// 获取任务列表
	// 		entries := gcron.Entries()
	// 		for k, v := range entries {
	// 			g.Log().Debug(ctx, "task", k, v.Name, v.Time)
	// 		}
	// 		// service.DownloadConfig(ctx, "238FDEDC-20DC-B06E-1D32-AD511C637A23", "cert.pem")
	// 		select {}
	// 		// return nil
	// 	},
	// }
)
