# Garage Door Server

## Setup Systemd Service

To set up systemd to automatically start the garage door server, follow these steps:

### 1. Copy the service file to systemd directory
```bash
sudo cp garage.service /etc/systemd/system/
```

### 2. Set proper permissions
```bash
sudo chmod 644 /etc/systemd/system/garage.service
```

### 3. Reload systemd to recognize the new service
```bash
sudo systemctl daemon-reload
```

### 4. Enable the service to start on boot
```bash
sudo systemctl enable garage.service
```

### 5. Start the service immediately
```bash
sudo systemctl start garage. service
```

### 6. Verify the service is running
```bash
sudo systemctl status garage. service
```

## Managing the Service

### Stop the service
```bash
sudo systemctl stop garage.service
```

### Restart the service
```bash
sudo systemctl restart garage.service
```

### Disable auto-start on boot
```bash
sudo systemctl disable garage.service
```

### View service logs
```bash
sudo journalctl -u garage.service -f
```

## Prerequisites

Make sure the `/root/garage` executable exists and has execute permissions before starting the service: 

```bash
chmod +x /root/garage
```

## Build

```bash
GOOS=linux GOARCH=amd64 go build -o garage .
```

## Docker

Build the image with Docker:

```bash
docker build -t garage:latest .
```

Run with Docker:

```bash
docker run --rm -p 8090:8090 garage:latest
```

Or use docker-compose:

```bash
docker-compose up --build
```
