package database

type MySqlProvider struct {
	dbProvider DataBaseProvider
}

func NewMySqlProvider(config *Configuration) (IDatabase, error) {
	return &MySqlProvider{dbProvider: MYSQL_DB}, nil
}

func (r *MySqlProvider) GetProvider() string {
	return string(r.dbProvider)
}
