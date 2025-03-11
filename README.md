## Requirements

To run this program you will need to install:
- PostresSQL
- Go CLI

For Dev:
- Goose (for migrations) `go install github.com/pressly/goose/v3/cmd/goose@latest`
  - from `sql/schema` directory of the project run `goose postgres "postgres://postgres:postgres@localhost:5432/gator" up` to do the migrations
- SQLC (for generating models) `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
  - from root of the project run `sqlc generate` to build models from SQL files

## Setup

set up a config in your home directory `~/.gatorconfig.json`

```json
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable",
  "current_user_name": "username_goes_here"
}
```


## Installation

Because GO does not allow custom names when installing application (as far as I know) you can either:
- install from source with `go install` and use `boot-dev-blog-aggregator` to run the program, or
- build the code with custom name `go build -o gator` and use generated executable instead

## Usage

first: make an account
> ./gator register YOUR_USERNAME

then you can follow some topics:
> ./gator follow https://blog.boot.dev/index.xml

let them update:
> ./gator agg 1m

and see the content:
> ./gator browse X
(where X is number of articles to be shown, they are sorted from latest or oldest by publishing data)
