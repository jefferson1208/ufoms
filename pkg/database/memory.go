package database

type MemoryProvider struct {
	dbProvider DataBaseProvider
}

func NewMemoryProvider(config *Configuration) (IDatabase, error) {
	return &MemoryProvider{dbProvider: MEMORY_DB}, nil
}

func (r *MemoryProvider) GetProvider() string {
	return string(r.dbProvider)
}
