package healthcheck

import "github.com/abmpio/app"

func init() {
	app.RegisterStartupAction(healthcheckStartup)
}
