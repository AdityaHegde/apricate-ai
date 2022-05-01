package markets

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"AdityaHegde/apricate-ai/src/main/go/apricate/warehouses"
	"gorm.io/gorm"
)

type TradeService struct {
	commons.EntityService[TradeEntity]
}

func (service *TradeService) LoadFromMarketEntity(marketEntities *[]MarketEntity) {
	for _, marketEntity := range *marketEntities {
		service.loadFromMarketItems(marketEntity, marketEntity.Imports, false)
		service.loadFromMarketItems(marketEntity, marketEntity.Exports, true)
	}
}
func (service *TradeService) loadFromMarketItems(
	marketEntity MarketEntity, wares warehouses.Wares, exported bool,
) {
	service.loadFromItemsMap(marketEntity, wares.Produce, "PRODUCE", exported)
	service.loadFromItemsMap(marketEntity, wares.Goods, "GOODS", exported)
	service.loadFromItemsMap(marketEntity, wares.Seeds, "SEEDS", exported)
	service.loadFromItemsMap(marketEntity, wares.Tools, "TOOLS", exported)
}
func (service *TradeService) loadFromItemsMap(
	marketEntity MarketEntity, items map[string]int,
	itemCategory string, exported bool,
) {
	for itemName, itemValue := range items {
		tradeEntity := TradeEntity{
			LocationSymbol: marketEntity.LocationSymbol,
			ItemName:       itemName,
			ItemCategory:   itemCategory,
			Price:          itemValue,
			Exported:       exported,
		}
		service.CreateOrSave(&tradeEntity)
	}
}

func NewTradeService(db *gorm.DB) TradeService {
	tradeService := TradeService{}
	tradeService.Db = db
	tradeService.FieldsService = &commons.SimpleEntityFieldsService[TradeEntity]{
		PrimaryKey: "location_symbol",
		PrimaryValueGetter: func(entity *TradeEntity) string {
			return entity.LocationSymbol
		},
	}
	tradeService.Handler = new(commons.DummyEntityHandler[TradeEntity])
	return tradeService
}
