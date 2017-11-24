package spider

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	ZUser struct {
		Uid            bson.ObjectId `json:"userId" bson:"_id"`
		IsFollowed     bool          `json:"is_followed"`
		AvartTemplate  string        `json:"avatar_url_template"`
		Type           string        `json:"user_type"`
		IsFollowing    bool          `json:"is_following"`
		UrlToken       string        `json:"url_token"`
		ZId            string        `json:"id"`
		Name           string        `json:"name"`
		Headline       string        `json:"headline"`
		Gender         int           `json:"gender"`
		Location       string        `json:"loction"`
		Business       string        `json:"business"`
		Employment     string        `json:"employment"`
		IsOrg          bool          `json:"is_org"`
		Position       string        `json:"position"`
		Education      string        `json:"education"`
		EducationExtra string        `json:"education_extra"`
		Avart          string        `json:"avart"`
		HomePageUrl    string        `json:"home_page_url"`
		FollowerCount  int           `json:"follower_count"`
		AnswerCount    int           `json:"answer_count"`
		AsksCount      int           `json:"asks_count"`
		ArticleCount   int           `json:"articles_count"`
		AgreeCount     int           `json:"agree_count"`
		ThxCount       int           `json:"thx_count"`
	}

	ZPaging struct {
		IsEnd    bool   `json:"is_end"`
		Totals   int    `json:"totals"`
		Previous string `json:"previous"`
		IsStart  bool   `json:"is_start"`
		Next     string `json:"next"`
	}

	ZUserForm struct {
		Paging ZPaging `json:"paging"`
		Users  []ZUser `json:"data"`
	}
)
