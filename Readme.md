# API

## Create a nodeport service
### /api/v1/createsvc

```
curl -XPOST -H "Content-Type: application/json" http://localhost:8080/api/v1/createsvc -d '{"port": 8080,"nodeport": 31191,"Namespace":"gladostest","Label": {"key": "app","value": "glados"}}'
```

### /api/v1/deletesvc
```
curl -XPOST -H "Content-Type: application/json" http://localhost:8080/api/v1/deletesvc -d '{"port": 8080,"nodeport": 31191,"Namespace":"gladostest","Label": {"key": "app","value": "glados"}}'
```

### /api/v1/getcache
Get service cache
```
curl http://localhost:8080/api/v1/getcache
```