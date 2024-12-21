package datasources

import "clean-code/datasources/database"

// DataSources is a struct that contains all the data sources
// It is used to pass different data sources to the server and services
type DataSources struct {
	DB database.Database
}
