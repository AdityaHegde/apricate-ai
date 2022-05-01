package assistants

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type AssistantService struct {
	commons.EntityService[AssistantEntity]
}

func NewAssistantService(db *gorm.DB) AssistantService {
	assistantService := AssistantService{}
	assistantService.Db = db
	assistantService.FieldsService = &commons.SimpleEntityFieldsService[AssistantEntity]{
		PrimaryKey: "uuid",
		PrimaryValueGetter: func(entity *AssistantEntity) string {
			return entity.Uuid
		},
	}
	assistantService.Handler = new(commons.DummyEntityHandler[AssistantEntity])
	assistantService.HttpService = commons.HttpService[AssistantEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/assistants",
		},
	}
	return assistantService
}
