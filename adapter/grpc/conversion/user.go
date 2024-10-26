package conversion

import (
	ucdto "github.com/todennus/oauth2-client-service/usecase/dto"
	ucresource "github.com/todennus/oauth2-client-service/usecase/dto/resource"
	pbdto "github.com/todennus/proto/gen/service/dto"
	pbresource "github.com/todennus/proto/gen/service/dto/resource"
	"github.com/todennus/shared/enumdef"
	"github.com/xybor-x/snowflake"
)

func NewPbOAuth2Client(client *ucresource.OAuth2Client) *pbresource.OAuth2Client {
	return &pbresource.OAuth2Client{
		Id:           client.ClientID.Int64(),
		Name:         client.Name,
		OwnerId:      client.OwnerID.Int64(),
		AllowedScope: client.AllowedScope,
	}
}

func NewUsecaseOAuth2ValidateRequest(req *pbdto.OAuth2ClientValidateRequest) *ucdto.OAuth2ClientValidateRequest {
	return &ucdto.OAuth2ClientValidateRequest{
		ClientID:                snowflake.ID(req.ClientId),
		ClientSecret:            req.ClientSecret,
		ConfidentialRequirement: enumdef.ConfidentialRequirementTypeFromGRPC(req.Requirement),
		Scope:                   req.RequestedScope,
	}
}

func NewUsecaseOAuth2GetByIDRequest(req *pbdto.OAuth2ClientGetByIDRequest) *ucdto.OAuth2ClientGetByIDRequest {
	return &ucdto.OAuth2ClientGetByIDRequest{
		ClientID: snowflake.ID(req.ClientId),
	}
}

func NewPbOAuth2GetByIDResponse(resp *ucdto.OAuth2ClientGetByIDResponse) *pbdto.OAuth2ClientGetByIDResponse {
	if resp == nil {
		return nil
	}

	return &pbdto.OAuth2ClientGetByIDResponse{
		Client: NewPbOAuth2Client(resp.Client),
	}
}
