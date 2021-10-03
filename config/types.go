package config

type Config struct {
	Port            string `json:"port"`
	MongoDbName     string `json:"mongo_db_name"`
	MongoURI        string `json:"mongo_uri"`
	KafkaBrokerAddr string `json:"kafka_broker_addr"`
	SqlmapPath      string `json:"sqlmap_path"`
}
