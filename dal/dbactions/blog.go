package dbactions

import "github.com/antiphy/mememe/dal/models"

func CreateBlogArticle(a *models.Article) error {
	return db.Create(a).Error
}

func QueryBlogArticle(params *models.QueryParams) (*models.ExtendArticle, error) {
	var ea models.ExtendArticle
	err := db.Raw(sqlQueryArticle, params.ID).Scan(&ea).Error
	return &ea, err
}

func QueryBlogArticles(params *models.QueryParams) ([]models.ExtendArticle, error) {
	var eas []models.ExtendArticle
	err := db.Raw(sqlQueryArticles, params.PageSize, (params.Page-1)*params.PageSize).Scan(eas).Error
	return eas, err
}

func UpdateBlogArticle(a *models.Article) error {
	return db.Model(a).Update(a).Error
}

const (
	sqlQueryArticle  = ``
	sqlQueryArticles = ``
)
