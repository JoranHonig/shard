# Shard

Shard is a light mythril client
## Installation
```
snap install --jailmode --edge shard
```
## Configuration
You can put a config file in ```$HOME/.config/shard.yaml``` containing your api key.
shard.yaml contents:
```yaml
api-key: <put your api key here>
```
This way you don't have to put it in the cli every time. Alternatively shard also looks in the current directory for
a configuration file if it can't find one in the aforementioned directory.

## Usage
As any with any tool, the help command can be very useful
```
> $ ./shard                                                                                                                                                                              [±master ●]
Shard is a mythril light client

Usage:
  shard [command]

Available Commands:
  analyze     Analyzes the contract
  help        Help about any command
  version     Print the version number of Shard

Flags:
      --config string   config file (default is $HOME/.config/shard.yaml)
  -h, --help            help for shard
  -v, --verbose         Enable verbose logging.

Use "shard [command] --help" for more information about a command.

```

To analyze a contract execute:
```
> shard analyze 0x606b...
```
