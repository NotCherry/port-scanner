# Intentions
To better understand golang i wrote port scanner

## Keypoint
- Concurrency and wait groups (sync.WaitGroup)
- Dealing with multiple gorutines writes to one variable (sync.Mutex)
- Logging with channels (chan)
- Usage of orm in go (gorm)

# Usage
- ```port-scanner -p 22-25,80 -t 0.5 192.168.0.1-255```

## Flags
- ```-p ``` ports to scan single or range separate by ```,``` range with ```-```
- ```-t``` scan range
- ```-T``` timeout in seconds
- ```-s``` save output to sqlite database
