
## filter(1)

  Filter JSON streams using javascript.

  ```bash
  $ echo '{"event":true}{"event":1}{}' | filter "if (j.event) return j"
  {"event":true}
  {"event":1}
  ```

## Installation
  
  ```bash
  $ go get github.com/segmentio/filter
  ```

## Usage

  ```bash
  Usage: 
    filter <js>

    filter -h | --help
    filter -v | --version

  Options:
    -h, --help      show help information
    -v, --version   show version information
  ```

## License

  MIT