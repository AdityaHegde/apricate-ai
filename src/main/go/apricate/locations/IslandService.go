package locations

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"AdityaHegde/apricate-ai/src/main/go/utils"
	"gorm.io/gorm"
)

type IslandService struct {
	commons.EntityService[IslandEntity]
}

func NewIslandService(db *gorm.DB) IslandService {
	islandService := IslandService{}
	islandService.Db = db
	islandService.FieldsService = &commons.SimpleEntityFieldsService[IslandEntity]{
		PrimaryKey: "symbol",
		PrimaryValueGetter: func(entity *IslandEntity) string {
			return entity.Symbol
		},
	}
	islandService.Handler = &commons.SimpleEntityHandler[IslandEntity]{
		LoadedHandler: func(entity *IslandEntity) {
			for _, port := range entity.Ports {
				location := utils.ParseLocationSymbol(port.Symbol)
				entity.Symbol = location.Island
				db.Save(&port)
			}
		},
	}
	islandService.HttpService = commons.HttpService[IslandEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/islands",
		},
	}

	return islandService
}
