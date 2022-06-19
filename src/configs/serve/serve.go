package serve

import ("github.com/spf13/cobra"
		"log"
		"net/http"
		"carRent/src/routers"
		"os"
// "fmt"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Short: "start api server",
	RunE: serve,
}

func serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var address string = "127.0.0.1:1616"

		port := os.Getenv("PORT");
		
		if port != "" {
			address = ":" + port
		}

		log.Println("App running")

		err := http.ListenAndServe(address, mainRoute)
		
		if err != nil {
			return err
		}
		
		return nil
	} else {
		return err
	}
}