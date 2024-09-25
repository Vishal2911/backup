# 

 swag init --parseDependency  --parseInternal --parseDepth 1 -d api -g ../cmd/main.go
 go run cmd/main.go -log-level info

  http://localhost:8080/swagger/index.html