package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

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

// article related SQLs

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

// insurance related SQLs

func (handler *databaseHandler) ApplyForMobilePhoneInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	result, err := handler.Databaseconn().Prepare("insert into mobileinsurance values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		handler.logger.Debug("could not add new insurance data for mobile", zap.String("error", err.Error()))
		return false, err
	}
	data, err := result.Exec(name, email, phonenumber, nameofphone, purchasedate, imeinumber, model, color, description, paid, createdAt, updatedAt)
	if err != nil {
		handler.logger.Debug("error populating database", zap.String("error", err.Error()))
		return false, err
	}
	check_success, err := data.LastInsertId()
	if err != nil || check_success == 0 {
		handler.logger.Debug("last insert id failed :error populating database", zap.String("error", err.Error()))
		return false, err
	}
	return true, nil
}

func (handler *databaseHandler) ApplyForLaptopInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	result, err := handler.Databaseconn().Prepare("insert into laptopinsurance values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		handler.logger.Debug("could not add new insurance data for laptop", zap.String("error", err.Error()))
		return false, err
	}
	data, err := result.Exec(name, email, phonenumber, nameofphone, purchasedate, imeinumber, model, color, description, paid, createdAt, updatedAt)
	if err != nil {
		handler.logger.Debug("error populating database", zap.String("error", err.Error()))
		return false, err
	}
	check_success, err := data.LastInsertId()
	if err != nil || check_success == 0 {
		handler.logger.Debug("last insert id failed :error populating database", zap.String("error", err.Error()))
		return false, err
	}
	return true, nil
}

func (handler *databaseHandler) AllMobilePhoneInsuranceApplications(email string) (*[]models.MobileInsurance, error) {
	var (
		mobile_insurance_model = &models.MobileInsurance{}
		mobile_insurance_slice = make([]models.MobileInsurance, 0)
		err                    error
	)
	result, err := handler.Databaseconn().Query("select * from mobileinsurance where `email`=?")
	if err != nil {
		handler.logger.Debug("could not select all mobile insurance plans", zap.String("error", err.Error()))
		return &mobile_insurance_slice, err
	}
	for result.Next() {
		err = result.Scan(
			&mobile_insurance_model.Id,
			&mobile_insurance_model.Name,
			&mobile_insurance_model.Email,
			&mobile_insurance_model.Phonenumber,
			&mobile_insurance_model.Nameofphone,
			&mobile_insurance_model.Purchasedate,
			&mobile_insurance_model.Imeinumber,
			&mobile_insurance_model.Model,
			&mobile_insurance_model.Color,
			&mobile_insurance_model.Description,
			&mobile_insurance_model.CreatedAt,
			&mobile_insurance_model.UpdatedAt,
		)
		mobile_insurance_slice = append(mobile_insurance_slice, *mobile_insurance_model)
	}
	if result.Err() != nil {
		return nil, err
	}
	return &mobile_insurance_slice, nil
}

func (handler *databaseHandler) AllLaptopInsuranceApplications(email string) (*[]models.LaptopInsurance, error) {
	var (
		laptop_insurance_model = &models.LaptopInsurance{}
		laptop_insurance_slice = make([]models.LaptopInsurance, 0)
		err                    error
	)
	result, err := handler.Databaseconn().Query("select * from laptopinsurance where `email`=?")
	if err != nil {
		handler.logger.Debug("could not select all laptop insurance plans", zap.String("error", err.Error()))
		return &laptop_insurance_slice, err
	}
	for result.Next() {
		err = result.Scan(
			&laptop_insurance_model.Id,
			&laptop_insurance_model.Name,
			&laptop_insurance_model.Email,
			&laptop_insurance_model.Phonenumber,
			&laptop_insurance_model.Nameofphone,
			&laptop_insurance_model.Purchasedate,
			&laptop_insurance_model.Imeinumber,
			&laptop_insurance_model.Model,
			&laptop_insurance_model.Color,
			&laptop_insurance_model.Description,
			&laptop_insurance_model.CreatedAt,
			&laptop_insurance_model.UpdatedAt,
		)
		laptop_insurance_slice = append(laptop_insurance_slice, *laptop_insurance_model)
	}
	if result.Err() != nil {
		return nil, err
	}
	return &laptop_insurance_slice, nil
}

func (handler *databaseHandler) Close() error {
	return nil
}
