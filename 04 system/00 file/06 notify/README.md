# Notifier for file system events

## Installation

... depends on ...

```
go get -u golang.org/x/sys/...
```

... install ...

```
go get -u ...
```




## Vergleich

```
if event.Op&fsnotify.Create == fsnotify.Create {
    log.Println("File created: ", event.Name)
}
```
