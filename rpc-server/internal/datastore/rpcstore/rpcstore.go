package rpcstore

/* Driven Adapter (Driven by the core)
This package implements the interface aka port of the Driven Adapter.

Its purpose is to allow the core to ineract with the database

By implementing it this way, we allow ourselves to be loosely coupled to any specific database types.

To change a database, we simple create anothe adapter implementing the same interface
*/
