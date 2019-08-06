# KintoHub Go Template

This is the standard format of how the Kinto Goons Engineering team will collaborating in making awesome go applications.

All Pull Requests should keep the following rules and standardizations in mind and **DO NOT ACCEPT** PRs that do not meet this requirements.

We need to build a team that is focused on quality code and take the time to learn and grow each other.

If you have any ideas on new standardizations, you can branch and submit a PR to this repository with your proposal. Once the proposal is accepted, you can code and submit the changes for review+merge.

Otherwise, no funky rules when it comes to coding go @ KintoHub!


## Requirements

### Mac

* Install Go version 1.12 or higher
* Install Goland and get a license from joseph@kintohub.com
* Install Git

When creating a project run `./setup_git` upon creation

* Setup your IDE to run Auto Tests by right clicking base folder in Goland -> Run -> go test {proj name}. Left bar has auto circular icon for rerunning tests on save
* Ensure when you commit in Goland that you *DO NOT* uncheck run git hooks

### Linux (TODO)
### Windows (TODO)

## Creating a new project

1) Click on `use this template` of this repostiory
2) `git clone {repo}` in your workspace directory of choice
3) `./setup_git` in the root of the project
4) Rename the `/cmd/common-go-example` folder to `/cmd/{repo-name}`
5) Delete and or modify the desired `/internal/controller` and `/internal/router` and `/internal/config`files.

NOTE: Would be nice if there was a script that cleaned up this project to do the above :).

## Code Style & Folder Structure

### Folder & file Overview 
Rule of thumb is ensure the scope of data is packaged in the scope of where it's used:

1) If string, number or value is used only in a function, at beginning of function `const name = "value"`
2) If string, number or value is used within a file, at beginning of file `const name = "value"`
3) If string, number or value is used across multiple files, place in `internal/package/constants/constants.go`
4) If string, number or value is used accross multiple packages, place in `internal/constants/constants.go`
5) If string, number or value is configurable, place in `internal/config/config.go` as an env var.

### Standard Folders & File 
* `.githooks` has required git checks for commit and pushing
* `cmd/{repo-name}` has the main file which requires no custom changes.
* `internal` has all global packages for the project
* `internal/config` has all env vars for the **entire** project in a single place
* `internal/controller` has all business logic for API endpoints for the app
* `internal/controller/queries` has all graphql/database queries for the project & tests for them
* `internal/controller/api.go` has all API Request and Response structs
* `internal/controller/types.go` has common controller structs
* `internal/controller/controller.go` handles all fasthttp request response Marshal/Unmarshal. After unmarshalling,
sends request struct to a functon to process.
* `internal/controller/{route-name}.go` has functionality for all *methods* IE: GET,POST for a specific endpoint
* `internal/router/route.go` has mapping to all Controller APIs to the fasthttp router.

### Style & Settings

We decided to build middleware for error handling.  When an error occurs, we panic the code to be caught
and written to the response. There are several utilities that automatically do this for us that start with
the function name of Panic such as `middelware.PanicClientErrorWithMessage`.

You should never do a raw panic yourself, always use the `middleware.Panic....` functions including internal
errors.

