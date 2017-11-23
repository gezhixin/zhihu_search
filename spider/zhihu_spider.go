package spider

import (
	"fmt"
)

const (
	Zgz_url string = `https://www.zhihu.com/api/v4/members/zhang-jia-wei/followers?include=data%5B*%5D.answer_count%2Carticles_count%2Cgender%2Cfollower_count%2Cis_followed%2Cis_following%2Cbadge%5B%3F(type%3Dbest_answerer)%5D.topics&offset=20&limit=20`
)

func Start() {
	form := ZUserForm{}
	GetJson(Zgz_url, &form)
	for i := 0; i < len(form.Users); i++ {
		user := form.Users[i]
		fmt.Println(user)
	}
}
