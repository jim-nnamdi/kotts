package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jim-nnamdi/kotts/internal/models"
	"go.uber.org/zap"
)

var _ Client = &databaseHandler{}

type databaseHandler struct {
	logger *zap.Logger
}

func NewDatabaseHandler(logger *zap.Logger) *databaseHandler {
	return &databaseHandler{
		logger: logger,
	}
}

func (handler *databaseHandler) Databaseconn() (db *sql.DB) {
	// config, err := loadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	var (
		err error
	)
	db, err = sql.Open("mysql", "root:M@etroboomin50@tcp(localhost:3306)/kotts")
	if err != nil {
		handler.logger.Debug("could not connect to the database")
		return
	}
	return db
}

func (handler *databaseHandler) GetUserByUsername(username string) bool {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where username = ?", username)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return false
	}
	return true
}

func (handler *databaseHandler) GetUserByEmail(email string) (*models.User, error) {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ?", email)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return user_response, errors.New(err.Error())
	}
	return user_response, nil
}

func (handler *databaseHandler) GetByUsernameAndPassword(email string, password string) (*models.User, error) {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ? and password = ?", email, password)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return nil, err
	}
	return user_response, nil
}

func (handler *databaseHandler) GetUserHash(email string) []byte {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ?", email)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return nil
	}
	return []byte(user_response.Password)
}

func (handler *databaseHandler) GetAllArticles() (*[]models.Articles, error) {
	var (
		err            error
		articles_slice = make([]models.Articles, 0)
	)

	all_articles, err := handler.Databaseconn().Query("SELECT * FROM articles where approved = true")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("No data returned from the database")
			return nil, errors.New(err.Error())
		}
	}

	for all_articles.Next() {
		var article models.Articles
		err := all_articles.Scan(
			article.Id,
			article.Title,
			article.Description,
			article.Author,
			article.CreatedAt,
			article.UpdatedAt,
			article.Category,
			article.NoOfViews,
		)
		if err != nil {
			log.Print("could not scan articles details into the database")
			return nil, errors.New(err.Error())
		}

		// append value of article to articles slice
		articles_slice = append(articles_slice, article)
	}
	return &articles_slice, nil
}

func (handler *databaseHandler) GetByAuthor(author string) (*[]models.Articles, error) {
	var (
		err           error
		article_slice = make([]models.Articles, 0)
	)
	results, err := handler.Databaseconn().Query("SELECT * FROM articles where author=", author)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("could not retrieve authors articles")
			return nil, errors.New(err.Error())
		}
	}
	for results.Next() {
		var article models.Articles
		if err = results.Scan(
			article.Id,
			article.Title,
			article.Description,
			article.Author,
			article.CreatedAt,
			article.UpdatedAt,
			article.Category,
			article.NoOfViews,
		); err != nil {
			log.Print("could not scan data into database properly")
			return nil, errors.New(err.Error())
		}
		article_slice = append(article_slice, article)
		log.Print("article slice populated", article_slice)
	}
	return &article_slice, nil
}

func (handler *databaseHandler) GetSingleArticle(articleID int) (*models.Articles, error) {
	var (
		single_article = models.Articles{}
		err            error
	)
	result, err := handler.Databaseconn().Query("SELECT * FROM articles where id=?", articleID)
	if err != nil {
		log.Print("error fetching article")
		if err == sql.ErrNoRows {
			log.Print("No article found with required ID")
			return nil, errors.New(err.Error())
		}
	}
	for result.Next() {
		if err = result.Scan(
			single_article.Id,
			single_article.Title,
			single_article.Description,
			single_article.Author,
			single_article.CreatedAt,
			single_article.UpdatedAt,
			single_article.Category,
			single_article.NoOfViews,
		); err != nil {
			log.Print("error scanning data from database to view")
			return nil, errors.New(err.Error())
		}
	}
	return &single_article, nil
}
func (handler *databaseHandler) Close() error {
	return nil
}
