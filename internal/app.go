package internal

import (
	"sync"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

// App stitches together all the parts to make a fully formed application.
type App struct {
	api.HTTPClient
}

var (
	//nolint:gochecknoglobals //it's private cann't be accessed outside this package
	instance *App
	//nolint:gochecknoglobals //it's private cann't be accessed outside this package
	once sync.Once
)

// Get obtains the initialized instance of the application.
func Get() *App { return instance }

// Instance instantiates a new Application from a config, if it's already
// been instantiated, an error is returned.
func Instance(httpClient api.HTTPClient) *App {
	once.Do(func() {
		instance = &App{
			HTTPClient: httpClient,
		}
	})

	return instance
}
