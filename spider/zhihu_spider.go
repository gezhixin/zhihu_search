package spider

import (
	"fmt"
	// "github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
	// "zhihu_search/utils"
)

const (
	Zgz_url  string = `https://www.zhihu.com/api/v4/members/elyn/followers?include=data[*].answer_count,articles_count,gender,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics&offset=0&limit=100`
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
					time.Sleep(time.Second * time.Duration(1))
				} else {
					timeCounter2++
					time.Sleep(time.Second * time.Duration(60))
				}

				if timeCounter1 > 60*30 {
					timeCounter1 = 0
					isFast = false
				}

				if timeCounter2 > 3 {
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
