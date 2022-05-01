package farms

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type FarmPlotService struct {
	commons.EntityService[FarmPlotEntity]
}

func NewFarmPlotService(db *gorm.DB) FarmPlotService {
	farmPlotService := FarmPlotService{}
	farmPlotService.Db = db
	farmPlotService.FieldsService = &commons.SimpleEntityFieldsService[FarmPlotEntity]{
		PrimaryKey: "uuid",
		PrimaryValueGetter: func(entity *FarmPlotEntity) string {
			return entity.Uuid
		},
	}
	farmPlotService.Handler = new(commons.DummyEntityHandler[FarmPlotEntity])
	farmPlotService.HttpService = commons.HttpService[FarmPlotEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/plots",
		},
	}
	return farmPlotService
}
