# Go API Template

- Added script for migration and seeding
- Added support for hot reload
- Added script for testing
- Added example feature: user registration and jwt login
- Added middlewares: auth, timeout, error, logging
- Configured server with graceful shutdown
- Tested in go1.18.10

## Usage

```bash
# in new project
git init

git remote add skeleton git@github.com:alvinfebriando/gin-gorm-skeleton.git
git pull skeleton
git checkout skeleton
# update go module
MODULE=example.com/example-project make rename
git add .
git commit
git checkout -b main

cp .env.example .env
# change env value
nano .env

# hot reload
make reload
```