package locations

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type LocationService struct {
	commons.EntityService[LocationEntity]
}

func (service *LocationService) GetNearbyLocations() (*map[string][]string, error) {
	var nearbyLocationsResp commons.ApricateResponse[map[string]map[string]string]
	_, err := commons.HttpClientGet(&nearbyLocationsResp, "/my/nearby-locations")
	if err != nil {
		return nil, err
	}

	var resp = map[string][]string{}
	for islandSymbol, island := range nearbyLocationsResp.Data {
		resp[islandSymbol] = []string{}
		for locationSymbol, _ := range island {
			var location LocationEntity
			existing := service.Db.First(
				location,
				service.FieldsService.GetPrimaryKey()+" = ?",
				locationSymbol,
			)
			if existing.RowsAffected == 0 {
				resp[islandSymbol] = append(resp[islandSymbol], locationSymbol)
			}
		}
	}

	return &resp, nil
}

func NewLocationService(db *gorm.DB) LocationService {
	locationService := LocationService{}
	locationService.Db = db
	locationService.FieldsService = &commons.SimpleEntityFieldsService[LocationEntity]{
		PrimaryKey: "symbol",
		PrimaryValueGetter: func(entity *LocationEntity) string {
			return entity.Symbol
		},
	}
	locationService.Handler = new(commons.DummyEntityHandler[LocationEntity])
	locationService.HttpService = commons.HttpService[LocationEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/locations",
		},
	}
	return locationService
}
