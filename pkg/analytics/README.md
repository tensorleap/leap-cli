# Analytics Package

A minimal Mixpanel analytics package for tracking server installation events in the Leap CLI.

## Features

- Track server installation events only
- Automatic common properties (OS, architecture, timestamp, username)
- HTTP timeout handling
- Non-blocking (errors don't affect server operations)

## Usage

```go
import "github.com/tensorleap/leap-cli/pkg/analytics"

// Track server installation
properties := map[string]interface{}{
    "port": 8080,
    "gpus": 2,
    "cpu": true,
}
err := analytics.TrackServerInstall(properties)
```

## Event Properties

### server_install

Tracked when `leap server install` completes successfully.

**Custom Properties:**

- `port` - HTTP server port
- `data_dir` - Data directory path
- `domain` - Server domain
- `gpus` - Number of GPUs
- `gpu_devices` - GPU device specification
- `cpu` - Whether CPU mode is enabled
- `cpu_limit` - CPU resource limit
- `disable_metrics` - Whether metrics are disabled
- `airgap` - Whether air-gap installation
- `local` - Whether using local helm charts
- `tag` - Installation tag/version
- `tls_port` - TLS port
- `registry_port` - Docker registry port
- `clear_images` - Whether to clear images
- `proxy_url` - Whether proxy is configured
- `cert_file` - Whether TLS certificate is provided
- `key_file` - Whether TLS key is provided
- `chain_file` - Whether TLS chain is provided

**Automatic Properties:**

- `token` - Mixpanel project token
- `time` - Unix timestamp
- `os` - Operating system (e.g., "darwin", "linux")
- `arch` - Architecture (e.g., "amd64", "arm64")
- `timestamp` - RFC3339 formatted timestamp
- `username` - Current username (if available)
- `device_id` - Unique persistent device identifier (32-character hex string, stored in `~/.tensorleap/device_id`)

## Configuration

The Mixpanel token is hardcoded in the package. To change it, modify the `MixpanelToken` constant in `mixpanel.go`.

## Device ID Persistence

The device ID is automatically generated on first use and stored persistently in `~/.tensorleap/device_id`. This ensures:

- **Consistency**: Same device ID across all sessions and installations
- **Uniqueness**: Each device gets a truly random, unique identifier
- **Persistence**: Device ID survives system reboots and CLI updates
- **Privacy**: Random generation ensures no personal information is encoded

The device ID file is created with restricted permissions (0600) for security.

## Error Handling

Analytics tracking errors are logged as warnings but do not cause the main operations to fail. This ensures that analytics issues don't impact user experience.
