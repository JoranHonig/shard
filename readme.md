# Shard

Shard is a light mythril client that is amazingly awesome

## Configuration
You can put a config file in ```$HOME/.config/shard.yaml``` containing your api key.
This way you don't have to put it in the cli every time. Alternatively shard also looks in the current directory for
a configuration file if it can't find one in the aforementioned directory.

## Usage
As any with any tool, the help command can be very useful
```
shard help
```

[WIP] To analyze a contract execute:
```
> shard analyze -f filename.sol
# or
> shard analyze -c 0x606b...
```