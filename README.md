# Microservice Boilerplate #

Repository ini merupakan template yang akan digunakan untuk membuat microservice baru. Template ini sudah terintegrasi dengan beberapa tools yang akan membantu developer dalam membuat microservice. Jika ingin membuat microservice dengan dependency, Anda dapat menggunakan library di repository ini tanpa perlu clone atau copy-paste library di sini ke repository Anda.

## Pembuatan Service Baru ##
Jika service yang dibuat memiliki banyak sub service yang saling berkaitan, maka sebaiknya dikumpulkan dalam satu repository. Package `main` dari service-service tersebut dikumpulkan di folder `service`. Pada template ini, sudah ada service `sum_caller` dan service `sum`.

### Build Image Service Baru DB2 ###

contoh script build:
```
docker build -t example-db2:latest -f ./service/example-db2/Dockerfile .
```

contoh script running:
```
docker run -p 8001:8001 --env DB_HOST=<<db host>> --env DB_PORT=<<db port>> --env DB_NAME=<<db name>> --env DB_USER=<<db user>> --env DB_PASSWORD=<<db password>> --env JAEGER_ENDPOINT=<<jaeger url>> --name=example-db2 example-db2:latest
```

isi value env sesuai config yang di gunakan 

### Build Image Service Baru Postgres ###

contoh script build:
```
docker build -t example-postgres:latest -f ./service/example-postgres/Dockerfile .
```

contoh script running:
```
docker run -p 8001:8001 --env DB_HOST=<<db host>> --env DB_PORT=<<db port>> --env DB_NAME=<<db name>> --env DB_USER=<<db user>> --env DB_PASSWORD=<<db password>> --env DB_TIMEZONE=<<db timezone>> --env JAEGER_ENDPOINT=<<jaegar_url>> --name=example-postgres example-postgres:latest
```

isi value env sesuai config yang di gunakan 


# Open Telemetry #
Untuk menggunakan open telemetry, Anda harus melakukan konfigurasi terlebih dahulu. Konfigurasi yang harus dilakukan adalah sebagai berikut:
```
url := os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT") // url jaeger misal: http://localhost:14268/api/traces
serviceName := os.Getenv("OTEL_SERVICE_NAME") // nama service, misal: user-service
environment := os.Getenv("OTEL_ENVIRONMENT") // environment misal: production, staging, development
tracer.InitOtel(url, serviceName, version, environment)
```