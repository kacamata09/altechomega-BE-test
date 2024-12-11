### First Setup
1. Make copy from .env.yaml.example and rename to .env.yaml
1. You can settings your environment in env.yaml  
3. Run command below to install module dependency
```
go mod tidy
```

### Database migration
1. Create database as in your env.yaml
2. Change directory to migrations in terminal
```
cd migrations
```
3. Run file main.go in that directory
```
go run main.go
```

### Run App
Run file main.go in app directory  
```
go run main.go
```

### Add Docs
1. Run command below:  
```
swag init
```
2. Visit route /swagger/, example "http://localhost:9999/swagger/
