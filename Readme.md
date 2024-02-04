# Build Your Own Google Keep

## Description
The challenge was to use build your own Google Keep. The backend is a simple server that can 
store and retrieve notes. The frontend is a simple web app that can create, update, delete and 
view notes (currently very basic - even ugly).


## Components
`TODO (vamsaty): Update the components and their descriptions` 

There are 3 main components to this project -
1. Frontend web app - powers the UI (very basic/ugly currently)
2. Backend server - Provides RESTful APIs for - 
   1. Authenticating users (JWT token - fake secret key used for now)
   2. CRUD operations on Notes (only GetAllNotes and CreateNote integrated with the frontend 
      currently)
3. Persistent store - A document store for persisting Notes, it runs as a mongodb instance on a 
   docker container.

## Configurations and Ports
* Backend configurations defined in - `back-end/src/configs/app_config.json`
  * Go code to parse config : `back-end/src/internal/config/config_load.go`
  * Go code to store config : `back-end/src/internal/config/app_config.go`
* The database credentials are stored in - `/tmp/mongo_creds.json`
  * Can be configured by updating the `secrets_path` in `apps_config.json` above.
  * The `default_setup` step in `Makefile` creates the `/tmp/mongo_creds.json` using default values
* The frontend web app runs on port `3000` and view notes.

TODO(vamsaty):  Support arguments for custom configuration

Default backend configuration (as defined in `back-end/src/configs/app_config.json`) -
```
{
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
  }
}
```



## Usage
Steps to deploy the application locally (requires docker and go installed) -
```
make force_run_all
```


## Flags
TODO (vamsaty):  Update the flags used in the application
Currently, there are no flags used in the application.

## Issues
the `front-end` directory is a blob (in github) and needs to be cleaned up.
