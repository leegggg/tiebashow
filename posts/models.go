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

// GetPostHeader ...
func GetPostHeader(pid string, flr bool) (PostHeader, error) {
	db := common.GetDB()
	var postHeader PostHeader
	err := db.Where(&PostHeader{Pid: &pid, Flr: flr}).First(&postHeader).Error
	return postHeader, err
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

// GetFirstPostHeaderByKz ...
func GetFirstPostHeaderByKz(kz int64) (PostHeader, error) {
	db := common.GetDB()
	var header PostHeader
	err := db.Where(&PostHeader{Parent: kz, Flr: false, Floor: 1}).First(&header).Error
	return header, err
}

// GetLastPostHeaderByKz ...
func GetLastPostHeaderByKz(kz int64) (PostHeader, error) {
	db := common.GetDB()
	var header PostHeader
	err := db.Where(&PostHeader{Parent: kz, Flr: false}).Order("floor desc").First(&header).Error
	return header, err
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
	Downloaded *time.Time `json:"downloaded"`
	Title      *string    `json:"title"`
	Source     *string    `json:"source"`
	Link       *string    `json:"link" gorm:"primary_key"`
	Path       *string    `json:"path"`
	Pid        *string    `json:"pid"`
	ModDate    *time.Time `json:"mod_date"`
	Status     int64      `json:"status"`
	Comment    *string    `json:"comment"`
}

// TableName ...
func (AttachementHeader) TableName() string {
	return "ATTACHEMENT_HEADER"
}

// PostAttachement ...
type PostAttachement struct {
	Parent  int64      `json:"parent" gorm:"primary_key"`
	Floor   int64      `json:"floor" gorm:"primary_key"`
	Link    *string    `json:"link" gorm:"primary_key"`
	AttID   *string    `json:"attId"`
	Status  int64      `json:"status"`
	Comment *string    `json:"comment"`
	ModDate *time.Time `json:"mod_date"`
}

// TableName ...
func (PostAttachement) TableName() string {
	return "POST_ATTACHEMENT"
}

// GetAttachementHeadersByPostHeader ...
func GetAttachementHeadersByPostHeader(post PostHeader) ([]AttachementHeader, error) {
	db := common.GetDB()
	var attachementHeaders []AttachementHeader
	err := db.Table("ATTACHEMENT_HEADER").
		Joins("JOIN POST_ATTACHEMENT on (POST_ATTACHEMENT.attid = ATTACHEMENT_HEADER.link) ").
		Where(&PostAttachement{Parent: post.Parent, Floor: post.Floor}).
		Find(&attachementHeaders).Error
	return attachementHeaders, err
}

// Post ...
type Post struct {
	PostHeader
	Content      Content  `json:"content"`
	NbFlr        int64    `json:"nb_flr"`
	Attachements []string `json:"attachements"`
}
