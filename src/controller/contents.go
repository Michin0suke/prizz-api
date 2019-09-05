package controller

import (
	"database/sql"
	"net/http"
	"regexp"
	"time"

	// "encoding/base64"
	"fmt"
	"strconv"

	"github.com/Michin0suke/prizz-api/src/model"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func errorJson(c *gin.Context, s string, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": s})
	panic(err)
}

func addCORS(ctx *gin.Context) {
	// ctx.Header("Access-Control-Allow-Origin", "*")
	// ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// ctx.Header("Access-Control-Max-Age", "86400")
	// ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func generateJson(c *gin.Context, query string, isSearch bool) {
	limit := "10"
	r := regexp.MustCompile(`\d+,\d+`)

	if c.Query("limit") != "" {
		limit = c.Query("limit")

		if !r.MatchString(limit) {
			_, err := strconv.Atoi(limit)

			if err != nil {
				errorJson(c, "Parameter [ limit ] is invalid.", err.Error())
			}
		}
	}

	isRaw := false
	if c.Query("raw") == "true" {
		isRaw = true
	}

	db := model.DBConnect()
	result, err := db.Query(query + " LIMIT " + limit)

	if err != nil {
		errorJson(c, "An error occurred in the database connection.", err.Error())
	}

	// json返却用
	contents := []model.Content{}
	for result.Next() {
		content := model.Content{}

		var id uint
		var name string
		var winner uint
		var imageURL sql.NullString
		var createdAt time.Time
		var updatedAt mysql.NullTime
		var limitDate time.Time
		var link string
		var provider sql.NullString
		var way sql.NullString
		var categories []string
		var twitterWays []string
		var isOneclick bool
		var twitterRaw string

		err = result.Scan(
			&id,
			&name,
			&winner,
			&imageURL,
			&createdAt,
			&updatedAt,
			&limitDate,
			&link,
			&provider,
			&way,
			&isOneclick,
		)
		if err != nil {
			errorJson(c, "An error occurred while scanning from the database.", err.Error())
		}

		strID := strconv.Itoa(int(id))
		result, err := db.Query("SELECT category FROM categories WHERE id = " + strID)
		if err != nil {
			errorJson(c, "An error occurred in the SQL related to the category table.", err.Error())
		}
		for result.Next() {
			var category string
			err = result.Scan(&category)
			if err != nil {
				errorJson(c, "An error occurred when scanning category data from the database.", err.Error())
			}
			categories = append(categories, category)
		}

		result, err = db.Query("SELECT twitter_way FROM twitter_way WHERE id = " + strID)
		if err != nil {
			errorJson(c, err.Error(), err.Error())
		}
		for result.Next() {
			var twitterWay string
			err = result.Scan(&twitterWay)
			twitterWays = append(twitterWays, twitterWay)
		}

		if isRaw {
			result, err = db.Query("SELECT raw FROM twitter WHERE id = " + strID)
			if err != nil {
				errorJson(c, "An error occurred in the SQL related to the twitter_raw data.", err.Error())
			}
			for result.Next() {
				err = result.Scan(&twitterRaw)
				if err != nil {
					errorJson(c, "An error occurred when scanning twitter_raw data from the database.", err.Error())
				}
			}
		}

		content.ID = id
		content.Name = name
		content.Winner = winner
		if imageURL.Valid {
			content.ImageURL = imageURL.String
		} else {
			content.ImageURL = ""
		}
		content.CreatedAt = createdAt
		if updatedAt.Valid {
			content.UpdatedAt = updatedAt.Time
		} else {
			// content.UpdatedAt = time.Time{}
			content.UpdatedAt = time.Time{}
		}
		content.LimitDate = limitDate
		content.Link = link
		if provider.Valid {
			content.Provider = provider.String
		} else {
			content.Provider = ""
		}
		if way.Valid {
			content.Way = way.String
		} else {
			content.Way = ""
		}
		content.Category = categories
		content.IsOneclick = isOneclick
		content.TwitterWays = twitterWays
		content.TwitterRaw = twitterRaw

		contents = append(contents, content)
	}
	defer db.Close()

	if isSearch {
		c.JSON(http.StatusOK, contents[0])
	} else {
		c.JSON(http.StatusOK, contents)
	}
	// c.JSON(http.StatusOK, gin.H{"contents": contents})
}

func convCategory(s string) string {
	list := []string{
		"appliance",
		"baby",
		"books",
		"cash",
		"cosmetics",
		"daily",
		"fashion",
		"foods",
		"gift",
		"goods",
		"kitchen",
		"movie",
		"sports",
		"stationery",
		"ticket",
		"toy",
		"travel",
		"vehicle",
		"other",
	}

	if s == "" {
		return ""
	}

	for _, elem := range list {
		if s == elem {
			return " AND category = '" + elem + "'"
		}
	}
	return "error"
}

func convOrder(order string) string {
	switch order {
	case "new":
		return "created_at DESC"
	case "winner":
		return "winner DESC"
	default:
		return "limit_date ASC"
	}
}

func ContentsGET(c *gin.Context) {
	addCORS(c)
	order := convOrder(c.Query("order"))
	category := convCategory(c.Query("category"))

	if category == "error" {
		errorJson(c, "parameter [ category ] is invalid.", "parameter [ category ] is invalid.")
	}

	way := ""
	if c.Query("way") != "" {
		way = " AND way = '" + c.Query("way") + "'"
	}

	// if c.Query("way") != "" {
	// 	r := regexp.MustCompile(`%`)

	// 	if r.MatchString(c.Query("way")) {
	// 		wayQuery, err := base64.StdEncoding.DecodeString(c.Query("way"))

	// 		if err != nil {
	// 			errorJson(c, "An error occurred when decode from base64", err.Error())
	// 		} else {
	// 			way = " AND way = '" + string(wayQuery) + "'"
	// 		}

	// 	} else {
	// 		way = " AND way = '" + c.Query("way") + "'"
	// 	}
	// }

	query := `
	SELECT DISTINCT c1.id, name, winner, image_url, created_at, updated_at, limit_date, link, provider, way, is_oneclick 
	FROM contents c1 
	LEFT OUTER JOIN categories c2 ON c1.id = c2.id 
	WHERE limit_date > CURRENT_TIMESTAMP ` + category + way + `
	ORDER BY ` + order

	fmt.Println(query)

	generateJson(c, query, false)
}

// タスク一覧
// func DeadlineGET(c *gin.Context) {
// 	generateJson(c, "SELECT * FROM contents WHERE limit_date > CURRENT_TIMESTAMP ORDER BY limit_date ASC")
// }

// func NewGET(c *gin.Context) {
// 	generateJson(c, "SELECT * FROM contents WHERE limit_date > CURRENT_TIMESTAMP ORDER BY created_at DESC")
// }

// func WinnerGET(c *gin.Context) {
// 	generateJson(c, "SELECT * FROM contents WHERE limit_date > CURRENT_TIMESTAMP ORDER BY winner DESC")
// }

// func CategoryGET(c *gin.Context) {
// 	order := convOrder(c.Query("order"))

// 	category := c.Param("category")
// 	if !category_contain(category) {
// 		errorJson(c, "Invalid category", "Invalid category: "+category)
// 	}

// 	query := `
// 	SELECT DISTINCT c1.id, name, winner, image_url, created_at, updated_at, limit_date, link, provider, way, is_oneclick
// 	FROM contents c1
// 	LEFT OUTER JOIN categories c2 ON c1.id = c2.id
// 	WHERE category = '` + category + `' AND limit_date > CURRENT_TIMESTAMP
// 	ORDER BY ` + order

// 	generateJson(c, query)
// }

func SearchGET(c *gin.Context) {
	addCORS(c)
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorJson(c, "Invalid id", err.Error())
	}
	generateJson(c, "SELECT * FROM contents WHERE id = "+c.Param("id"), true)
}

