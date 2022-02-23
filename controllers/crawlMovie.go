package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	"web/models"
)

type CrawlMovieController struct {
	beego.Controller
}

/*
	1，判断是否是电影，是的话就插入MySQL数据库
	2，判断是否是电影，是的话就插入Redis队列中
	3，将url加入到set中，表示这个url已经get到了
*/
func (c *CrawlMovieController) CrawlMovie() {
	var movieinfo models.MovieInfo

	//redis连接
	models.ConnectRedis("115.238.30.132:6379")

	//爬虫入口url
	//一度怀疑自己正则或者哪里代码写错了，其实https://movie.douban.com   这个主页做了防爬措施的应该
	SuperUrl := "https://movie.douban.com/subject/27038228/"
	models.PutinSuperQueue(SuperUrl)
	for {
		fmt.Println("单条循环开始 =======================================")
		length := models.GetQueueLength()
		fmt.Println("++++++++++++++,", length)
		if length == 0 {
			break //如果url队列为空，则退出当前循环
		}
		SuperUrl = models.PopfromQueue()

		//判断SuperUrl是否被访问过了，访问过的话就跳过这个
		if models.IsVisit(SuperUrl) {
			continue
		}
		rsp := httplib.Get(SuperUrl)
		sMoviHtml, err := rsp.String() //获取到页面的源代码
		if err != nil {
			panic(err)
		}
		movieinfo.Movie_name = models.GetMovieName(sMoviHtml)
		if movieinfo.Movie_name != "" { //这里简单判断下，通过电影名称是否为空，作为最后一步判断是否是电源url
			movieinfo.Movie_director = models.GetMovieDirector(sMoviHtml)
			movieinfo.Movie_country = models.GetMovieCountry(sMoviHtml)
			movieinfo.Movie_grade = models.GetMovierCatingNum(sMoviHtml)
			movieinfo.Movie_language = models.GetMovieLanguage(sMoviHtml)
			movieinfo.Movie_main_character = models.GetMovieMainActs(sMoviHtml)
			movieinfo.Movie_type = models.GetMovieType(sMoviHtml)
			movieinfo.Movie_on_time = models.GetMovieIinitialReleaseDate(sMoviHtml)
			movieinfo.Movie_span = models.GetMovieRuntime(sMoviHtml)
			movieinfo.Movie_url = SuperUrl
			models.AddMovie(&movieinfo)
		}
		urls := models.GetMovieUrls(sMoviHtml)
		for _, url := range urls {
			models.PutinSuperQueue(url)
			c.Ctx.WriteString("<h1>" + url + "</h1>")
		}
		models.AddtoSet(SuperUrl)
		time.Sleep(1 * time.Second)
		fmt.Println("单条循环结束=======================================")
	}
	c.Ctx.WriteString("end of abc=======================================")
}
