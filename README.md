# config

## Introduction

'config' is a simple golang library for handling configration files.

The main feature provided is a Config type, which represents a set of configuration properties. Properties are name-spaced, and access via 'x.y' dot-notation.

## Import

Use:

    go get github.com/mrmorphic/config

Then in your program:

    import (
        "github.com/mrmorphic/config"
    )

## Usage

The following is a simple program fragment that shows the basic usage of the library:

    import (
        "github.com/mrmorphic/config"
    )

    func main() {
        // construct a config object from a json file
        conf, e := config.ReadFromFile("config.json")
        if e != nil {
           panic(e)
        }

        // Get a property from the config
        v := conf.AsString("app.myAppName")
    }

The JSON file should contain a single object, whose properties form the top-level of the namespace.

If a key doesn't exist, Get() returns nil.

The types of values returned are the same as for JSON parsing. In particular, numeric literals in the json file are returned as float64, even if they look like int literals.