The reason why we are doing this is to avoid returning error objects in everything and writing if err != nil then...
100 times in a single project :). Simply call a function and it will not continue to execute if something went wrong.
This is similar to exceptions in Java/DotNet and general error middleware in NodeJs.
(inspired by [Koa middleware](https://github.com/koajs/koa/blob/master/docs/error-handling.md))

* We currently use tabs for indentation (following the [official guide](https://golang.org/doc/effective_go.html#formatting))
* All go error.New("messages should be 100% lowercase")
* go fmt on pre-commit and general style guide.
* We use int not int64,32,etc unless special case required
* Goland go PROXY is setup as `direct`
* TODO: More details one day here :)

**Always run `go vet ./...` to prevent common styling mistakes!**

### Naming Convention
In order to prepare for the future code linting and static code analysis, please follow the convention. 
* Variable name: Use `installationID`/`someURL` instead of `InstallationId`/`someUrl` ([official guide](https://github.com/golang/go/wiki/CodeReviewComments#initialisms))

Potential tools: 
* [golint](https://github.com/golang/lint)
  * Usage: `golint ./...`
* [go vet](https://golang.org/cmd/vet/)
  * Usage: `go vet ./...`

#### Interface names
As go has no concrete and scalable standard on interface names we decided to got with `I<InterfaceName>` standard

### Local Debugging and Dev

We are working on CLI tools to make it easier to debug in the cloud. But when you want to make changes
to commons-go, use go 1.12 feature of `replace {package} => {local/path/to/your/package}` for local testing
before pushing changes to commons-go.

Example (go.mod)
```
module github.com/kintohub/githubapp-go

go 1.12

require (
	github.com/kintohub/common-go v0.0.1-dev-a3f12bb
)

replace github.com/kintohub/common-go v0.0.1-dev-a3f12bb => ~/local/common-go
```

### Json Format

Hasura docs and standard is following PostgresDb with column names in underscore format `created_at`.

With json request/response we will stick to the standard `createdAt` lowerCamelCase

This means we need to ensure there is a struct for Request/Response of APIs in json format and a Hasura
response struct. We will translate and/or manipulate Hasura's response struct => our API Response.

A simple example has been written in `create_account.go` and `api.go`

## Error Handling

Error handling is completely handled by the common-go project. When KINTO_DEBUG_LEVEL=DEBUG is on,
all error objects will be passed to the client to be seen on the browser for easy debugging.
 
There are three types of errors:

### Internal Errors

Internal errors should occur when something happened that is not possible to happen in production.

* A function is not implemented
* Some sort of data came back that was unexpected or correct
* An uncaught exception

All uncaught exceptions are automatically logged and written as 500s to the client

But if you call Hasura's DB and you received data that is incorrect and has nothing to do with the actual
request that the client has sent, this error should be logged as `json.PanicValidateBytesToStruct` which
throws a 500 error if validation and/or json data is corrupt. Without validation, you can use `json.PanicBytesToStruct`

To do simple errors such as not implemented, simply do `panic("message in all lowercase")`

### Client Errors

All client errors are handled by `json.PanicClientBytesToStruct` or `json.PanicValidateClientBytesToStruct`.

If you have a custom client error due to some sort of validation or edge case, use `middleware.PanicClientError`
You should also use `errors.Wrap` to add additional information
```go
func GetDataFromDB() error {
	return errors.New("connection failed")
}
 
func FetchUsers() error {
    err := GetDataFromDB()
    if err != nil {
        middleware.PanicClientError(
            fasthttp.StatusInternalServerError,
            errors.Wrap(err, "getting data from db"),
        )
	}
}
```

### Hasura Errors

Hasura errors occur when there's a bad query or unhandled error which will automatically be logged and thrown
as a 500 request.

Otherwise you will get a HasuraError object which are errors that should be manually handled and thrown
with the appropriate message. IE: If a ConstraintViolationUniqueness error occurs, you should transform
that into a middleware.PanicClientErrorWithMessage with some message of {"error"{"name":"name already in use!"}}

Currently we support:

* Uniqueness
* ForeignKey (The ID provided does not exist in another table)
* NotNull (When you passed in nul when it shouldn't be) - TODO: Coop think this should be a 500 :).

## Logging

Logging uses a custom logger to notify us about the line of code & file which tihngs are happening.
This detailed logger should only be enabled when during `KINTO_DEBUG_LEVEL=debug`. By default our production
debug levels should be `KINTO_DEBUG_LEVEL=info` unless we are debugging something on live :).

**DO NOT EVER LOG REQUEST DATA** If you need to debug API requests, manually extract the information. We do not
want to accidentally log users passwords in clear text.

All messages for the logger should go as `logger.Debug` unless it is critical information about the service.
Critical information includes but is not limited to:

* Server has started
* Non sensitive (no passwords) config from env vars or files
* API Routes that are registered
* Jobs that are running every few seconds such as `/health` checks or caching
* A new user has been created with username X.

Non critical information includes:

* User X is searching with text Y
* Build X status has updated To Step 1/1000

Overall rule of thumb is if the log is going to be called > 10 times per user in a single minute, do not log it as 
`logger.Info` consider logging it as `logger.Debug`

## Docker file
For deployment use provided in the template dockerfile and change `./cmd/<your-project-name>/main.go` to 
corresponding to your project's name.  

## Main

Main function does the following logic in order:

0) Runs `_ .../autoload` for env file
1) All `init()` code is called automatically as per go's standard. see `config.go` for example
2) Initialize Logger
3) Create router
4) Create server & run

This file should not change for any project.

## Environment Variables

All env vars are put into `internal/config/config.go`. Every project **MUST** have a .env-example up to date

Copy the `.env-example` to be `.env` to have it run in Goland / local testing. This file is in `.gitignore`

KintoHub default env vars are as follows:

* KINTO_LOG_LEVEL=debug || info || warning || error || panic || trace
* SERVER_PORT=8081
* HASURA_HOST=localhost:8080/v1/graphql
* HASURA_ADMIN_SECRET= // keep empty if hasura is not set with admin secret 

Rule of thumb for environment variables is that they should be static information that is used in the project.

Env vars *SHOULD NOT* include logic, format properties. All formatting should be done in code for the final implementation
of the project.

## Router

Router creates a controller and routes maps all HTTP calls by their method + path to a specific function in `controller/controller.go`

## Controller

Controller is responsible for the following in order:

0) Obtain body + query data from request
1) Validate that data is clean by processing it into a struct from `api.go` using our common json utils
2) Pass the request struct to a api handler function in the controller package (see ping.go)
3) Receive the response
4) Write the response to the context.

