# Hoopoe

![Hoopoe](http://bit.ly/1jPUlUI)

**Hoopoe** is the REST API service for [satisfeet](https://satisfeet.me).

## Issues

### Minor

Minor issues which should be easy to solve with:

* store.Search is not able to read tag from slices
* code style differs accross files

### Major

* store.Store does not auto-index (internal hooks)
* Cannot remove model without fetching before
* define mongo independent id value
* (json) marshal logic on models is redundant = extract logic to package

## License

Copyright © 2014 Bodo Kaiser <i@bodokaiser.io>
