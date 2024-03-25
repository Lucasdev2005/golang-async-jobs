docker compose up -d
echo "--------------------------------------------------------------------------------------"
echo "Load testing with Grafana dashboard http://localhost:3000/d/k6/k6-load-testing-results"
echo "--------------------------------------------------------------------------------------"
docker compose run --rm k6 run -e TEST_URL=http://publisher:8080 /scripts/transfer.test.js