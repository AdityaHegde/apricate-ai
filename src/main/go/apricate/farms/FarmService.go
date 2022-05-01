package farms

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type FarmService struct {
	commons.EntityService[FarmEntity]
}

func NewFarmService(db *gorm.DB) FarmService {
	farmService := FarmService{}
	farmService.Db = db
	farmService.FieldsService = &commons.SimpleEntityFieldsService[FarmEntity]{
		PrimaryKey: "uuid",
		PrimaryValueGetter: func(entity *FarmEntity) string {
			return entity.Uuid
		},
	}
	farmService.Handler = new(commons.DummyEntityHandler[FarmEntity])
	farmService.HttpService = commons.HttpService[FarmEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/farms",
		},
	}
	return farmService
}
