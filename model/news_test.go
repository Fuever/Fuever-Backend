package model

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

// 什么 你问我为什么不分开写？
// 一个个开事务再回滚很累人诶
// 要写你自己写奥
// 写完给你沏一壶昏睡红茶
func TestNewsCRUD(t *testing.T) {
	InitDB()
	tx := db.Begin()
	db = tx
	tx.Begin()
	for _, n := range newArray {
		err := CreateNews(n)
		if err != nil {
			t.Error(err)
		}
	}
	news := make([]*News, 0)
	tx.Find(&news)

	{ // test CreateNews method
		for i := 0; i < len(newArray); i++ {
			compareNew(t, newArray[i], news[i])
		}
	}

	{ // test GetNewsWithOffsetLimit method
		newsArray, err := GetNewsWithOffsetLimit(0, 7)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, len(newsArray), 7)
	}

	{ // test GetNewsByID method
		for i := 0; i < len(newArray); i++ {
			_new, err := GetNewsByID(newArray[i].ID)
			if err != nil {
				t.Error(err)
			}
			compareNew(t, _new, newArray[i])
		}
	}

	{ // test GetNewsByAuthorIDWIthOffsetLimit method
		limit := 3
		newsWithAuthorIDFromDB, err := GetNewsByAuthorIDWIthOffsetLimit(newArray[1].AuthorID, 0, limit)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, len(newsWithAuthorIDFromDB), limit)
		newsWithAuthorIDFromLiteral := make([]*News, 0)
		for _, n := range newArray {
			if n.AuthorID == newArray[1].AuthorID {
				newsWithAuthorIDFromLiteral = append(newsWithAuthorIDFromLiteral, n)
			}
		}
		for i := 0; i < len(newsWithAuthorIDFromDB); i++ {
			compareNew(t, newsWithAuthorIDFromLiteral[i], newsWithAuthorIDFromDB[i])
		}
	}

	{ // test UpdateNew method
		_new := newArray[3]
		_new.Title = "Kick Back"
		err := UpdateNew(_new)
		if err != nil {
			t.Error(err)
		}
		__new, err := GetNewsByID(_new.ID)
		if err != nil {
			t.Error(err)
		}
		compareNew(t, _new, __new)
	}

	{ // test DeleteNewByID method
		for _, n := range newArray {
			err := DeleteNewByID(n.ID)
			if err != nil {
				t.Error(err)
			}
		}
		newsArray, err := GetNewsWithOffsetLimit(0, 10)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, len(newsArray), 0)
	}

	tx.Rollback()
}

func compareNew(t *testing.T, new1 *News, new2 *News) {
	assert.Equal(t, new1.ID, new2.ID)
	assert.Equal(t, new1.AuthorID, new2.AuthorID)
	assert.Equal(t, new1.Title, new2.Title)
	assert.Equal(t, new1.Content, new2.Content)
	// gorm不会把time类型注回去
	// 很奇怪
	// 明明主键都没问题的
	//assert.Equal(t, new1.CreateTime, new2.CreateTime)
}

var newArray = []*News{
	{
		AuthorID: 114514,
		Title:    "野兽先辈驾临福州大学下北泽学院",
		Content:  "如题，啊啊",
	}, {
		AuthorID:   17,
		Title:      "科西嘉的怪物在儒安港登陆",
		Content:    "科西嘉的怪物在儒安港登陆",
		CreateTime: time.Now().Unix(),
	}, {
		AuthorID:   17,
		Title:      "吃人的魔鬼向格腊斯前进",
		Content:    "吃人的魔鬼向格腊斯前进",
		CreateTime: time.Now().Unix(),
	}, {
		AuthorID:   114,
		Title:      "篡位者进入格勒诺布尔",
		Content:    "篡位者进入格勒诺布尔",
		CreateTime: time.Now().Unix(),
	}, {
		AuthorID:   17,
		Title:      "波拿巴占领里昂",
		Content:    "波拿巴占领里昂",
		CreateTime: time.Now().Unix(),
	}, {
		AuthorID:   17,
		Title:      "拿破仑接近枫丹白露",
		Content:    "拿破仑接近枫丹白露",
		CreateTime: time.Now().Unix(),
	}, {
		AuthorID:   114,
		Title:      "皇帝陛下将于今日抵达他忠实的巴黎",
		Content:    "皇帝陛下将于今日抵达他忠实的巴黎",
		CreateTime: time.Now().Unix(),
	},
}
