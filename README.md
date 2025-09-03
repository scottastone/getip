# getip

`getip` is a simple Go command-line tool to retrieve your public IP address. It makes an HTTP request to a public IP service and prints the result.

## Features

- Fetches your public IP address using a reliable external service.
- Minimal dependencies.
- Fast and lightweight.

## Usage

```bash
go run getip.go
```

Or build and run:

```bash
go build getip.go
./getip
```

## Example Output

```
Your public IP is: 203.0.113.42
```

## Requirements

- Go 1.18 or newer

## License

MIT