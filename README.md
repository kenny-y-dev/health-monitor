# health-monitor

health-monitor is a lightweight tool written in Go for easily monitoring the status of a service and reporting via HTTP if the status changes.

**This project has not reached an initial stable release, and all features and implementations are subject to change**

# Features

- Check and endpoint via ping to check for response
- Alert via various HTTP methods in case of health check failure

# Configuration

health-monitor is configured via the following environment variables:

- `MONITOR_TARGET` ip address or dns name of service to monitor
- `MONITOR_STRICT` determines if a single monitor test failure will send alert. Bool, defaults to `true`
- `MONITOR_TIMEOUT` timeout of monitor test sent to target, formatted as time string (such as `5s` for 5 seconds). Defaults to 5 seconds
- `NOTIFY_TARGET` HTTP endpoint to send notification on monitor health change
- `NOTIFY_METHOD` HTTP method to use when sending notification. Supports `GET`, `POST`, and `PUT`
- `NOTIFY_UP_JSON` JSON formatted payload to send to notify target when monitor reports status as up
- `NOTIFY_DOWN_JSON` JSON formatted payload to send to notify target when monitor reports status as down