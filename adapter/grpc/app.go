package grpc

import (
	"github.com/todennus/oauth2-client-service/wiring"
	"github.com/todennus/proto/gen/service"
	"github.com/todennus/shared/config"
	"github.com/todennus/shared/interceptor"
	"google.golang.org/grpc"
)

func App(config *config.Config, usecases *wiring.Usecases) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.NewUnaryInterceptor().
				WithBasicContext().
				WithLogRoundTripTime().
				WithTimeout().
				WithAuthenticate().
				Interceptor(config),
		),
	)

	service.RegisterOAuth2ClientServer(s, NewOAuth2ClientServer(usecases.OAuth2ClientUsecase))

	return s
}
