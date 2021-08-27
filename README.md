# SejutaCita Technical Test

## Technical Test Requirements
1. Dengan bahasa apa saja, diutamakan NodeJS, agar saudara membuat Rest API CRUD User dan User Login.
2. Framework dibebaskan, tetapi disarankan ExpressJS. Database bebas, tetapi disarankan MongoDB.
3. User Login digunakan user (username, password) untuk mengakses API CRUD (token, tetapi mendapatkan nilai tambahan jika menggunakan refresh token).
5. Bikin 2 users dengan role: 1 Admin, 1 User.
6. Admin bisa melakukan/mengakses semua API CRUD, sedangkan User hanya bisa mengakses data user bersangkutan saja (Read)
7. Implementasi arsitektur Microservices, menggunakan Kubernetes dengan Docker container deploy di VPS (1 node dengan beberapa pod di dalamnya). Bagi yang belum memiliki VPS, maka cukup (a) menyiapkan semua YML agar aplikasi bisa dijalankan secara containerize dan siap di deploy di Kubernetes dan (b) di-deploy di lokal dan sertakan screenshoot.
8. Upload source code ke Github beserta script YML Kubernetes.
9. Bikin dokumentasi API nya (Postman atau Swagger) yang bisa diakses ke server Rest API nya.
10. Bikin diagram arsitektur nya yang menjelaskan flow API CRUD dan Login.
11. Lampirkan credential Admin di Readme.

The API - Simple User API
A user simple API build using Go gin framework with PostgreSQL. Can support Docker containerization and Kubernetes.

### There is two module for this API:

AuthNode, contains API for user authentication,
POST .../users/auth/login
POST .../users/auth/register
UserNode, contains API for user CRUD operation
GET .../users/
GET .../users/:id
PUT .../users/:id
DELETE .../users/:id

The user authentication is implemented using JWT mechanism, and user RBAC using casbin.

The API request examples can be seen at doc. The exported Postman documentation can also be downloaded from .

The app were made this way so it can be scale independently. E.g. we can have two user node and one auth node. Just change the replication from 1 to 2 for the user node.

#### System Requirements
This application has been tested on Ubuntu 20.04 with these settings

> $ go --version  
> go version go1.16.5 darwin/amd64

> $ docker --version
Docker version 20.10.3, build 48d30b5

> psql (PostgreSQL) 12.7

### How to Install
#### Get the Code 
Make sure you already install all the requirements then clone or download the repository.

git clone https://github.com/iqbal482/goginCasbin.git

#### Install Application
First setup postgresql using helm kubernetes.
> $ helm repo add bitnami https://charts.bitnami.com/bitnami

> $ kubectl create namespace postgre

> $ helm install postgresql bitnami/postgresql -n postgre --set storageClass=hostPath --set global.postgresql.postgresqlPassword=q6PENVTLBW"

Second running kubernetes Go Gin using helm kubernetes.

> $ helm install gogin kubernetes 

```
Release "gogin" has been upgraded. Happy Helming!
NAME: gogin
LAST DEPLOYED: Fri Aug 27 15:39:00 2021
NAMESPACE: default
STATUS: deployed
REVISION: 6
NOTES:
1. Get the application URL by running these commands:
  export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services gogin-kubernetes)
  export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
  echo http://$NODE_IP:$NODE_PORT
```

### Future Improvement
1. Optimize the restful API
2. Added for refresh token
3. Added environment Variable