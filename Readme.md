


# Destinations Services
I implemented two different destination services for logging and saving messages.
### Destination 1 (without database, using file)
First destination have responsibility of counting messages and calculating total size of data received
we used a file for storing counters for recovering data in case of shutting down the service,
and we flush data to it every 1 minute not for every message to get best performance 
### Destination 2 (PostgreSQL,TimescaleDB)


# Architecture and Design Patterns

## Clean Architecture

## Singleton DPI Design Pattern
I used singleton design pattern to access instances
that have interaction with limited resources.
the files with name "dpi" in application layers are exactly for this purpose.

## Saga Pattern
Since we have two destination services I decided to have a simple saga pattern implementation
for interaction of broker and destinations services.



# Extra features
#### implemented unit test for all services
#### I implemented two destination with two different solutions
#### Saving average time to process (destination 2)


