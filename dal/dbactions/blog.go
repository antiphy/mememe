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

func CreateTableBlogArticle() error {
	return db.Exec("CREATE TABLE `mememe_article` (  `id` int(11) NOT NULL AUTO_INCREMENT,  `category_id` int(11) DEFAULT 0,  `title` varchar(255) DEFAULT '',  `content` text,  `status` tinyint(4) DEFAULT 0,  `created_by` int(11) DEFAULT 0,  `create_ts` bigint(20) DEFAULT 0,  `update_ts` bigint(20) DEFAULT 0,  PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;").Error
}

const (
	sqlQueryArticle  = `select b.id, b.category_id, b.title, b.content, b.status, b.created_by, a.name creator, b.create_ts, b.update_ts from mememe_article b inner join mememe_account a on b.created_by = a.id where b.id = ?`
	sqlQueryArticles = `select b.id, b.category_id, b.title, b.content, b.status, b.created_by, a.name creator, b.create_ts, b.update_ts from mememe_article b inner join mememe_account a on b.created_by = a.id limit ? offset ?`
)
