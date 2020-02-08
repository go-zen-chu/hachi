package cmd

import (
	"fmt"

	"github.com/go-zen-chu/hachi/pkg/hachi"
	"github.com/go-zen-chu/hachi/pkg/interface/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run hachi server",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		port := viper.Get("port").(int)
		h := handler.NewHandler()
		hs := hachi.NewHttpServer()
		hs.ConfigureRoute(h)
		if err := hs.Run(port); err != nil {
			fmt.Errorf("Failed running server: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().Int("port", 8080, "Port for http server")
	viper.GetViper().BindPFlag("port", serveCmd.Flags().Lookup("port"))
}
