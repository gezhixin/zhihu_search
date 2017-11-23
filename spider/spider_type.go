package spider

type (
	ZUser struct {
		IsFollowed    bool   `json:"is_followed"`
		Avart         string `json:"avatar_url_template"`
		Type          string `json:"user_type"`
		AnswerCount   int    `json:"answer_count"`
		IsFollowing   bool   `json:"is_following"`
		UrlToken      string `json:"url_token"`
		Id            string `json:"id"`
		Gender        int    `json:"gender"`
		FollowerCount int    `json:"follower_count"`
		Location      string `json:"loction"`
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
