package conversion

import (
	ucdto "github.com/todennus/oauth2-client-service/usecase/dto"
	ucresource "github.com/todennus/oauth2-client-service/usecase/dto/resource"
	pbdto "github.com/todennus/proto/gen/service/dto"
	pbresource "github.com/todennus/proto/gen/service/dto/resource"
	"github.com/todennus/shared/enumdef"
	"github.com/todennus/x/conversion"
	"github.com/xybor-x/snowflake"
)

func NewPbOAuth2Client(client *ucresource.OAuth2Client) *pbresource.OAuth2Client {
	return &pbresource.OAuth2Client{
		Id:             client.ClientID.Int64(),
		Name:           conversion.ConvertPointer(client.Name),
		OwnerId:        conversion.ConvertPointer(client.OwnerID).Int64(),
		IsAdmin:        conversion.ConvertPointer(client.IsAdmin),
		IsConfidential: conversion.ConvertPointer(client.IsConfidential),
	}
}

func NewUsecaseOAuth2ValidateRequest(req *pbdto.OAuth2ClientValidateRequest) *ucdto.OAuth2ClientValidateRequest {
	return &ucdto.OAuth2ClientValidateRequest{
		ClientID:                snowflake.ID(req.GetClientId()),
		ClientSecret:            req.GetClientSecret(),
		ConfidentialRequirement: enumdef.OAuth2ClientConfidentialRequirementTypeFromGRPC(req.GetRequirement()),
	}
}

func NewUsecaseOAuth2ValidateResponse(resp *ucdto.OAuth2ClientValidateResponse) *pbdto.OAuth2ClientValidateResponse {
	if resp == nil {
		return nil
	}

	return &pbdto.OAuth2ClientValidateResponse{
		Client: NewPbOAuth2Client(resp.OAuth2Client),
	}
}

func NewUsecaseOAuth2GetByIDRequest(req *pbdto.OAuth2ClientGetByIDRequest) *ucdto.OAuth2ClientGetByIDRequest {
	return &ucdto.OAuth2ClientGetByIDRequest{
		ClientID: snowflake.ID(req.GetClientId()),
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
