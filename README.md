# go-ent-experiment

Personal opinion about `ent.`

## Pros

* Solid code generation
* Handles migrations, both automated and versioned
* Excellent documentation
* GraphQL integration (have not tried)
* Can hook into the library with custom code

## Cons

* Not a fan of how you need to query data. Requires learning a whole ORM in order to work with a database, which has limited transferable knowledge to other projects
* The long chaining of functions is not something often seen in Golang, and feels out of place. `go fmt` doesn't force these on a newline which requires manual formatting or hard-to-read code
* Seems focussed on graph-like data structures, which might not be what you want (at least the naming of things like edges)


## Development

```bash
# start dev environment
$ docker compose up -d

# run tests
$ docker compose exec dev ./run-test.sh

# Generate entities
$ docker compose exec dev bash -c 'go generate ./ent'
```
