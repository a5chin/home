# Backend
- go: 1.20.0
- air: 1.43.0
- db
  - mysql: 8.0.28

## How to Develop
```sh
## Example for backend/config/.env
HOSTNAME="127.0.0.1"
PORT="8080"
DB_HOSTNAME="db"
DB_USER="root"
DB_PWD=""
DB_NAME="home"
DB_PORT="3306"
```

```
docker-compose up backend
open http://localhost:8080
```
