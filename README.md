# report-transaction

Service in charge of generating and notifying user transaction reports.

## Content

1. [Introduction](#introduction)
   1. [Execution flow](#execution-flow)
   2. [Database schema](#database-schema)
2. [Quick start](#quick-start)
3. [Local testing](#local-testing)
   1. [Program](#program)
   2. [Database](#database)
4. [Program modes](#program-modes)
	1. [CLI](#cli)
	2. [Lambda](#lambda)
5. [Configuration](#configuration)
6. [Dependencies](#dependencies)
7. [Project structure](#project-structure)


## Introduction

The aim of this project is to create a flexible program able to run in any mode or deployed in any environment needed. 
Currently, there is support to execute it as a CLI application or as a [Lambda](https://aws.amazon.com/lambda/) with an 
[API Gateway](https://aws.amazon.com/api-gateway/), but potentially, it can be extended to support any other mode you 
want. <br>
The program objective is to process a CSV file stored in Amazon [S3](https://aws.amazon.com/s3/) that contains a list of 
credit and debit transactions of an account. It receives the filename and account ID and generates a report with 
information extracted from the file. Once generated, an email will be sent to the account owner using Amazon 
[SES](https://aws.amazon.com/ses/). <br>

An example summary email will contain:
- Total balance
- Average total debit
- Average total credit
- Average debit and credit grouped by month
- Total transactions grouped by month

### Execution flow
![Execution flow](./assets/execution_flow.png "Execution flow")
_This graph is not following any rules or convention, it's just a quick hint of the program execution flow. Created
using [Excalidraw](https://excalidraw.com/)._

### Database schema
![Database schema](./assets/database_schema.png "Database schema")
_Where an Account has a one-to-many relationship with Transaction. Created using [dbdiagram](https://dbdiagram.io/home)._

## Quick start

**1.** Clone the repository:

```shell
$ git clone git@github.com:ifarinho/report-transaction.git
```

**2.** Build the image
```shell
$ make build-image
```

**3.** Done!

## Local testing

For local testing the image you can run the program in CLI mode. Keep in mind that you will need valid AWS credentials
and a proper user that have permissions to both S3 and SES. For the database, you could just use a Postgres image.

### Program

**1.** Generate a local `.env` file and set the corresponding values. Check the [configuration](#configuration) section
for hints.

```shell
$ make dotenv
```

A sample CSV file is provided, but you can generate a new one with random values using Python. The file will be created 
at the root level of the project with the name `txns_<current-time-iso-format>.csv`, then upload this file to your S3 
Bucket and set the correct env file values:

```shell
$ make csv
```

**2.** Run the program locally. Example where `txns.csv` is the filename and `69420` is the account ID:

```shell
$ make local FILENAME="txns.csv" ACCOUNT="69420"
```

### Database

To generate a testing database just build a default Postgres [image](https://hub.docker.com/_/postgres):

**1.** Build image:

```shell
$ docker run --name dev-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
```

**2.** Source the database using the sample SQL files from the `storage/` directory. Don't forget to set a valid email 
to test SES correctly.

## Program modes

The program can run in two different modes: CLI and Lambda. The mode must be set via an environmental variable, but 
more on that in the [configuration](#configuration-and-environmental-variables) section.

### CLI

The program must be run with two arguments:
- `--filename`: string. The CSV filename inside the S3 Bucket. Don't include the Bucket name or the prefix, these values
are loaded via env.
- `--account`: uint. The account ID related to the file.

### Lambda

The program is also ready to be deployed with AWS Lambda. For this particular case, you will need to set 
`ENV_RUN_MODE` to `2` (Lambda), set the allowed cors origin with `ENV_CORS_ORIGIN` and `ENV_ALLOWED_METHODS` to work 
with the API Gateway. Program arguments are no longer needed and can be ignored. <br>
An example API Gateway request proxy event is:

```json
{
	"filename": "txns.csv",
	"account_id": 69420
}
```

## Configuration

Some environmental variables must be set to the program be able to run, also there some optional ones related to the
specific run mode. A `template.env` file with reference values is provided.

#### General

- `ENV_RUN_MODE`: Set the program run mode.
  - `1`: CLI mode
  - `2`: Lambda mode
- `ENV_SERVICE_EMAIL`: Sender email address, this must be verified by AWS SES and the user must be in the right group 
with permissions to use the service.
- `ENV_CORS_ORIGIN`: Cors origin rules for the API Gateway. This value is not needed for the program to run in CLI mode.
- `ENV_ALLOWED_METHODS`: Allowed origin methods for the API Gateway. Default only `POST`.
- `ENV_PROJECT_NAME`: The name of the project. Default is `report-transaction`.

#### Amazon

- `ENV_AWS_ACCESS_KEY_ID`: AWS access key ID.
- `ENV_AWS_ACCESS_SECRET_KEY`: AWS access secret key.
- `ENV_AWS_CREDENTIAL_TOKEN`: This variable is optional.
- `ENV_AWS_REGION`: Region.
- `ENV_AWS_S3_BUCKET`: Bucket name. Ex: `report-transaction`.
- `ENV_AWS_S3_PREFIX`: File prefix. Ex: `account`. The full path would be `report-transaction/account/<id>/<file>`.

#### Postgres

- `ENV_POSTGRES_HOST`: Database host.
- `ENV_POSTGRES_USER`: Database user.
- `ENV_POSTGRES_PASSWORD`: User password.
- `ENV_POSTGRES_NAME`: Database name.
- `ENV_POSTGRES_PORT`: Database exposed port.
- `ENV_POSTGRES_SSL_MODE`: Enable/disable SSL mode.
- `ENV_POSTGRES_TIME_ZONE`: Database timezone.
- `ENV_POSTGRES_DATA_SOURCE_NAME`: Connection string generated from all the previous values.
- `ENV_POSTGRES_BATCH_CREATE_SIZE`: Default create batch size.

## Dependencies

The program uses some dependencies to execute Postgres database operations, connect to Amazon services and handle 
transaction amounts more precisely using fixed points decimals.

- [GORM](https://gorm.io/): Full featured ORM to handle SQL databases.
- [Postgres driver](https://github.com/go-gorm/postgres): Driver used by GORM  to execute database operations.
- [Amazon SDK](https://github.com/aws/aws-sdk-go): Official SDK tools to consume AWS services.
- [shopspring/decimal](https://github.com/shopspring/decimal): Arbitrary-precision fixed-point decimal numbers with GORM
and SQL compatibility.

## Project structure

The project is not following any specific rule, just some common conventions from this 
[repository](https://github.com/golang-standards/project-layout) and experience from past developments. This structure
aims to prevent circular import errors and to be as simple as possible to organize packages by function.

- `/`: Project directories. Build, instructions and template files.
  - `cmd`:
    - `app`: Application `main.go` file.
  - `assets`: Miscellaneous support files.
  - `internal`:
    - `app`:
      - `args`: Program arguments definitions and parsing.
      - `awsdk`: AWS session initialization, clients and services methods.
      - `db`: Postgres database connection and repository methods.
      - `env`: Environmental variables load.
      - `file`: CSV file reading.
      - `notification`: Email creation and sending.
      - `tools`:
        - `calculate`: Calculation and numeric parsing operations.
        - `datetime`: Time related functions.
        - `decode`: Deserialization and decoding functions.
        - `stringify`: To string conversion.
      - `transaction`: Transaction manipulation and Report generation.
    - `pkg`:
      - `event`: Main program execution process and run mode. Other modes should be added here.
  - `scripts`: Collection of scripts used in the project.
  - `storage`: Example SQL files to source a testing database.
  - `web`:
    - `templates`: HTML email template.
