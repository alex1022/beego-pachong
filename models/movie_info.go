package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
	Remark               string
	Create_time          string
	Movie_url            string
}

func init() {
	orm.Debug = true //调试模式会打印出sql
	orm.RegisterDataBase("default", "mysql", "lankai:AD1oor00t@tcp(127.0.0.1:3306)/douban?")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

//电源信息插入数据库
func AddMovie(movie_info *MovieInfo) (int64, error) {
	movie_info.Id = 0
	id, err := db.Insert(movie_info)
	return id, err
}

//获得导演名称
func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}

	return result[0][1]
}

//获得电影名称
func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]
}

//获得主演名称
func GetMovieMainActs(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a href=.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	MainActs := ""
	for _, v := range result {
		MainActs += v[1] + " / "
	}
	return MainActs
}

//电影类型
func GetMovieType(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	MovieType := ""
	for _, v := range result {
		MovieType += v[1] + " / "
	}
	return MovieType
}

//电影语言
func GetMovieLanguage(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">语言:</span> (.*)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]
}

//电影上映日期
func GetMovieIinitialReleaseDate(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	MovieType := ""
	for _, v := range result {
		MovieType += v[1] + " / "
	}
	return MovieType
}

//电影时长
func GetMovieRuntime(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]
}

//制片国家
func GetMovieCountry(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span class="pl">制片国家/地区:</span> (.*)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]
}

//电影评分
func GetMovierCatingNum(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]

}

//
func GetMovieUrls(movieHtml string) []string {
	if movieHtml == "" {
		return nil
	}

	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/subject/.*?from=subject-page)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	var movieSets []string
	for _, v := range result {
		movieSets = append(movieSets, v[1])
	}
	return movieSets
}
