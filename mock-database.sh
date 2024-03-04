docker run -d \
    --name asyncjobs-postgres \
    -e POSTGRES_PASSWORD=asyncjobs@123 \
    -e POSTGRES_USER=asyncjobs \
    -e POSTGRES_DB=asyncjobs \
    -p 5432:5432 \
    postgres:latest