# DXEP Specification

DXEP stands for DictX encoding protocol, it is the protocol used by client and server to communicate. DXEP is a simplified version of RESP.

## Data types

There are 4 data types supported by DXEP

- **Integers** - first byte of is "%"
- **Strings** - first byte is "\$". Strings are binary safe.
- **Arrays** - first byte is "#"
- **Errors** - first byte is "-"

## Integers

Integers are CRLF terminated strings.
Example:

- 0 is encoded as: `"%0\r\n"`.
- 1000 is encoded as : `"%1000\r\n"`.

## Strings

Strings are of the form `"$str_len\r\nstrings\r\n"`
Example:

- foobar is encoded as `$6\r\nfoobar\r\n`, where 6 is the length of the string.
- An empty string can be enocded as `$0\r\n\r\n`.

### Nil type

A nil type is just a string with length `-1`
Example:
`$-1\r\n\r\n`

## Arrays

Arrays are of the form `#arr_len\r\ndata`.
For example an array of strings can be encoded as:

```
#2\r\n
$6\r\nfoobar\r\n
$5\r\nfoobar\r\n
```

for the sake of documentation it is written in separate lines but in application it is just sent as a single string this, `#2\r\n$6\r\nfoobar\r\n$5\r\nfoobar\r\n`

The client sends command to the dictX server as array of strings for example

```
#2
$3\r\nGET\r\n
$3\r\nkey\r\n
```

again this is done just for the purpose of documentation and in reality it is sent as a single string.

## Errors

Errors are of the form `-Error message\r\n`. Here `Error` can be a generic error or a specific error, it's the job of the client to handle the error on the client side.
