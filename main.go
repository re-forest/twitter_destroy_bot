package main

import (
	"errors"
	"time"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/k0kubun/pp"
)
const (
	API_key = "................................"
	API_key_secret = "................................"
	Access_token = "................................"
	Access_token_secret = "................................"
	ScreenName = "............................."
	day = 30
)
func main() {
	// 登入
	config := oauth1.NewConfig(API_key,API_key_secret)
	token := oauth1.NewToken(Access_token,Access_token_secret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// 使用者資訊
	// params := &twitter.UserShowParams{
	// 	ScreenName:      "betaghost1",
	// }
	// user, _, _ := client.Users.Show(params)
	// pp.Printf("user: %v", user)

	// 高級搜索
	// PremiumSearch,_,_ := client.PremiumSearch.SearchFullArchive(&twitter.PremiumSearchTweetParams{
	// 	Query:"参加者募集",
	// 	MaxResults:10,
	// },"Test")
	// pp.Println("FollowersCount", PremiumSearch)

	// 使用者貼文
	// params := &twitter.UserTimelineParams{
	// 	Count:           200,
	// 	ScreenName:      ScreenName,
	// }
	// tweets, _, _ := client.Timelines.UserTimeline(params)
	// pp.Printf("%v \n",tweets[0].Text)

	// 刪除貼文
	// 使用 UserTimeline 查詢貼文 ID 再使用 Destroy 刪除指定 ID 貼文
	// tweetID := int64(tweets[0].ID)
	// _,res,err:= client.Statuses.Destroy(tweetID,&twitter.StatusDestroyParams{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// pp.Printf("%v\n",res.StatusCode)

	// 刪除指定天數 ( day ) 以前po文
	lastTweetID := int64(0)
	var allTweets []twitter.Tweet
	now := time.Now()
	for{
		params := &twitter.UserTimelineParams{
			Count:           200,
			ScreenName:      ScreenName,
		}
		if lastTweetID != 0 {
			params.MaxID = lastTweetID - 1
		}
		tweets, _, err := client.Timelines.UserTimeline(params)
		if len(tweets) == 0 {
			break
		}
		if err != nil {
			return 
		}

		allTweets = append(allTweets, tweets...)
		
		for _, t := range tweets {
			lastTweetID = t.ID
		}
	}
	pp.Printf("All post count : %v \n",len(allTweets))
	
	for _, tweet := range allTweets {
		createdAt, err := tweet.CreatedAtTime()
			if err != nil {
				pp.Println(errors.New("errors tweet.CreatedAtTime()"))
			}
			pp.Printf("%v \n",createdAt)
			daysAgo := now.Sub(createdAt).Hours() / 24
			if int(daysAgo) >= day {
				_, _, err := client.Statuses.Destroy(tweet.ID, nil)
				if err != nil {
					pp.Println(errors.New("errors Statuses.Destroy")) 
				}
			}
		}
}