# Questions

## Where should we put creating valuer objects?

Transport (+) The domain is clean.
(+) There is no intermediate raw model to pass to the application, which then converts it into domain value objects.
(+) It seems logical that it is the responsibility of the transport to convert data into domain objects.
(-) If there are multiple transports, identical tests will have to be written to cover the scenarios when the value object constructors return an error.

Command/Query
(+) The domain is clean.
(+) Transports do not have identical tests to cover scenarios when the value object constructors return an error.
(-) If there are multiple commands/queries, identical tests will have to be written to cover scenarios when the value object constructors return an error.

Constructor with Raw Data in the Domain
(+) There are no conditions in the transport that need to be covered by tests for each transport. These conditions will be tested once in the domain.
(-) Conditions are found within constructors that accept raw data. The domain is not clean.
(-) Sometimes we use constructors to create objects in the domain itself, and in the domain, we want to work with value objects rather than raw data. We will have to create multiple constructors; we already have a special constructor for the database.
(-) If we create a raw constructor, properties of the same type, such as strings, can be listed. Using domain value objects allows for validation at the compiler level to ensure the correct types are passed. Raw constructors break this.