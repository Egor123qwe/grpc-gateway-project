package mongoDb

import "github.com/Egor123qwe/grpc-gateway-project/internal/tools"

type ConnectionMongo struct {
	MongoHost       string
	MongoPort       string
	MongoDBName     string
	MongoCollection string
}

func NewConfig() ConnectionMongo {
	return ConnectionMongo{
		MongoHost:       tools.GetEnv("HOST_MONGO", "mongo"),
		MongoPort:       tools.GetEnv("PORT_MONGO", "27017"),
		MongoDBName:     tools.GetEnv("DBNAME_MONGO", "mongo"),
		MongoCollection: tools.GetEnv("COLLECTION_MONGO", "quotes"),
	}
}
