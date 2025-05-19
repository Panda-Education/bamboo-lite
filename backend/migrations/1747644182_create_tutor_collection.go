package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

var tutorCollectionName = "tutors"

func init() {
	m.Register(func(app core.App) error {
		// Create a tutor collection
		collection := core.NewAuthCollection(tutorCollectionName)

		// Set default rules
		collection.ListRule = types.Pointer("id = @request.auth.id")
		collection.ViewRule = types.Pointer("id = @request.auth.id")

		collection.Fields.Add(
			&core.TextField{
				Name:     "first_name",
				Required: true,
			},
			&core.TextField{
				Name:     "last_name",
				Required: true,
			},
		)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId(tutorCollectionName)
		if err != nil {
			return err
		}
		return app.Delete(collection)
	})
}