**KEEP CONTROLLERS CLEAN WITHOUT EXTRA LOGIC.** controllers should contain only high level method calls 
and request processing flow *MUST* be obvious and easy readable. Abstract all the extra logic in to helper
methods per controller.  

### Headers

If there are any headers, you should pass them as a struct and/or map in the response of your handler func.
After getting them, you can write them manually using fasthttp

Standard headers (just as Content-Type) are written in the fasthttputility

### Validation of Struct

We are using ozzo-validator to validate. There's the root package then the /is package.

Root package has basic validation and the is package has `is.Email` for advanced valdations

See `api.go` for an example of writing validation code. The `json` package in `commons-go` has a function
called `json.JsonValidateClientBytesToStruct` which will validate and deserialize your byte code. If
there is any issue, it will throw an error to the client in the correct format for KintoHub to handle.

## Graphql Hasura & Queries

Each query has its own dedicated file due to how large queries can get. create_account and get_account would
have two separate files.

In a query file, will have:

* {QueryType}{ActionnName} - QueryType = Mutation || Query and ActioName = the same name as the file.
* {QueryType}{ActionName}Name - the name / action inside of the query
* {QueryType}{ActionName}Variables - a struct with the variables to match your query
* {QueryType}{ActionName}Response - a struct to handle the response of the query

A file called `models.go` inside of queries can contain common database models

## Documentation

@Edward to define

Official tool is provided, similar to APIDoc in JS. 

https://blog.golang.org/godoc-documenting-go-code

## Onboarding

If you have gotten to this point, you have successfully read the docs. To prove that you read all of this, create a PR
to this document by adding yourself to the Team section below of this readme. Submit the PR request to Ben.

## Team

* ${full-name} - ${email} | ${github-alias}
* Roman Z. - roman@kintohub.com | ronanamsterdam
* Joseph Cooper - joseph@kintohub.com | disturbing
* Edward Yu - edward@kintohub.com | edwardkcyu
* Laura Ambrose - laura@kintohub.com | lambro
* Nandi Wong - nandi@kintohub.com | nandiheath

## Misc Technologies

* Goland
* Ozzovalidation
* Fasthttp


### Misc Todo list

*[ ] Refactor HTTP Client / Requests to use channels versus ResponseHandler function

*[ ] Linter

*[ ] Static code analyzer

*[ ] Automatic documentation

*[ ] Json logger for logging metrics 