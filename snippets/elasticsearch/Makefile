up:
	sudo podman-compose -f 1-node.yml up

upd:
	sudo podman-compose -f 1-node.yml up -d

down:
	sudo podman-compose -f 1-node.yml down

del:
	sudo podman-compose -f 1-node.yml down -v

getca:
	sudo podman cp elasticsearch_es01_1:/usr/share/elasticsearch/config/certs/ca/ca.crt .