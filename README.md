# ip-addr

`ip-addr` is a simple cli tool for printing the ip address for a 
given interface by specifying it's name.  This is useful for scripting 
so that cli flags and config files can be created with the correct
ip address in a scriptable way.

#### Help:

```bash
NAME:
   ip-addr - return the IP address for a given interface name

USAGE:
   ip-addr [global options] command [command options] [arguments...]

VERSION:
   1

AUTHOR:
  @crosbymichael - <crosbymichael@gmail.com>

COMMANDS:
   help, h      Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
``` 

#### Usage:

```bash
ip-addr eth0
192.168.1.10
```

### License - MIT

```
Copyright (c) 2015 Michael Crosby. crosbymichael@gmail.com

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation 
files (the "Software"), to deal in the Software without 
restriction, including without limitation the rights to use, copy, 
modify, merge, publish, distribute, sublicense, and/or sell copies 
of the Software, and to permit persons to whom the Software is 
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be 
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. 
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, 
DAMAGES OR OTHER LIABILITY, 
WHETHER IN AN ACTION OF CONTRACT, 
TORT OR OTHERWISE, 
ARISING FROM, OUT OF OR IN CONNECTION WITH 
THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
