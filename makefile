.PHONY: intodb
intodb:
	docker container exec -it postgres bash

.PHONY: intogo
intogo:
	docker container exec -it todo_app-web-1 bash

.PHONY: upd
upd:
	docker compose up -d