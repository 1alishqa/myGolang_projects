# Ascii - Art
This project consists on receiving a string as an argument and outputting the string in a graphic representation of ASCII. Also, the second argument must be the name of the template.

### Usage
```
name@ubuntu:~/ascii-art$ go build
name@ubuntu:~/ascii-art$ ./ascii-art "hello" standard
  _                _    _
 | |              | |  | |
 | |__      ___   | |  | |    ___
 |  _ \    / _ \  | |  | |   / _ \
 | | | |  |  __/  | |  | |  | (_) |
 |_| |_|   \___|  |_|  |_|   \___/


name@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow

_|    _|          _| _|                _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|


name@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" thinkertoy

o  o     o o           o-O-o o
|  |     | |             |   |                o
O--O o-o | | o-o         |   O--o o-o o-o o-o |
|  | |-' | | | |         |   |  | |-' |   |-' o
o  o o-o o o o-o         o   o  o o-o o   o-o
                                              O

```
