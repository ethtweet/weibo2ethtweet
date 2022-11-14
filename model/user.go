package model

// User [...]
type User struct {
	ID               string `gorm:"primaryKey;column:id;type:varchar(20);not null" json:"id"`
	ScreenName       string `gorm:"column:screen_name;type:varchar(30);default:null" json:"screen_name"`
	Gender           string `gorm:"column:gender;type:varchar(10);default:null" json:"gender"`
	StatusesCount    int    `gorm:"column:statuses_count;type:int;default:null" json:"statuses_count"`
	FollowersCount   int    `gorm:"column:followers_count;type:int;default:null" json:"followers_count"`
	FollowCount      int    `gorm:"column:follow_count;type:int;default:null" json:"follow_count"`
	RegistrationTime string `gorm:"column:registration_time;type:varchar(20);default:null" json:"registration_time"`
	Sunshine         string `gorm:"column:sunshine;type:varchar(20);default:null" json:"sunshine"`
	Birthday         string `gorm:"column:birthday;type:varchar(40);default:null" json:"birthday"`
	Location         string `gorm:"column:location;type:varchar(200);default:null" json:"location"`
	Education        string `gorm:"column:education;type:varchar(200);default:null" json:"education"`
	Company          string `gorm:"column:company;type:varchar(200);default:null" json:"company"`
	Description      string `gorm:"column:description;type:varchar(400);default:null" json:"description"`
	ProfileURL       string `gorm:"column:profile_url;type:varchar(200);default:null" json:"profile_url"`
	ProfileImageURL  string `gorm:"column:profile_image_url;type:varchar(200);default:null" json:"profile_image_url"`
	AvatarHd         string `gorm:"column:avatar_hd;type:varchar(200);default:null" json:"avatar_hd"`
	Urank            int    `gorm:"column:urank;type:int;default:null" json:"urank"`
	Mbrank           int    `gorm:"column:mbrank;type:int;default:null" json:"mbrank"`
	Verified         bool   `gorm:"column:verified;type:tinyint(1);default:null;default:0" json:"verified"`
	VerifiedType     int    `gorm:"column:verified_type;type:int;default:null" json:"verified_type"`
	VerifiedReason   string `gorm:"column:verified_reason;type:varchar(140);default:null" json:"verified_reason"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}
