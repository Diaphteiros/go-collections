# Go Collections
A collection of Golang data structures, similar to the Java Collection Framework.

## Interfaces

### Collection[T]

A generic collection interface, similar to `java.util.Collection`.
Inherits from `Iterable`.

### List[T]

A list interface, similar to `java.util.List`.
Inherits from `Collection`.

### Queue[T]

A queue interface.
Inherits from `Collection`.

### Iterator[T]

An interface for iterators, similar to `java.util.Iterator`.

### Iterable[T]

An interface for iterables, similar to `java.util.Iterable`.

## Implementations

### ArrayList[T]

A list implementation that works on an underlying slice which stores the actual data.
Similar to `java.util.ArrayList`.
Implements `Collection`, `List`, and `Iterable`.

### LinkedList[T]

A list implementation that works on helper structs linked to each other by pointers.
Similar to `java.util.LinkedList`.
Implements `Collection`, `List`, `Queue`, and `Iterable`.

## Utilities

### Filters

The `filters` package provides utility functions to filter maps, slices, and collections with given filter functions.

### Maps

The `maps` package contains some utility functions for working with go maps (mainly a `Merge` function).
