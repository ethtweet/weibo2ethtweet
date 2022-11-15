package model

import (
	"time"
)

// Weibo [...]
type Weibo struct {
	ID             string    `gorm:"primaryKey;column:id;type:varchar(20);not null" json:"id"`
	Bid            string    `gorm:"column:bid;type:varchar(12);not null" json:"bid"`
	UserID         string    `gorm:"column:user_id;type:varchar(20);default:null" json:"user_id"`
	ScreenName     string    `gorm:"column:screen_name;type:varchar(30);default:null" json:"screen_name"`
	Text           string    `gorm:"column:text;type:longtext;default:null" json:"text"`
	ArticleURL     string    `gorm:"column:article_url;type:varchar(100);default:null" json:"article_url"`
	Topics         string    `gorm:"column:topics;type:varchar(200);default:null" json:"topics"`
	AtUsers        string    `gorm:"column:at_users;type:varchar(1000);default:null" json:"at_users"`
	Pics           string    `gorm:"column:pics;type:varchar(3000);default:null" json:"pics"`
	VideoURL       string    `gorm:"column:video_url;type:varchar(1000);default:null" json:"video_url"`
	Location       string    `gorm:"column:location;type:varchar(100);default:null" json:"location"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;default:null" json:"created_at"`
	Source         string    `gorm:"column:source;type:varchar(30);default:null" json:"source"`
	AttitudesCount int       `gorm:"column:attitudes_count;type:int;default:null" json:"attitudes_count"`
	CommentsCount  int       `gorm:"column:comments_count;type:int;default:null" json:"comments_count"`
	RepostsCount   int       `gorm:"column:reposts_count;type:int;default:null" json:"reposts_count"`
	RetweetID      string    `gorm:"column:retweet_id;type:varchar(20);default:null" json:"retweet_id"`
	TweetID        string    `gorm:"column:tweet_id;type:varchar(200);default:''" json:"tweet_id"`
}

// TableName get sql table name.获取数据库表名
func (m *Weibo) TableName() string {
	return "weibo"
}
