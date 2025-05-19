package pb

import "github.com/pocketbase/pocketbase"

func CreateApp() *pocketbase.PocketBase {
	app := pocketbase.New()
	return app
}
