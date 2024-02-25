package config

/*
This package defines the configuration for the server.
The AppConfig is populated from `src/configs/app_config.json` file.
Each Config struct specifies the json tags to be used for unmarshalling.

Author: Satyam Shivam Sundaram
*/

import "fmt"

// MongoDBConfig holds configurations to connect to a MongoDB database
type MongoDBConfig struct {
	Url            string `json:"url"`
	SecretsPath    string `json:"secrets_path"`
	DatabaseName   string `json:"database_name"`
	CollectionName string `json:"collection_name"`
}

// HttpConfig holds configurations to run the server
type HttpConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

func (hc *HttpConfig) Address() string {
	return fmt.Sprintf("%s:%d", hc.Hostname, hc.Port)
}

// AppConfig holds all the configurations and is used by the server
// It's populated in LoadAppConfig and LoadAppConfigFrom functions.
type AppConfig struct {
	DocDBConfig  MongoDBConfig `json:"doc_db_config"`
	AuthDBConfig MongoDBConfig `json:"auth_db_config"`
	UserDBConfig MongoDBConfig `json:"user_db_config"`
	ServerConfig HttpConfig    `json:"server_config"`
}

// Sample returns a sample json string for the AppConfig (this is the default)
func (ac *AppConfig) Sample() string {
	return `{
			"server_config": {
				"hostname": "",
				"port": 8099
			},
			"doc_db_config": {
				"url": "mongodb://localhost:27017/admin",
				"database_name": "google_note_db",
				"collection_name": "note_collection",
				"prev_secrets_path": "configs/mongo_credentials.json",
				"secrets_path": "/tmp/mongo_creds.json"
			},
			"auth_db_config": {
				"url": "mongodb://localhost:27017/admin",
				"database_name": "auth_db_v0",
				"collection_name": "auth",
				"prev_secrets_path": "configs/mongo_credentials.json",
				"secrets_path": "/tmp/mongo_creds.json"
			},
			"user_db_config": {
				"url": "mongodb://localhost:27017/admin",
				"database_name": "user_db_v0",
				"collection_name": "user_collection",
				"prev_secrets_path": "configs/mongo_credentials.json",
				"secrets_path": "/tmp/mongo_creds.json"
			}
		}`
}
