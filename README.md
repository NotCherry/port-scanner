# Intentions
To better understand golang i rewrited port scanner that i created in rust and add more functionalitis

## Keypoint
- Multithreading and wait groups (sync.WaitGroup)
- Dealing with multiple gorutines writes to one variable (sync.Mutex)
- Usage of orm in go (gorm)

# Usage
- ```port-scanner -p 22-25,80 -t 0.5 192.168.0.1-255```

## Flags
- ```-p ``` ports to sscan single or range separate by ```,``` range with ```-```
- ```-t``` scan range
- ```-T``` timeout in seconds
- ```-s``` save output to sqlite database

# 