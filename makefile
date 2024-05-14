run-service:
	docker compose up -d
	docker build --progress=plain -t listingservice -f Dockerfile .
	docker run -p 8081:8081 -e "DEPLOY=dev" --network host listingservice:latest
