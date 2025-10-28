# go-meo

A Go library and CLI tools for managing MEO router settings programmatically.

## Features

- Assign static IP addresses to devices by MAC address
- Configure WiFi settings (SSID, password, channels, bandwidth, transmit power)
- Support for both 2.4GHz and 5GHz bands

## Installation

```bash
go get github.com/pedrobarbosak/meo
```

## Usage

### Static IP Assignment

Assign static IP addresses to devices on your network based on their MAC addresses.

#### Configuration

Create a `config.json` file:

```json
{
  "username": "meo",
  "password": "meo",
  "hostname": "192.168.1.254",
  "macs": {
    "00:00:00:00:00:00": "192.168.1.100",
    "11:11:11:11:11:11": "192.168.1.101"
  }
}
```

#### Running

```bash
go run cmd/static-ip/main.go
```

### WiFi Settings Configuration

Configure your WiFi network settings including SSID, password, and band-specific settings.

#### Configuration

Create a `config.json` file:

```json
{
  "username": "meo",
  "password": "meo",
  "hostname": "192.168.1.254",
  "wifi": {
    "network": {
      "ssid": "MyNetwork",
      "password": "MyPassword"
    },
    "2.4ghz": {
      "bandwidth": 3,
      "channel": 6,
      "transmitPower": 100
    },
    "5ghz": {
      "bandwidth": 7,
      "channel": 100,
      "transmitPower": 100
    }
  }
}
```

#### Bandwidth Options

- **2.4GHz Band:**
  - `1` = 20MHz
  - `3` = 20MHz/40MHz

- **5GHz Band:**
  - `1` = 20MHz
  - `3` = 20MHz/40MHz
  - `7` = 20MHz/40MHz/80MHz

#### Running

```bash
go run cmd/wifi-settings/main.go
```

## Configuration Parameters

### Common Parameters

- `username`: Router admin username (default: "meo")
- `password`: Router admin password (default: "meo")
- `hostname`: Router IP address (default: "192.168.1.254")

### Static IP Parameters

- `macs`: Map of MAC addresses to IP addresses

### WiFi Parameters

- `wifi.network.ssid`: WiFi network name
- `wifi.network.password`: WiFi network password
- `wifi.2.4ghz.bandwidth`: Bandwidth setting for 2.4GHz band
- `wifi.2.4ghz.channel`: Channel number for 2.4GHz band
- `wifi.2.4ghz.transmitPower`: Transmit power percentage (0-100)
- `wifi.5ghz.bandwidth`: Bandwidth setting for 5GHz band
- `wifi.5ghz.channel`: Channel number for 5GHz band
- `wifi.5ghz.transmitPower`: Transmit power percentage (0-100)

## Library Usage

```go
import "github.com/pedrobarbosak/meo/pkg/meo"

meoClient, err := meo.New("username", "password", "192.168.1.254")
if err != nil {
    log.Fatal(err)
}

ctx := context.Background()
err = meoClient.AssignStaticIP(ctx, "00:00:00:00:00:00", "192.168.1.100")
if err != nil {
    log.Fatal(err)
}
```

## Requirements

- Go 1.24 or higher
- Access to MEO router admin interface

## License

See LICENSE file for details.

