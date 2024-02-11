docker-compose down
docker-compose --profile dev up -d --build
docker compose exec -ti dev-app /bin/bash
