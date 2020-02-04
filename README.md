# web-service-example
An example on how to implement a web service in Golang using the clean architecture.

The project is divided in:

- Domain
- Usecases
- Interfaces
- Infrastructure

## Domain
The `domain` describes all the high level entities and their repoisitory's interface.

## Usecases
The `usecases` package contains entities related strictly to the use case portion of the service (i.e. users), as well as a simple and clear implementation of what each use case should do.

## Interfaces
The `interfaces` are all those little parts that puts everything together:

- Web service interface
- All the repositories' interfaces

## Infrastructure
The `infrastructure` package defines all the low level logic (i.e. the database connection handler).



Finally to put all together we use quite the handful of dependency injection. The idea is that every single part is interchangeable, making both testing and service swithching extremely simple.

