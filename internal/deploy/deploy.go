package deploy

var watchtower *Watchtower

// Deployment function
// Install and prepare program environment for persistence
func Deploy() {

	// Setup watchtower
	watchtower = NewWatchtower()
	watchtower.BuildWatchtower()

}

func GetWatchtower() *Watchtower {
	return watchtower
}
