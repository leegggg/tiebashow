package threads

import (
	"time"

	"github.com/leegggg/tiebashow/common"
	"github.com/leegggg/tiebashow/posts"
)

// ThreadHeader ...
// Models should only be concerned with database schema, more strict checking should be put in validator.
// More detail you can find here: http://jinzhu.me/gorm/models.html#model-definition
// HINT: If you want to split null and "", you should use *string instead of string.
// ThreadHeader ...
type ThreadHeader struct {
	Kz             int64      `json:"kz" gorm:"primary_key"`
	Kw             *string    `json:"kw"`
	Title          string     `json:"title"`
	Good           bool       `json:"good"`
	Top            bool       `json:"top"`
	Click          int64      `json:"click"`
	Reply          int64      `json:"reply"`
	AuthorName     *string    `json:"author_name"`
	AuthorNickname *string    `json:"author_nickname"`
	AuthorPortrait *string    `json:"author_portrait"`
	FirstPostID    int64      `json:"first_post_id"`
	Bakan          bool       `json:"bakan"`
	Vid            *string    `json:"vid"`
	Protal         bool       `json:"protal"`
	Membertop      bool       `json:"membertop"`
	MultiForum     bool       `json:"multi_forum"`
	FrsTpoint      *string    `json:"frs_tpoint"`
	Comment        *string    `json:"comment"`
	LastDate       *time.Time `json:"last_date"`
	ModDate        *time.Time `json:"mod_date"`
}

// TableName ...
func (ThreadHeader) TableName() string {
	return "THREAD_HEADER"
}

// GetThreadHeaders ...
func GetThreadHeaders(kw string, pn int64, nb int64) ([]ThreadHeader, error) {
	db := common.GetDB()
	var threadHeaders []ThreadHeader
	err := db.Where(&ThreadHeader{Kw: &kw}).
		Order("top desc, kz desc").Limit(nb).Offset(pn).Find(&threadHeaders).Error
	return threadHeaders, err
}

// GetGoodThreadHeaders ...
func GetGoodThreadHeaders(kw string, pn int64, nb int64) ([]ThreadHeader, error) {
	db := common.GetDB()
	var threadHeaders []ThreadHeader
	res := db.Where(&ThreadHeader{Kw: &kw}).
		Limit(nb).Where(&ThreadHeader{Good: true}).
		Order("top desc, kz desc").Offset(pn).Find(&threadHeaders)
	err := res.Error
	return threadHeaders, err
}

// SearchThreadHeaders ...
func SearchThreadHeaders(kw string, pn int64, nb int64, query string) ([]ThreadHeader, error) {
	db := common.GetDB()
	var threadHeaders []ThreadHeader
	var err error
	if "" == kw {
		err = db.Limit(nb).Where("title LIKE ?", "%"+query+"%").
			Order("kz desc").Offset(pn).Find(&threadHeaders).Error
	} else {
		err = db.Where("title LIKE ?", "%"+query+"%").Where(&ThreadHeader{Kw: &kw}).
			Order("kz desc").Limit(nb).Offset(pn).Find(&threadHeaders).Error
	}
	return threadHeaders, err
}

// Thread ...
type Thread struct {
	ThreadHeader
	First posts.Post `json:"first"`
	Last  posts.Post `json:"last"`
}
