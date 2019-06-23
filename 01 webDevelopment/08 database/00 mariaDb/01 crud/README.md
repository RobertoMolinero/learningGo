# Using Curl

Get Lines

```
curl http://localhost:8080/line
```

Get Line with Id

```
curl http://localhost:8080/line/1
```

Create Line

```
curl -X POST -H "Content-Type: application/json" -d '{"Publisher":"Marvel","Title":"X-Men","Id":17}' http://localhost:8080/line
```

Update

```
curl -X PUT -d '{"Publisher":"Marvel","Title":"X-Men","Id":1}' http://localhost:8080/line
```

Delete Line

```
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/line/20
```
