package apricate

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/assistants"
	"AdityaHegde/apricate-ai/src/main/go/apricate/caravans"
	"AdityaHegde/apricate-ai/src/main/go/apricate/farms"
	"AdityaHegde/apricate-ai/src/main/go/apricate/locations"
	"AdityaHegde/apricate-ai/src/main/go/apricate/markets"
	"AdityaHegde/apricate-ai/src/main/go/apricate/plants"
	"AdityaHegde/apricate-ai/src/main/go/apricate/users"
	"gorm.io/gorm"
)

type ApricateService struct {
	UserService      users.UserService
	LocationsService locations.LocationService
	IslandService    locations.IslandService
	FarmService      farms.FarmService
	FarmPlotService  farms.FarmPlotService
	PlantService     plants.PlantService
	MarketService    markets.MarketService
	TradeService     markets.TradeService
	AssistantService assistants.AssistantService
	CaravanService   caravans.CaravanService
}

func NewApricateService(db *gorm.DB) (*ApricateService, error) {
	err := db.AutoMigrate(users.UserEntity{})

	err = db.AutoMigrate(locations.RegionEntity{})
	err = db.AutoMigrate(locations.IslandEntity{})
	err = db.AutoMigrate(locations.LocationEntity{})
	err = db.AutoMigrate(locations.PortEntity{})

	err = db.AutoMigrate(farms.FarmPlotEntity{})
	err = db.AutoMigrate(farms.FarmEntity{})

	err = db.AutoMigrate(plants.ConsumableOption{})
	err = db.AutoMigrate(plants.GrowthStageEntity{})
	err = db.AutoMigrate(plants.PlantEntity{})

	err = db.AutoMigrate(markets.MarketEntity{})
	err = db.AutoMigrate(markets.TradeEntity{})

	err = db.AutoMigrate(assistants.AssistantEntity{})

	err = db.AutoMigrate(caravans.CaravanEntity{})

	if err != nil {
		return nil, err
	}

	return &ApricateService{
		UserService:      users.NewUserService(db),
		LocationsService: locations.NewLocationService(db),
		IslandService:    locations.NewIslandService(db),
		FarmService:      farms.NewFarmService(db),
		FarmPlotService:  farms.NewFarmPlotService(db),
		PlantService:     plants.NewPlantService(db),
		MarketService:    markets.NewMarketService(db),
		TradeService:     markets.NewTradeService(db),
		AssistantService: assistants.NewAssistantService(db),
		CaravanService:   caravans.NewCaravanService(db),
	}, nil
}

func (service *ApricateService) LoadInitialData() error {
	_, err := service.UserService.GetOrFetch(
		service.UserService.FieldsService.GetPrimaryValue(nil))

	_, err = service.IslandService.LoadAllMapFromRemote("")
	_, err = service.LocationsService.LoadAllFromRemote("")

	_, err = service.FarmService.LoadAllFromRemote("")
	_, err = service.FarmPlotService.LoadAllMapFromRemote("")

	_, err = service.PlantService.LoadAllMapFromRemote("")

	marketEntities, err := service.MarketService.LoadAllFromRemote("")
	service.TradeService.LoadFromMarketEntity(marketEntities)

	_, err = service.AssistantService.LoadAllMapFromRemote("")
	_, err = service.CaravanService.LoadAllFromRemote("")
	return err
}
