package database

import "errors"

var (
	ErrorUnloadDBProvider = errors.New("unable to load a db provider")
)

var dbProviders = map[DataBaseProvider]func(config *Configuration) (IDatabase, error){
	MYSQL_DB:  NewMySqlProvider,
	MEMORY_DB: NewMemoryProvider,
}

func ConfigureDBProvider(config *Configuration) (IDatabase, error) {

	p := DataBaseProvider(config.DBProvider)

	callback, found := dbProviders[p]

	if !found {
		return nil, ErrorUnloadDBProvider
	}

	provider, err := callback(config)

	return provider, err
}
