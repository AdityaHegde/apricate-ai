package markets

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type MarketService struct {
	commons.EntityService[MarketEntity]
}

func NewMarketService(db *gorm.DB) MarketService {
	marketService := MarketService{}
	marketService.Db = db
	marketService.FieldsService = &commons.SimpleEntityFieldsService[MarketEntity]{
		PrimaryKey: "location_symbol",
		PrimaryValueGetter: func(entity *MarketEntity) string {
			return entity.LocationSymbol
		},
	}
	marketService.Handler = new(commons.DummyEntityHandler[MarketEntity])
	marketService.HttpService = commons.HttpService[MarketEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/markets",
		},
	}
	return marketService
}
