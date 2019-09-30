# dynocsv

Exports DynamoDB table into CSV [![Build Status](https://travis-ci.org/zshamrock/vmx.svg?branch=master)](https://travis-ci.org/zshamrock/dynocsv) [![dynocsv](https://snapcraft.io/dynocsv/badge.svg)](https://snapcraft.io/dynocsv)                                                                                                                                                    

```
NAME:
   dynocsv - Export DynamoDB table into CSV file

USAGE:
   dynocsv
              --table/-t <table>
              [--columns/-c <comma separated columns>]
              [--limit/-l <number>]
              [--profile/-p <AWS profile>]
              [--output/-o <output file name>]

VERSION:
   1.0.0

AUTHOR:
   (c) Aliaksandr Kazlou

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --table value, -t value    table to export
   --columns value, -c value  optional columns to export from the table, if skipped, all columns will be exported
   --limit value, -l value    limit number of records returned, if not set (i.e. 0) all items are fetched (default: 0)
   --profile value, -p value  AWS profile to use to connect to DynamoDB, otherwise the value from AWS_PROFILE env var is used if available, or then "default" if it is not set or empty
   --output value, -o value   output file, or the default <table name>.csv will be used
   --help, -h                 show help
   --version, -v              print the version
```

Table of Contents
=================

* [Installation](#installation)
* [Usage](#usage)
* [AWS Connection](#aws-connection)
* [Limits](#limits)

## Installation                                                                                                                                              
                                                                                                                                                             
Use the `go` command:                                                                                                                                        
                                                                                                                                                             
    $ go get github.com/zshamrock/dynocsv
    
Or using `snap`:                                                                                                                                             
                                                                                                                                                             
    $ snap install dynocsv                                                                                                                                   
                                                                                                                                                             
[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/dynocsv)

## Usage                                                                                                                                                     
                                                                                                                                                             
    $ dynocsv -t <table name>
    
## AWS Connection

Connection to the AWS is established using profile credentials. There are 2 options to provide the AWS profile to use:

1. use explicit `--profile/-p` option to set the AWS profile to use, i.e. `dynocsv -p <profile name> -t <table name>`
2. set the env var `$AWS_PROFILE` before running the app, i.e. `AWS_PROFILE=<profile name> dynocsv -t <table name>`

If no explicit profile value is set, it looks for the env var `$AWS_PROFILE` if present or otherwise fallbacks to the `default` profile.

## Limits

Currently there are the following limitations:

- only `String`, `Boolean`, `Number`, `Map`, `StringSet` and `NumberSet` data types are supported to export the data from, attributes with other data type will still be present, but the value will be "" (empty string)
- there is no pause or proper throttling according to the current set table's RCU, so you might need manually to increase the RCU value temporarily for the period of running the export
    
## Copyright                                                                                                                                                 
                                                                                                                                                             
Copyright (C) 2019 by Aliaksandr Kazlou.                                                                                                                     
                                                                                                                                                             
dynocsv is released under MIT License.                                                                                                                       
See [LICENSE](https://github.com/zshamrock/dynocsv/blob/master/LICENSE) for details.      
