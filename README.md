# getip

`getip` is a simple Go command-line tool to retrieve your public IP address. It makes an HTTP request to a public IP service and prints the result.

It also knows to ignore local loopback addresses when listing local IPs.

## Features

- Fetches your public IP address using a reliable external service (ipify).
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
Public IPv4: 75.150.150.20
Local IPv4:  10.0.0.32
Local IPv4:  100.112.158.63
```

## Requirements

- Go 1.18 or newer

## License

MIT