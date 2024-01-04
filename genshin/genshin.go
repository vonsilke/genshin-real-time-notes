package genshin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"resin/helper"
	"time"
)

type GenshinAttendanceRewards struct {
	Status   string `json:"status"`
	Progress int    `json:"progress"`
}

type GenshinTaskRewards struct {
	Status string `json:"status"`
}

type GenshinDailyTask struct {
	TotalNum                  int                        `json:"total_num"`
	FinishedNum               int                        `json:"finished_num"`
	IsExtraTaskRewardReceived bool                       `json:"is_extra_task_reward_received"`
	TaskRewards               []GenshinTaskRewards       `json:"task_rewards"`
	AttendanceRewards         []GenshinAttendanceRewards `json:"attendance_rewards"`
	AttendanceVisible         bool                       `json:"attendance_visible"`
}

type GenshinTransformerRecovery struct {
	Day     int  `json:"Day"`
	Hour    int  `json:"Hour"`
	Minute  int  `json:"Minute"`
	Second  int  `json:"Second"`
	Reached bool `json:"reached"`
}

type GenshinTransformer struct {
	Obtained     bool                       `json:"obtained"`
	RecoveryTime GenshinTransformerRecovery `json:"recovery_time"`
	Wiki         string                     `json:"wiki"`
	Noticed      bool                       `json:"noticed"`
	LatestJobID  int                        `json:"latest_job_id"`
}

type GenshinExpedition struct {
	AvatarSideIcon string `json:"avatar_side_icon"`
	Status         string `json:"status"`
	RemainedTime   string `json:"remained_time"`
}

type GenshinData struct {
	CurrentResin              int                    `json:"current_resin"`
	MaxResin                  int                    `json:"max_resin"`
	ResinRecoveryTime         string                 `json:"resin_recovery_time"`
	FinishedTaskNum           int                    `json:"finished_task_num"`
	TotalTaskNum              int                    `json:"total_task_num"`
	IsExtraTaskRewardReceived bool                   `json:"is_extra_task_reward_received"`
	RemainResinDiscountNum    int                    `json:"remain_resin_discount_num"`
	ResinDiscountNumLimit     int                    `json:"resin_discount_num_limit"`
	CurrentExpeditionNum      int                    `json:"current_expedition_num"`
	MaxExpeditionNum          int                    `json:"max_expedition_num"`
	Expeditions               []GenshinExpedition    `json:"expeditions"`
	CurrentHomeCoin           int                    `json:"current_home_coin"`
	MaxHomeCoin               int                    `json:"max_home_coin"`
	HomeCoinRecoveryTime      string                 `json:"home_coin_recovery_time"`
	CalendarURL               string                 `json:"calendar_url"`
	Transformer               string                 `json:"transformer"`
	DailyTask                 GenshinDailyTask       `json:"daily_task"`
	ArchonQuestProgress       map[string]interface{} `json:"archon_quest_progress"`
}

type GenshinResponse struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    GenshinData `json:"data"`
}

func LoadJSON(reader io.Reader) GenshinResponse {
	var cfg GenshinResponse
	bytesValue, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytesValue)
	json.Unmarshal(bytesValue, &cfg)
	return cfg
}

func GenerateDS() string {
	salt := "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	t := time.Now().Unix()
	r := helper.RandStringBytes(6)
	h := helper.GetMD5Hash(fmt.Sprintf("salt=%s&t=%d&r=%s", salt, t, r))
	return fmt.Sprintf("%d,%s,%s", t, r, h)
}

func GenerateRequest(server string, genshinUUID string, ltoken string, ltuid string) *http.Request {
	url := fmt.Sprintf("https://bbs-api-os.hoyolab.com/game_record/genshin/api/dailyNote?server=%s&role_id=%s", server, genshinUUID)
	rOrig, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	rOrig.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/14.2 Chrome/87.0.4280.141 Mobile Safari/537.36")
	rOrig.Header.Add("Accept", "application/json, text/plain, */*")
	rOrig.Header.Add("Accept-Language", "en-US,en;q=0.5")
	rOrig.Header.Add("Accept-Encoding", "gzip, deflate, br")
	rOrig.Header.Add("x-rpc-client_type", "5")
	rOrig.Header.Add("x-rpc-app_version", "1.5.0")
	rOrig.Header.Add("x-rpc-language", "en-us")
	rOrig.Header.Add("Origin", "https://act.hoyolab.com")
	rOrig.Header.Add("Connection", "keep-alive")
	rOrig.Header.Add("Referer", "https://act.hoyolab.com/")
	rOrig.Header.Add("Cookie", fmt.Sprintf("ltoken_v2=%s; ltuid_v2=%s", ltoken, ltuid))
	rOrig.Header.Add("Sec-Fetch-Dest", "empty")
	rOrig.Header.Add("Sec-Fetch-Mode", "cors")
	rOrig.Header.Add("Sec-Fetch-Site", "same-site")
	return rOrig
}