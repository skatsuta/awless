[![Build Status](https://api.travis-ci.org/wallix/triplestore.svg?branch=master)](https://travis-ci.org/wallix/triplestore)
[![Go Report Card](https://goreportcard.com/badge/github.com/wallix/triplestore)](https://goreportcard.com/report/github.com/wallix/triplestore)
[![GoDoc](https://godoc.org/github.com/wallix/triplestore?status.svg)](https://godoc.org/github.com/wallix/triplestore) 

# Triple Store

Triple Store is a library to manipulate RDF triples in a fast and fluent fashion.

RDF triples allow to represent any data and its relations to other data. It is a very versatile concept and is used in [Linked Data](https://en.wikipedia.org/wiki/Linked_data), graphs traversal and storage, etc....

Here the RDF triples implementation follows along the [W3C RDF concepts](https://www.w3.org/TR/rdf11-concepts/). (**Note that blank nodes and reification are not implemented**.). More digestible info on [RDF Wikipedia](https://en.wikipedia.org/wiki/Resource_Description_Framework)

## Features overview

- Create and manage triples through a convenient DSL
- Snapshot and query RDFGraphs
- Encode triples to binary, [DOT](https://en.wikipedia.org/wiki/DOT_(graph_description_language)), NTriples format
- Decode triples from binary
- CLI (Command line interface) utility to read and convert triples files.

Roadmap
- Simple RDF graph traversals API
- RDF graph comparison
- Encode to [Turtle syntax](https://en.wikipedia.org/wiki/Turtle_(syntax))

## Library 

This library is written using the [Golang](https://golang.org) language. You need to [install Golang](https://golang.org/doc/install) before using it.

Get it:

```sh
go get -u github.com/wallix/triplestore
```

Test it:

```
go test -v -cover -race github.com/wallix/triplestore
```

Import it in your source code:

```go
import (
	"github.com/wallix/triplestore"
	// tstore "github.com/wallix/triplestore" for less verbosity
)
```

Get the CLI with:

```
go get -u github.com/wallix/triplestore/cmd/triplestore
```

## Concepts

A triple is made of 3 components:

    subject -> predicate -> object

... or you can also view that as:

    entity -> attribute -> value

So

- A **triple** consists of a *subject*, a *predicate* and a *object*.
- A **subject** is a unicode string.
- A **predicate** is a unicode string.
- An **object** is a *resource* (or IRI) or a *literal* (blank node are not supported).
- A **literal** is a unicode string associated with a datatype (ex: string, integer, ...).
- A **resource**, a.k.a IRI, is a unicode string which point to another resource.

And

- A **source** is a persistent yet mutable source or container of triples.
- A **RDFGraph** is an **immutable set of triples**. It is a snapshot of a source and queryable .
- A **dataset** is a basically a collection of *RDFGraph*.

You can also view the library through the [godoc](https://godoc.org/github.com/wallix/triplestore)

## Usage

#### Create triples

Although you can build triples the way you want to model any data, they are usually built from known RDF vocabularies & namespace. Ex: [foaf](http://xmlns.com/foaf/spec/), ...

```go
triples = append(triples,
	SubjPred("me", "name").StringLiteral("jsmith"),
 	SubjPred("me", "age").IntegerLiteral(26),
 	SubjPred("me", "male").BooleanLiteral(true),
 	SubjPred("me", "born").DateTimeLiteral(time.Now()),
 	SubjPred("me", "mother").Resource("mum#121287"),
)
```

or dynamically and even shorter with

```go
triples = append(triples,
 	SubjPredLit("me", "age", "jsmith"), // String literal object
 	SubjPredLit("me", "age", 26), // Integer literal object
 	SubjPredLit("me", "male", true), // Boolean literal object
 	SubjPredLit("me", "born", time.now()) // Datetime literal object
 	SubjPredRes("me", "mother", "mum#121287"), // Resource object
)
```

#### Create triples from a struct

As a convenience you can create triples from a singular struct:

```go
type Address struct {
	Street string `predicate:"street"`
}

type Person struct {
	Name     string    `predicate:"name"`
	Age      int       `predicate:"age"`
	Size     int64     `predicate:"size"`
	Male     bool      `predicate:"male"`
	Birth    time.Time `predicate:"birth"`
	Surnames []string  `predicate:"surnames"`
	Addr     Address   `subject:"address"`
}

addr := &Address{...}
person := &Person{Addr: addr, ....}

tris := TriplesFromStruct("jsmith", person)

src := NewSource()
src.Add(tris)
snap := src.Snapshot()

snap.Contains(SubjPredLit("jsmith", "name", "..."))
snap.Contains(SubjPredLit("jsmith", "size", 186))
snap.Contains(SubjPredLit("jsmith", "surnames", "..."))
snap.Contains(SubjPredLit("jsmith", "surnames", "..."))
snap.Contains(SubjPredLit("address", "street", "..."))

```

#### Equality

```go
	me := SubjPred("me", "name").StringLiteral("jsmith")
 	you := SubjPred("me", "name").StringLiteral("fdupond")

 	if me.Equal(you) {
 	 	...
 	}
)
```

### Triple Source

A source is a persistent yet mutable source or container of triples

```go
src := tstore.NewSource()

src.Add(
	SubjPredLit("me", "age", "jsmith"),
	SubjPredLit("me", "born", time.now()),
)
src.Remove(SubjPredLit("me", "age", "jsmith"))
```

### RDFGraph

A RDFGraph is an immutable set of triples you can query. You get a RDFGraph by snapshotting a source:

```go
graph := src.Snapshot()

tris := graph.WithSubject("me")
for _, tri := range tris {
	...
}
```

### Codec

Triples can be encoded & decoded using either a simple binary format or more common clear format like NTriples, ...

Triples can therefore be persisted to disk, serialized or sent over the network.

For example

```go
enc := NewBinaryEncoder(myWriter)
err := enc.Encode(triples)
...

dec := NewBinaryDecoder(myReader)
triples, err := dec.Decode()
```

Create a file of triples under the NTriples format:

```go
f, err := os.Create("./triples.nt")
if err != nil {
	return err
}
defer f.Close()

enc := NewNTriplesEncoder(f)
err := enc.Encode(triples)

``` 

Encode to a DOT graph
```go
tris := []Triple{
        SubjPredRes("me", "rel", "you"),
        SubjPredRes("me", "rdf:type", "person"),
        SubjPredRes("you", "rel", "other"),
        SubjPredRes("you", "rdf:type", "child"),
        SubjPredRes("other", "any", "john"),
}

err := NewDotGraphEncoder(file, "rel").Encode(tris...)
...

// output
// digraph "rel" {
//  "me" -> "you";
//  "me" [label="me<person>"];
//  "you" -> "other";
//  "you" [label="you<child>"];
//}
```

Load a binary dataset (i.e. multiple RDFGraph) concurrently from given files:

```go
path := filepath.Join(fmt.Sprintf("*%s", fileExt))
files, _ := filepath.Glob(path)

var readers []io.Reader
for _, f := range files {
	reader, err := os.Open(f)
	if err != nil {
		return g, fmt.Errorf("loading '%s': %s", f, err)
	}
	readers = append(readers, reader)
}

dec := tstore.NewDatasetDecoder(tstore.NewBinaryDecoder, readers...)
tris, err := dec.Decode()
if err != nil {
	return err
}
...
```
