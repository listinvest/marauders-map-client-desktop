package internal

var watchtower *Watchtower

// Deployment function
// Install and prepare program environment for persistence
func Deploy() *Watchtower {
	// Setup watchtower
	watchtower = NewWatchtower()
	watchtower.BuildWatchtower()
	watchtower.Daemonize()

	return watchtower
}

func GetWatchtower() *Watchtower {
	return watchtower
}

func GetWatchTower() *Watchtower {
	return watchtower
}
