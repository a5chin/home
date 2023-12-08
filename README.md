# Home

## docker-compose structure
- frontend:
  - node:20.0.0
  - next: 14.0.3
  - chakra-ui/next-js: 2.2.0
  - prettier: 3.1.0
- backend:
  - go: 1.20.0
  - air: 1.43.0
  - db
    - mysql: 8.0.28

## How to develop
```sh
$ docker-compose up
```

- frontend: http://localhost:3000
- backend: http://localhost:8080
