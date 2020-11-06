# List of contents
- [What is Parabola](#what-is-parabola)
- [Motivation](#motivation)
- [How it works](#how-it-works)
- [Installation](#installation)
  - [Building from source](#building-from-source)
- [Usage](#usage)
  - [Starting the server](#starting-the-server)
  - [Interacting with the server](#interacting-with-the-server)
  -	[Using parabola-cli](#using-parabola-cli)
- [Parabola commands](#parabola-commands)
  - [set](#set)
  - [get](#get)
  - [del](#del)
- [Parabola protocol](#parabola-protocol)

# What is Parabola?
Parabola is an in-memory data store, currently it can be used as a database cache. In future, I plan to add support for more data structures, but until it matures and all the bugs are ironed out, i plan to keep it simple.
# Motivation
My plan for writing Parabola is not to overthrow industry standards like redis and memcached, I don't want to and can't compete with them. My motivation for writing Parabola is to provide a simple solution for all the caching needs of small to medium sized apps, which previously may not have considered the existing options due to the complexity or other problems.
# How it works?
At it's core, Parabola is just a server which operates on TCP. Currently it supports key-value caching, with LRU eviction policy. Parabola uses a custom communication protocol for the client-server communication. If you wish to read more about the implementation details feel free to read ahead.

# Installation
You can download the latest release from the releases page.
## Building from source
- Clone the repository and `cd` into it.
- Make sure you have latest version of `go` and `make` installed.
- From the root of the repository, run the `make` command.
- The compiled binaries will be present in the `bin` directory.

# Usage
## Starting the server
The Parabola server takes two optional flags.
- `port` is the port on which Parabola runs, by default it is `9898`, you can use any available port.
- `mem` is the max memory available to Parabola server in `bytes`, by default it's `1GB`. It's just a limit for the Parabola server and does not affect the performance of your system.
Example:
```bash
$ parabola -port=9000 -mem=1073741824
```
## Interacting with the server
Though you can use traditional tools like `telnet` or `ncat` to interact with the server, it requires knowledge about the Parabola protocol and it's pretty easy to mess up and might result in undesireable behaviour. I suggest using the `parabola-cli` to interact with the Parabola server which takes care of all the encoding and decoding.
## Using `parabola-cli`
`parabola-cli` takes one optional flag, that is the port that Parabola server is running on.
Example:
```bash
$ parabola-cli -port=9000
```

# Parabola commands
Currently there are 3 basic commands
- `set` which inserts a key-value pair with a ttl.
- `get` to retrieve the value of a specified key.
- `del` to delete a specified key.

## `set`
The set command is of the form:
```
> set [key] [value] [ttl (in seconds)]
```
All of the arguments are required.
Example:
```
> set mykey myvalue 120
```
This sets the following relationship.
`mykey` ---120---> `myvalue`
Currently the keys are only checked for expiration when they are queries, i am working on a solution which will poll keys in the background and evict expired keys.

## `get`
The get command is of the form:
```
> get [key]
```
All of the arguments are required.
Example:
```
> get mykey
```
This will return the value associated with the key.

## `del`
The del command is of the form:
```
> del [key]
```
All of the arguments are required.
Example:
```
> del mykey
```
This will delete the key from the key space.

# Parabola protocol
The Parabola protocol is a modified version of RESP.
There are 5 data types
- Strings
- Integers
- Arrays
- Errors
- Nil

The first byte of the message contains the data type of the message. Following are the data type specfiers. 
- Strings begin with `$`
- Integers begin with `%`
- Arrays begin with `#`
- Errors begin with `!`
- Nil begins with `-`

CRLF is used as a delimiter.
## Strings
All the strings in Parabola are binary safe, unlike RESP, Parabola protocol does not allow unsafe strings. 
The length of the string is embedded within the message.
Example:
`thetinygoat` will be encoded as
```
$11
\r\n
thetinygoat
\r\n
```
Where `$` specifies that it's a string followed by `11` which is the length of the string. It's written on separate lines for better readability, in real it would be written as `$11\r\nthetinygoat\r\n`.

## Integers
Integers are specified by `%` and are fairly straight forward.
Example:
```
%1000
\r\n
```
Where `%` is the specifier, `1000` is the encoded number followed by a CRLF.

## Arrays
Arrays are specified with `#` they can contain data of any of the other four types. Parabola protocol does not support nested arrays and `Nil` arrays.
Example:
```
#2
\r\n
$11
\r\n
thetinygoat
\r\n
%2
\r\n
```
Well this is a handful and might be a little complicated to understand, so let's break it down.
`#` specifies that it's an array followed by a `2` which is the number of elements in the array.
The first element is a string `$11\r\nthetinygoat\r\n` and the second element is an integer `%2\r\n`
Again this would be written  as:
 `#2\r\n$11\r\nthetinygoat\r\n%2\r\n`
## Errors
Errors are specified with `!`, they are of the form `!error details\r\n`
which is pretty straight forward.
## Nil
Nil is just a string with `size < 0`.
Example:
`$-1\r\n`
