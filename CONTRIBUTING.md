# Contributing

## Endpoints

Request methods should be written in a file named after the endpoint e.g. request methods to `/auth/*` belong in `auth.go`, and request methods to `/member/*` belong in `member.go`

Endpoint request method names should be of the form `($METHOD)($ENDPOINT)` where parameterized requests are omitted e.g request method for `GET /auth/token` becomes `GetAuthToken` and request method for `POST /member/{id}/report` becomes `PostMemberReport(id string)`

## Endpoint request method outputs

All request methods should return an unmarshalled struct with field equivalent to all data from the request response and an error if applicable. Response structs should be defined in [response_structs.go](letterboxd/response_structs.go), alphabetized.

## Opening a Pull Request

Relevant testing is required for pull request approval. CI is run using Github Actions on all commits. Upon approval an admistrator will update the library version according to [SemVer](https://semver.org/), update the `CHANGELOG.md`, and merge the PR to the `main` branch. Github Actions is configured to tag the commit with the released version update and create a [Github release](github.com/jtschelling/letterboxd-api-go-client/releases)
