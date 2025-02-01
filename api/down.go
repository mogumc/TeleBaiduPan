// 网盘文件下载地址返回
// @author MoGuQAQ
// @version 1.0.1

package api

import (
	"TeleBaidu/config"
	"TeleBaidu/core"
	"TeleBaidu/global"
	"TeleBaidu/utils"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

func Down(fid int64) string {
	if global.Log == nil {
		global.Log = core.InitLogger()
	}
	acclink := config.Config.User.AccLink
	BDUSS := config.Config.User.Bduss
	if BDUSS == "" {
		global.Log.Warnf("未填写BDUSS!")
		return "未填写BDUSS"
	}
	fsid := strconv.FormatInt(int64(fid), 10)
	if acclink == "" {
		global.Log.Infof("当前处于本地解析模式")
		url := "http://110.242.69.43/api/filemetas?dlink=1&clienttype=17&rt=third&vip=2&fsids=[%22" + fsid + "%22]"
		res := utils.Get(url, "netdisk;Mo", "BDUSS="+BDUSS+";PANPSC=;BAIDUID=1;ndut_fmt="+utils.Getndut())

		var JsonData map[string]interface{}
		if json.Unmarshal([]byte(res), &JsonData) == nil {
			errno := JsonData["errno"].(float64)
			global.Log.Infof("请求 Fid->%s 返回状态码 errno->%d", fsid, int(errno))
			if errno != 0 {
				return "failed"
			} else {
				info, ok := JsonData["info"].([]interface{})
				if !ok || len(info) == 0 {
					global.Log.Errorf("info 为空")
					return "获取下载地址失败"
				}
				odlink, ok := JsonData["info"].([]interface{})[0].(map[string]interface{})["dlink"].(string)
				if !ok {
					global.Log.Errorf("百度返回数据异常")
					return "获取下载地址失败"
				}
				dl := strings.Replace(odlink, "d.pcs.baidu.com", "218.93.204.36/b/pcs.baidu.com", -1) + "&clienttype=17&channel=0&version=7.22.0.8&" + utils.Getrand(BDUSS)
				headResult := utils.Head(dl, config.Config.User.User_Agent, "")
				dlink := headResult["Location"]
				if dlink[0] == "" {
					return "获取下载地址失败"
				}
				return dlink[0]
			}
		} else {
			global.Log.Errorf("解析Json失败 Url->%s", url)
			return "获取下载地址失败"
		}
	} else {
		global.Log.Infof("当前处于远程解析模式")
		res := utils.Get(acclink, "netdisk;Mo", "")
		var JsonData map[string]interface{}
		if json.Unmarshal([]byte(res), &JsonData) == nil {
			code := JsonData["code"].(string)
			if code != "0" {
				global.Log.Errorf("加速链接无效 Url->%s", acclink)
				return "无效的加速链接"
			} else {
				pdata := "bduss=" + BDUSS + "&fid=" + fsid + "&ua=" + base64.StdEncoding.EncodeToString([]byte(config.Config.User.User_Agent))
				res = utils.Post(acclink, "KinhWeb", "", pdata)
				var JsonData map[string]interface{}
				if json.Unmarshal([]byte(res), &JsonData) == nil {
					errno, ok := JsonData["errno"].(float64)
					if !ok {
						global.Log.Warnf("请求的加速链接返回了无效数据")
						return "获取下载地址失败"
					}
					if errno != 0 {
						global.Log.Errorf("获取下载地址失败%s", res)
						return "获取下载地址失败"
					}
					dlink := JsonData["dlink"].(string)
					if dlink == "" {
						global.Log.Errorf("获取下载地址失败")
						return "获取下载地址失败"
					}
					global.Log.Infof("获取到地址 %s", dlink)
					return dlink
				} else {
					global.Log.Errorf("解析Json失败 Url->%s", acclink)
					return "获取下载地址失败"
				}
			}
		} else {
			global.Log.Errorf("解析Json失败 Url->%s", acclink)
			return "获取下载地址失败"
		}
	}
}
