# Oaxaca
Restaurant website for Customers and Staff 
Team Project Module, Oaxaca

A group of 7 people did the Team Project, mostly worked on the Websockets, Database functionality and frontend side of the project.

Commit history has been wiped and forked. 

## Managment System

# Table of Contents

1. [Server](#server)

    1. [Requirements](#server-requirements)

    2. [Setup](#server-setup)

    3. [Swagger](#server-swagger)

    4. [Documentation](#server-docs)

2. [Website](#website)

    1. [Requirements](#website-requirements)

    2. [Setup](#website-setup)

### Server <a name="server"/>

#### Requirements <a name="server-requirements"/>

Please ensure that you have installed:

- Go (version 1.21 or higher)
- Swagger ([binary](https://github.com/swaggo/swag), optional)

#### Setup <a name="server-setup"/>

Create a file `.env` in the root of the `api` directory.

| Variable | Description |
| --- | --- |
| DB_ADDR | Address of the postgres database. |
| DB_USER | Username of the postgres database user. |
| DB_PASS | Password of the postgres database password. |
| DB_NAME | Name of the postgres database. |
| DB_PORT | Port that the postgres database is running on. |
| SQLITE_DB | Location of the sqlite database, if it is being used. |
| DB | Either sqlite or pg, to use the postgres database this should be pg, to use the sqlite database it should use sqlite. Note this will default to postgres if not variable exists. |
| ENVIRONMENT | should either be development or production, in development mode CORS is allowed while in production it is not. |

The server can be run by running

```sh
$ cd api
$ make
```

#### Swagger <a name="server-swagger"/>

When running the server you can access swagger docs to view the endpoints at *address/api/v1/docs*

The server runs on port 8080 by default.

#### Documentation <a name="server-docs"/>

This project uses golds documentation to generate both dynamic and statically served documentation.

To make run these commands in the `api` directory:

```sh
$ go install go101.org/golds@latest
$ make golds
```

You can then access the documentation with the `golds/index.html` file.

### Website <a name="website"/>

#### Requirements <a name="website-requirements"/>

Please ensure you have installed:

- npm & node.js

#### Setup <a name="website-setup"/>

Inside of the `web` directory please create a file called `.env.local`.

| Variable | Description |
| --- | --- |
| NEXT_PUBLIC_API_ADDR | Stores the address of the server, for example `localhost:8080` |

Running the website can be done by

```sh
$ cd web
$ npm install
$ npm run dev
```

Note npm i only has to be run once to install the npm packages.