// func TwitterGET(c *gin.Context) {
// 	order := convOrder(c.Query("order"))
// 	generateJson(c, "SELECT * FROM contents WHERE way = 'Twitter' AND limit_date > CURRENT_TIMESTAMP ORDER BY "+order)
// }

func TotalNumberGET(c *gin.Context) {
	addCORS(c)
	where := ""
	var totalNumber uint
	if c.Query("way") == "twitter" {
		where = " AND way = 'twitter'"
	}
	db := model.DBConnect()
	result, err := db.Query("SELECT COUNT(*) FROM contents WHERE limit_date > CURRENT_TIMESTAMP" + where + ";")
	if err != nil {
		errorJson(c, "Invalid id", err.Error())
	}
	for result.Next() {
		err = result.Scan(&totalNumber)
		if err != nil {
			errorJson(c, "Invalid id", err.Error())
		}
	}
	defer db.Close()
	c.JSON(http.StatusOK, gin.H{"total_number": totalNumber})
}

/*
// タスク検索
func FindByID(id uint) model.Content {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM contents WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	content := model.Content{}
	for result.Next() {
		var createdAt, updatedAt time.Time
		var title string

		err = result.Scan(&id, &createdAt, &updatedAt, &title)
		if err != nil {
			panic(err.Error())
		}

		content.ID = id
		content.CreatedAt = createdAt
		content.UpdatedAt = updatedAt
		content.Name = title
	}
	return content
}

// タスク登録
func ContentsPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	_, err := db.Exec("INSERT INTO contents (title, created_at, updated_at) VALUES(?, ?, ?)", title, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. title: %s", title)
}

// タスク更新
func ContentsPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	now := time.Now()

	_, err := db.Exec("UPDATE contents SET title = ?, updated_at=? WHERE id = ?", title, now, id)
	if err != nil {
		panic(err.Error())
	}

	content := FindByID(uint(id))

	fmt.Println(content)
	c.JSON(http.StatusOK, gin.H{"content": content})
}

// タスク削除
func ContentsDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.Query("DELETE FROM contents WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}
*/
