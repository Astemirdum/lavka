# Lavka services

Lavka is REST API service that registers couriers, add new orders and distribute them among couriers.

## The following concepts are applied in service:
- <a href="https://echo.labstack.com/">Echo</a> Framework 
- Clean Architecture
- Postgres <a href="https://github.com/volatiletech/sqlboiler">sqlboiler</a>.
- RateLimiter
- Swagger (OpenAPI)
- Docker compose
- CI/CD 
- K8s/helm

### Tasks
- [x] REST API
- [x] Courier Rating
- [x] Rate limiter
- [x] Orders allocations


### API
```vertica
http://localhost:8080/swagger/
```
or
```vertica
http://ingress-host:80/swagger/
```
#### couriers:
* POST /couriers
* GET /couriers/{courier_id}
* GET /couriers
* GET /couriers/meta-info/{courier_id}
#### orders:
* POST /orders
* GET /orders/{order_id}
* GET /orders
* POST /orders/complete
* POST /orders/assign
* GET /couriers/assignments

### Pre requirements
* golang 1.20
* minikube: `https://minikube.sigs.k8s.io/docs/start`
* helm: `https://helm.sh`


### Usage
#### run app
```sh
$ make run
```
#### linter
```sh
$ make lint
```
#### run tests
```sh
$ make test
```

Update helm dependencies:
```sh
$ eval $(minikube docker-env)
$ helm dep up ./helm/lavka-app/
```

Start the application:
```sh
$ make helm-init MY_RELEASE=my-release NAMESPACE=lavka-ns
```
