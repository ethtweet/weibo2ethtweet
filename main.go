package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"weibo2ethtweet/model"
	"weibo2ethtweet/utils"

	"github.com/golang-module/carbon"
	"github.com/spf13/viper"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

var ipfs *shell.Shell

func main() {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
	_ = utils.Db.AutoMigrate()

	ipfs = shell.NewShell("https://cdn.ipfsscan.io")
	id := strconv.Itoa(0)
	for {
		var users []model.User
		if utils.Db.Get().Where("id > ?", id).Order("id asc").Limit(100).Find(&users).RowsAffected < 1 {
			break
		}

		for _, user := range users {
			id = user.ID

			var items []model.Weibo
			if utils.Db.Get().Where("tweet_id = '' AND user_id = ?", user.ID).
				Order("created_at desc").Limit(10).Find(&items).RowsAffected < 1 {
				continue
			}

			for _, tweet := range items {
				sendTweet(tweet)
			}
		}
	}
}

func tweetFiles(tweet model.Weibo) []string {
	var list []string
	filename := strings.Join([]string{
		carbon.Time2Carbon(tweet.CreatedAt).Format("Ymd"),
		tweet.ID,
	}, "_")

	pics := tweet.Pics
	if pics != "" {
		picsArr := strings.Split(pics, ",")

		if len(picsArr) > 1 {
			for index, pic := range picsArr {
				u, _ := url.Parse(pic)
				path := strings.Join([]string{
					"./weibo-crawler/weibo/" + tweet.UserID + "/img/原创微博图片/",
					filename,
					"_",
					strconv.Itoa(index + 1),
					filepath.Ext(u.Path),
				}, "")

				if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
					list = append(list, path)
				}
			}
		} else {
			u, _ := url.Parse(picsArr[0])
			path := strings.Join([]string{
				"./weibo-crawler/weibo/" + tweet.UserID + "/img/原创微博图片/",
				filename,
				filepath.Ext(u.Path),
			}, "")

			if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
				list = append(list, path)
			}
		}
	}

	videoUrl := tweet.VideoURL
	if videoUrl != "" {
		u, _ := url.Parse(videoUrl)
		path := strings.Join([]string{
			"./weibo-crawler/weibo/" + tweet.UserID + "/video/原创微博视频/",
			filename + filepath.Ext(u.Path),
		}, "")

		if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
			list = append(list, path)
		}
	}

	return list
}

func sendTweet(tweet model.Weibo) {
	//files := tweetFiles(tweet)
	//attachment := addToIpfs(files)

	attachment := strings.Trim(strings.Join([]string{tweet.Pics, tweet.VideoURL}, ","), ",")

	params := url.Values{}
	params.Set("content", tweet.Text)
	params.Set("attachment", attachment)
	params.Set("key", tweet.UserID)

	_, body, _ := gorequest.New().Post("http://127.0.0.1:8080/api/v0/tweet/release").
		Type("multipart").
		Send(params).
		End()

	if gjson.Get(body, "code").Int() != 0 {
		log.Println(tweet.ID + "发送失败")
	} else {
		utils.Db.Get().Model(&tweet).Select("TweetID").Updates(model.Weibo{TweetID: gjson.Get(body, "data.Id").String()})
		log.Println(strings.Repeat("=", 50), tweet.ID+" 发送成功")
	}

	//wg.Done()
}

func addToIpfs(files []string) string {
	var cids []string
	for _, filePath := range files {
		if fl, err := ioutil.ReadFile(filePath); err == nil {
			if cid, er := ipfs.Add(strings.NewReader(string(fl))); er == nil {
				cid = "https://cdn.ipfsscan.io/ipfs/" + cid + "?filename=" + filepath.Base(filePath)
				cids = append(cids, cid)
				log.Println(cid + "  <========  " + filePath)
			}
		}
	}

	return strings.Join(cids, ",")
}
