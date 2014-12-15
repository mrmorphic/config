# config

## Introduction

'config' is a simple golang library for handling configuration files.

The main feature provided is a Config type, which represents a set of configuration properties. Properties are name-spaced, and access via 'x.y' dot-notation.

Initialising the configuration from json files and/or environment variables is supported.

## Import

Use:

    go get github.com/mrmorphic/config

Then in your program:

    import (
        "github.com/mrmorphic/config"
    )

## Usage

The following is a simple program fragment that shows the basic usage of the library in loading configuration from a single config.json file:

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

Here is a more complex example, where there are multiple config files for different system components, and we also want to map some environment variables into the config as well (ok, this is a somewhat contrived example, but demonstrates possible usage):

    func main() {
        // construct a config object from a json file. Properties from this
        // are added to the top level of the namespace.
        conf, e := config.ReadFromFile("app_config.json")
        if e != nil {
           panic(e)
        }

        // merge in db config, putting these settings under "database.*".
        // The third parameter (override) says that this config file will
        // override any settings of the same name that are already there.
        conf.AddFile("db_config.json", "database", true)

        // merge in environment variables of the form APP_*, into the
        // "env" namespace, but don't override.
        conf.AddEnvironment("APP_", "env", false)

        // Get a property from the config
        v := conf.AsString("app.myAppName")
    }

The intent is to not prescribe how your application represents configuration, but to support two common configuration modes (at the same time), and present a simple, uniform way for your application to read that configuration.