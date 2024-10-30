package seed

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/todennus/oauth2-client-service/usecase/dto"
	"github.com/todennus/oauth2-client-service/wiring"
	"github.com/todennus/shared/middleware"
	"github.com/xybor-x/snowflake"
)

var userID int64
var clientName string

var Command = &cobra.Command{
	Use:   "seed",
	Short: "Seed the first client",
	Run: func(cmd *cobra.Command, args []string) {
		envPaths, err := cmd.Flags().GetStringArray("env")
		if err != nil {
			panic(err)
		}

		system, err := wiring.InitializeSystem(envPaths...)
		if err != nil {
			panic(err)
		}

		ctx := middleware.WithBasicContext(context.Background(), system.Config)

		resp, err := system.Usecases.OAuth2ClientUsecase.CreateFirst(ctx, &dto.OAuth2ClientCreateFirstRequest{
			UserID:     snowflake.ID(userID),
			ClientName: clientName,
		})
		if err != nil {
			fmt.Println("Failed:", err)
			return
		}

		fmt.Println("Seed the client successfully")
		fmt.Println("ClientID:", resp.Client.ClientID)
		fmt.Println("ClientSecret:", resp.ClientSecret)
	},
}

func init() {
	Command.Flags().Int64VarP(&userID, "userid", "u", 0, "the admin user id")
	Command.Flags().StringVarP(&clientName, "clientname", "c", "Admin Client", "the client name")
	Command.MarkFlagRequired("userid")
}
