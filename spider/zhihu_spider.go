package spider

import (
	"fmt"
	// "github.com/PuerkitoBio/goquery"
	// "encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
	// "zhihu_search/utils"
)

const (
	Zgz_url  string = `https://www.zhihu.com/api/v4/members/akajiejie/followers?include=data[*].answer_count,articles_count,gender,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics&offset=0&limit=100`
	MongoUrl string = "127.0.0.1:9500"
)

func Start() {
	go func() {

		session, err := mgo.Dial(MongoUrl)
		if err != nil {
			return
		}

		db := session.DB("zhihu")

		defer session.Close()

		form := ZUserForm{}
		form.Paging.Next = Zgz_url
		form.Paging.IsEnd = false

		timeCounter1 := 0
		timeCounter2 := 0
		isFast := true

		for {
			GetJson(form.Paging.Next, &form)
			for i := 0; i < len(form.Users); i++ {

				user := form.Users[i]

				userInDB := ZUser{}
				err = db.C("User").Find(bson.M{"zid": user.ZId}).One(&userInDB)
				if userInDB.Uid.Valid() {
					continue
				}

				user.HomePageUrl = "https://www.zhihu.com/people/" + user.UrlToken
				user.Avart = strings.Replace(user.AvartTemplate, "{size}", "b", -1)
				getUserExtInfor(&user)
				fmt.Println(user)

				fmt.Println("insert")
				user.Uid = bson.NewObjectId()
				err = db.C("User").Insert(&user)
				if err != nil {
					fmt.Println(err)
				}

				if isFast {
					timeCounter1++
					time.Sleep(time.Second * time.Duration(2))
				} else {
					timeCounter2++
					time.Sleep(time.Second * time.Duration(60))
				}

				if timeCounter1 > 60*30 {
					timeCounter1 = 0
					isFast = false
				}

				if timeCounter2 > 10 {
					timeCounter2 = 0
					isFast = true
				}

			}

			if form.Paging.IsEnd {
				break
			}
		}

	}()

}

func getUserExtInfor(user *ZUser) {
	url := "https://www.zhihu.com/people/" + user.UrlToken + "/logs"
	doc, _ := GetHtmlDoc(url)

	boby := doc.Find("body").Text()
	fmt.Println(boby)

	user.Location, _ = doc.Find(".location").Attr("title")
	user.Business, _ = doc.Find(".business").Attr("title")
	user.Employment, _ = doc.Find(".employment").Attr("title")
	user.Position, _ = doc.Find(".position").Attr("title")
	user.Education, _ = doc.Find(".education").Attr("title")
	user.EducationExtra, _ = doc.Find(".education-extra").Attr("title")

	agreeCount := doc.Find(".zm-profile-header-user-agree").Find("strong").Text()
	user.AgreeCount, _ = strconv.Atoi(agreeCount)

	thxCount := doc.Find(".zm-profile-header-user-thanks").Find("strong").Text()
	user.ThxCount, _ = strconv.Atoi(thxCount)

	// doc.Find(".profile-navbar").Each(func(i int, s *goquery.Selection) {
	// 	s1 := s.Find(".item")
	// 	href, _ := s1.Attr("href")
	// 	fmt.Println(i, "  ", href)
	// 	if strings.Contains(href, "/asks") {
	// 		countStr := s.Find(".num").Text()
	// 		fmt.Println("提问 ： ", countStr, " href: ", href)
	// 		user.AsksCount, _ = strconv.Atoi(countStr)
	// 	}
	// })

}

type ZAnswer struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	Author  ZUser  `json:"author"`
	Content string `json:"content"`
}

func GetA() {

	type TForm struct {
		Paging  ZPaging                    `json:"paging"`
		Answers [](map[string]interface{}) `json:"data"`
	}

	session, err := mgo.Dial(MongoUrl)
	if err != nil {
		return
	}

	db := session.DB("zhihu")

	defer session.Close()

	url := `https://www.zhihu.com/api/v4/questions/24865212/answers?include=data[*].is_normal,admin_closed_comment,reward_info,is_collapsed,annotation_action,annotation_detail,collapse_reason,is_sticky,collapsed_by,suggest_edit,comment_count,can_comment,content,editable_content,voteup_count,reshipment_settings,comment_permission,created_time,updated_time,review_info,question,excerpt,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp,upvoted_followees;data[*].mark_infos[*].url;data[*].author.follower_count,badge[?(type=best_answerer)].topics&offset=20&limit=20&sort_by=created`
	form := TForm{}
	form.Paging.Next = url
	for {

		GetJson(form.Paging.Next, &form)

		for i := 0; i < len(form.Answers); i++ {
			ans := form.Answers[i]
			var content = ans["content"].(string)
			userMap := ans["author"].(map[string]interface{})
			gender := int(userMap["gender"].(float64))
			if strings.Contains(content, "深圳") {
				if gender == 0 {
					userHomePageUrl := "https://www.zhihu.com/people/" + ans["author"].(map[string]interface{})["url_token"].(string)
					fmt.Println(userHomePageUrl)
					aurl := ans["url"].(string)
					fmt.Println("url : ", aurl)

					user := ZUser{}
					user.Avart = userMap["avatar_url"].(string)
					user.UrlToken = userMap["url_token"].(string)
					user.ZId = userMap["id"].(string)
					user.ExTag = int64(ans["created_time"].(float64))
					user.Location = "深圳"

					userInDB := ZUser{}
					err = db.C("BUser").Find(bson.M{"zid": user.ZId}).One(&userInDB)
					if userInDB.Uid.Valid() {
						continue
					}

					user.Uid = bson.NewObjectId()
					db.C("BUser").Insert(&user)

				}

			}
		}

		if form.Paging.IsEnd {
			break
		}

		time.Sleep(time.Second * time.Duration(1))
	}

}
