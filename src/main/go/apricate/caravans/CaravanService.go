package caravans

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"AdityaHegde/apricate-ai/src/main/go/apricate/warehouses"
	"gorm.io/gorm"
	"strconv"
)

type CaravanService struct {
	commons.EntityService[CaravanEntity]
}

func NewCaravanService(db *gorm.DB) CaravanService {
	assistantService := CaravanService{}
	assistantService.Db = db
	assistantService.FieldsService = &commons.SimpleEntityFieldsService[CaravanEntity]{
		PrimaryKey: "uuid",
		PrimaryValueGetter: func(entity *CaravanEntity) string {
			return entity.Uuid
		},
	}
	assistantService.Handler = new(commons.DummyEntityHandler[CaravanEntity])
	assistantService.HttpService = commons.HttpService[CaravanEntity]{
		UrlService: &commons.SimpleHttpURLService{
			ApiPath: "/my/caravans",
		},
	}
	return assistantService
}

type CharterCaravanBody struct {
	Origin      string           `json:"origin"`
	Destination string           `json:"destination"`
	Assistants  []int            `json:"assistants"`
	Wares       warehouses.Wares `json:"wares"`
}

func (service *CaravanService) CharterCaravan(
	origin string, destination string,
	assistants []int, wares warehouses.Wares,
) (*CaravanEntity, error) {
	var charterCaravanResp commons.ApricateResponse[CaravanEntity]
	body := CharterCaravanBody{
		Origin:      origin,
		Destination: destination,
		Assistants:  assistants,
		Wares:       wares,
	}
	_, err := commons.HttpClientPatch(
		&charterCaravanResp,
		service.HttpService.UrlService.GetManyUrl(""),
		body,
	)
	if err != nil {
		return nil, err
	}
	return &charterCaravanResp.Data, nil
}

type UnpackCaravanResp struct {
	AssistantsReleased []int `json:"assistants_released"`
}

func (service *CaravanService) UnpackCaravan(caravanId int64) error {
	var unpackCaravanResp commons.ApricateResponse[UnpackCaravanResp]
	_, err := commons.HttpClientDelete(
		&unpackCaravanResp,
		service.HttpService.UrlService.GetSingleUrl(strconv.FormatInt(caravanId, 10)),
	)
	return err
}
