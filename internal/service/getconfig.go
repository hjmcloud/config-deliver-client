package service

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
)

// 获取配置文件
func GetConfig(ctx context.Context, taskName string, taskFile string, localDir string) (err error) {
	g.Log().Debug(ctx, "start task", taskName, taskFile, localDir)
	apiServer, err := g.Cfg().Get(ctx, "apiServer")
	if err != nil {
		return err
	}
	g.Log().Debug(ctx, "apiServer", apiServer)

	listUrl := apiServer.String() + "?id=" + taskName
	g.Log().Debug(ctx, "listUrl", listUrl)
	// localDir, err := g.Cfg().Get(ctx, "localDir")
	// if err != nil {
	// 	return err
	// }
	// g.Log().Debug(ctx, "localDir", localDir)
	if taskFile == "*" {
		type Result struct {
			Code    int      `json:"code"`
			Message string   `json:"message"`
			Data    []string `json:"data"`
		}
		var result *Result
		g.Client().GetVar(ctx, listUrl).Scan(&result)
		g.Log().Debug(ctx, "result", result.Code, result.Message, result.Data)
		if result.Code != 0 {
			g.Log().Error(ctx, "result", result.Code, result.Message)
			return err
		}

		var fileList []string = result.Data
		var runFlag bool = false
		for _, file := range fileList {
			update, err := DownloadConfig(ctx, taskName, file, localDir)
			if err != nil {
				return err
			}
			if update {
				runFlag = true
			}
		}
		if runFlag {
			g.Log().Info(ctx, "run shell", taskName, taskFile)
			// 判断是否有shell文件
			shellFile := localDir + "/run.sh"
			if gfile.Exists(shellFile) {
				g.Log().Info(ctx, "run shell", taskName, taskFile, shellFile)
				// g.Exec(ctx, shellFile)
				gproc.ShellExec("sh " + shellFile)
			}
			// g.Exec(ctx, "sh", "-c", "sh "+taskName+".sh")
		} else {
			g.Log().Info(ctx, "no update, no run shell", taskName, taskFile)
		}

	} else {
		update, err := DownloadConfig(ctx, taskName, taskFile, localDir)
		if err != nil {
			g.Log().Error(ctx, "task", err)
			return err
		}
		if update {
			g.Log().Info(ctx, "task", taskName, taskFile, "update")
		} else {
			g.Log().Info(ctx, "task", taskName, taskFile, "no update")
		}
	}
	return nil
}
func DownloadConfig(ctx context.Context, taskName string, taskFile string, localDir string) (update bool, err error) {
	g.Log().Debug(ctx, "start download config file", taskName, taskFile, localDir)
	// localDir, err := g.Cfg().Get(ctx, "localDir")
	// if err != nil {
	// 	return false, err
	// }
	// g.Log().Debug(ctx, "localDir", localDir)
	apiServer, err := g.Cfg().Get(ctx, "apiServer")
	if err != nil {
		return false, err
	}
	g.Log().Debug(ctx, "apiServer", apiServer)
	md5url := apiServer.String() + "?id=" + taskName + "&filename=" + taskFile
	g.Log().Debug(ctx, "md5url", md5url)
	downUrl := apiServer.String() + "?id=" + taskName + "&filename=" + taskFile + "&dl=true"
	type Result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
	var result *Result
	g.Client().GetVar(ctx, md5url).Scan(&result)
	g.Log().Debug(ctx, "result", result.Code, result.Message, result.Data)
	if result.Code != 0 {
		g.Log().Error(ctx, "result", result.Code, result.Message)
		return false, err
	}
	// 获取本地文件md5
	localMd5, err := gmd5.EncryptFile(localDir + "/" + taskFile)
	if err != nil {
		g.Log().Info(ctx, "get local md5 error", err)
	}
	g.Log().Debug(ctx, "localMd5", localMd5)
	// 比较md5
	if localMd5 != result.Data {
		g.Log().Info(ctx, "md5 not equal")
		// 下载文件
		if r, err := g.Client().Get(ctx, downUrl); err != nil {
			panic(err)
		} else {
			defer r.Close()
			err := gfile.PutBytes(localDir+"/"+taskFile, r.ReadAll())
			if err != nil {
				g.Log().Info(ctx, "put bytes error", err)
			}
		}
	} else {
		g.Log().Info(ctx, "md5 equal")
		return false, nil
	}
	// 获取下载文件md5
	downMd5, err := gmd5.EncryptFile(localDir + "/" + taskFile)
	if err != nil {
		g.Log().Info(ctx, "get down md5 error", err)
	}
	g.Log().Debug(ctx, "downMd5", downMd5)
	// 比较md5
	if downMd5 != result.Data {
		g.Log().Info(ctx, "md5 not equal")
		return false, err
	}
	return true, nil
}
