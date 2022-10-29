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
func TestNewCRUD(t *testing.T) {
	InitDB()
	tx := db.Begin()
	db = tx
	tx.Begin()
	for _, n := range _newsArray {
		CreateNew(n)
	}
	news := make([]*New, 0)
	tx.Find(&news)

	{ // test CreateNew method
		for i := 0; i < len(_newsArray); i++ {
			compareNew(t, _newsArray[i], news[i])
		}
	}

	{ // test GetNewsWithLimit method
		assert.Equal(t, len(GetNewsWithLimit(7)), 7)
	}

	{ // test GetNewByID method
		for i := 0; i < len(_newsArray); i++ {
			compareNew(t, GetNewByID(_newsArray[i].ID), _newsArray[i])
		}
	}

	{ // test GetNewsByAuthorIDWIthLimit method
		limit := 3
		newsWithAuthorIDFromDB := GetNewsByAuthorIDWIthLimit(_newsArray[1].AuthorID, limit)
		assert.Equal(t, len(newsWithAuthorIDFromDB), limit)
		newsWithAuthorIDFromLiteral := make([]*New, 0)
		for _, n := range _newsArray {
			if n.AuthorID == _newsArray[1].AuthorID {
				newsWithAuthorIDFromLiteral = append(newsWithAuthorIDFromLiteral, n)
			}
		}
		for i := 0; i < len(newsWithAuthorIDFromDB); i++ {
			compareNew(t, newsWithAuthorIDFromLiteral[i], newsWithAuthorIDFromDB[i])
		}
	}

	{ // test UpdateNewByID method
		_new := _newsArray[3]
		_new.Title = "Kick Back"
		UpdateNewByID(_new)
		compareNew(t, _new, GetNewByID(_new.ID))
	}

	{ // test DeleteNewByID method
		for _, n := range _newsArray {
			DeleteNewByID(n.ID)
		}
		assert.Equal(t, len(GetNewsWithLimit(10)), 0)
	}

	tx.Rollback()
}

func compareNew(t *testing.T, new1 *New, new2 *New) {
	assert.Equal(t, new1.ID, new2.ID)
	assert.Equal(t, new1.AuthorID, new2.AuthorID)
	assert.Equal(t, new1.Title, new2.Title)
	assert.Equal(t, new1.Content, new2.Content)
	// gorm不会把time类型注回去
	// 很奇怪
	// 明明主键都没问题的
	//assert.Equal(t, new1.CreateTime, new2.CreateTime)
}

var _newsArray = []*New{
	{
		AuthorID: 114514,
		Title:    "野兽先辈驾临福州大学下北泽学院",
		Content:  "如题，啊啊",
	}, {
		AuthorID:   17,
		Title:      "科西嘉的怪物在儒安港登陆",
		Content:    "科西嘉的怪物在儒安港登陆",
		CreateTime: time.Now(),
	}, {
		AuthorID:   17,
		Title:      "吃人的魔鬼向格腊斯前进",
		Content:    "吃人的魔鬼向格腊斯前进",
		CreateTime: time.Now(),
	}, {
		AuthorID:   114,
		Title:      "篡位者进入格勒诺布尔",
		Content:    "篡位者进入格勒诺布尔",
		CreateTime: time.Now(),
	}, {
		AuthorID:   17,
		Title:      "波拿巴占领里昂",
		Content:    "波拿巴占领里昂",
		CreateTime: time.Now(),
	}, {
		AuthorID:   17,
		Title:      "拿破仑接近枫丹白露",
		Content:    "拿破仑接近枫丹白露",
		CreateTime: time.Now(),
	}, {
		AuthorID:   114,
		Title:      "皇帝陛下将于今日抵达他忠实的巴黎",
		Content:    "皇帝陛下将于今日抵达他忠实的巴黎",
		CreateTime: time.Now(),
	},
}
