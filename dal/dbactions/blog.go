package dbactions

import "github.com/antiphy/mememe/dal/models"

func CreateBlogArticle(a *models.Article) error {
	return db.Create(a).Error
}

func QueryBlogArticle(params *models.QueryParams) (*models.ExtendArticle, error) {
	var ea models.ExtendArticle
	return &ea, nil
}

func QueryBlogArticles(params *models.QueryParams) ([]models.ExtendArticle, error) {
	var eas []models.ExtendArticle
	return eas, nil
}

func UpdateBlogArticle(a *models.Article) error {
	return db.Model(a).Update(a).Error
}
