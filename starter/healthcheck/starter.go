package healthcheck

import "github.com/shanluzhineng/app"

func init() {
	app.RegisterStartupAction(healthcheckStartup)
}
