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
