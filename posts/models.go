package posts

import (
	"strconv"
	"time"

	"github.com/leegggg/tiebashow/common"
)

// PostHeader ...
// Models should only be concerned with database schema, more strict checking should be put in validator.
// More detail you can find here: http://jinzhu.me/gorm/models.html#model-definition
// HINT: If you want to split null and "", you should use *string instead of string.
// ThreadHeader ...
type PostHeader struct {
	Pid        *string    `json:"pid"`
	Parent     int64      `json:"parent" gorm:"primary_key"`
	Floor      int64      `json:"floor" gorm:"primary_key"`
	Flr        bool       `json:"flr" gorm:"primary_key"`
	Un         *string    `json:"un"`
	To         *string    `json:"to"`
	CreateDate *time.Time `json:"create_date"`
	Cid        *string    `json:"cid"`
	ModDate    *time.Time `json:"mod_date"`
}

// TableName ...
func (PostHeader) TableName() string {
	return "POST_HEADER"
}

// GetPostHeaders ...
func GetPostHeaders(pid string, pn int64, nb int64) ([]PostHeader, error) {
	db := common.GetDB()
	var postHeaders []PostHeader
	err := db.Where(&PostHeader{Pid: &pid}).
		Order("floor").Limit(nb).Offset(pn).Find(&postHeaders).Error
	return postHeaders, err
}

// GetPostHeadersByParentAndFlrFlag ...
func GetPostHeadersByParentAndFlrFlag(parent int64, flr bool, pn int64, nb int64) ([]PostHeader, error) {
	db := common.GetDB()
	var postHeaders []PostHeader
	err := db.Where(&PostHeader{Parent: parent, Flr: flr}).
		Order("floor").Limit(nb).Offset(pn).Find(&postHeaders).Error
	return postHeaders, err
}

// Content ...
type Content struct {
	Cid     *string    `json:"cid" gorm:"primary_key"`
	Content *string    `json:"content"`
	Source  *string    `json:"source"`
	ModDate *time.Time `json:"mod_date"`
}

// TableName ...
func (Content) TableName() string {
	return "CONTENT"
}

// GetContent ...
func GetContent(cid string) (Content, error) {
	db := common.GetDB()
	var content Content
	err := db.Where(&Content{Cid: &cid}).First(&content).Error
	return content, err
}

// GetNbPostFlr ...
func GetNbPostFlr(pid string) (int64, error) {
	db := common.GetDB()
	parent, err := strconv.ParseInt(pid, 10, 64)
	var count int64
	db.Model(&PostHeader{}).Where(&PostHeader{Parent: parent, Flr: true}).Count(&count)
	return count, err
}

// AttachementHeader ...
type AttachementHeader struct {
    Downloaded      *time.Time   `json:"downloaded"`
    Title           *string      `json:"title"`
    Source          *string      `json:"source"`
    Link            *string      `json:"link" gorm:"primary_key"`
    Path            *string      `json:"path"`
    Pid             *string      `json:"pid"`
    ModDate         *time.Time   `json:"mod_date"`
    Status          int64        `json:"status"`
    Comment         *string      `json:"comment"` 
}

// TableName ...
func (AttachementHeader) TableName() string {
	return "ATTACHEMENT_HEADER"
}

// GetAttachementHeadersByPid ...
func GetAttachementHeadersByPid(pid int64, flr bool, pn int64, nb int64) ([]AttachementHeader, error) {
	db := common.GetDB()
	var attachementHeaders []AttachementHeader
	parent := strconv.FormatInt(pid,10)
	err := db.Where(&AttachementHeader{Pid: &parent}).Find(&attachementHeaders).Error
	return attachementHeaders, err
}

// Post ...
type Post struct {
	PostHeader
	Content      Content
	NbFlr        int64
	Attachements []string
}
