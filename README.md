# golang-hexagonal
## Hexagonal architecture

* Domain layer @ **internal** folder root
* Application layer: first level folders (except **platform**): creating
* Infrastructure layer @ **internal/platform** 

## Go folder structure

**kit** folder is what in other languages is known as **shared** folder

## "Production ready" goals

* Acceptable performance
* Support for concurrency
* Support for asynchrony
* Robust error handling
* Observability (metrics, log)