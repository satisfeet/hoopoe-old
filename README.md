# Hoopoe

![Hoopoe](http://bit.ly/1jPUlUI)

**Hoopoe** is the REST API service for [satisfeet](https://satisfeet.me).

## Issues

### Minor

Minor issues which should be easy to solve with:

* code style differs accross files
* mongo.Store.Index does not check nested structs
* mongo.Store does not set auto set id
* store.Search is not able to read tag from slices
* (json) marshal logic on models is redundant = extract logic to package

### Major

* define mongo independent id value
* Cannot remove model without fetching before
* store.Store does not auto-index (internal hooks)

## License

Copyright © 2014 Bodo Kaiser <i@bodokaiser.io>
