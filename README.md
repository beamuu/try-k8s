# Deploying Go service with Redis on Kubernetes

## 👷🏻 Go backend service example explained

This service connects to the Redis once the service is run. There are 2 API endpoints available for you to test if the connection between the service
and Redis (which will be deployed as an k8s pod later) works properly. Redis has a configured password as `helloworld`. This service needs to GET and SET
values on Redis with out any errors. You can SET a value through `HTTP POST /set?key=<your-key>&value=<your-value>` and GET a value `HTTP GET /get?key=<your-key>`. A `key` query string won't accept `""` (EMPTY STRING). If an internal error occurs, error will be returned in a http response with a status 500 and error message will be provided in a response body.


## 😂 Some short k8s definition
Kubernetes (k8s) is a kind of cluster that let us deploy pods(an abstraction of containers), services, ect., on it. 

## Project structures explained

- `/app` Go backend service source code.
  - The service uses clean architecture design. Try check [this](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) out.
- `/k8s` This one stores k8s object files.
  - `redis-development.yaml` Redis master pod(container).
  - `redis-secret.yaml` Redis secret which contains REDIS_PASSWORD (stored in base64 encoded).
  - `redis-service.yaml` Redis service maps a internal traffic/ip to an Redis master pod. This allows us to access Redis master pod internally (Inside the cluster or Pod-Pod).
  - `webapp.yaml` Runs our backend services. This is a bad practice. I should split deployment and service in to two files separately.
- `Dockerfile` Use for building **Go backend service** image. 

You can also get my built image from [beamuuuu/redis-getset:latest](https://hub.docker.com/repository/docker/beamuuuu/redis-getset) or
```shell
docker pull beamuuuu/redis-getset
```
- `docker-compose.yml` Use for running Redis locally on docker container. (Not a k8s deployment!)

## Getting Started

You need [Minikube](https://minikube.sigs.k8s.io/docs/start/) and [kubectl](https://kubernetes.io/docs/tasks/tools/)

1. Create our new Kubernetes cluster
```
minikube start
```

2. Apply our k8s secrets, deployments, services, pods, ect., to our cluster. (I'm not gonna explain each yaml files because this is just an k8s setup tutorial)
```
sh k8s-apply.sh
```

3. Access the app service by using
```sh
minikube service webapp-service --url
// output like http://127.0.0.1:58081 or something.
```
You should get an URL (above) to your backend service. Then try to GET some keys through your browser.
Open your browser and go to `<your-url>/get?key=helloredis`. Then you should see something like this.
```
{
    "success":false,
    "message":"redis: nil",
    "payload":null
}
```
This means our backend works perfectly. The `redis: nil` in message means on Redis, there is no value for key `helloredis` which we sent it to the backend.

Now create a HTTP POST to `<your-url>/set?key=helloredis&value=welcome`. You should get
```
{
    "success": true,
    "message": "",
    "payload": null
}
```
This means Redis has set our key-value pair. { "helloredis": "welcome" }

Then go to your browser and enter `<your-url>/get?key=helloredis`
```
{
    "success":true,
    "message":"",
    "payload":"welcome"
}
```
As we can see that right now we have a payload as "welcome". This means once we request for a value of a key "helloredis", Redis returns us a "welcome" 
which was a value that we have set before.

## k8s works fine
If you have finished the `Getting Started` section, your Kubernetes tutorial is now completed.
I recommend anyone who are reading this to learn more about what Kubernetes is, how it works, and some important components inside Kubernetes.
I have some FREE useful resources for someone who are getting started on Kubernetes.
- [Kubernetes Crash Course for Absolute Beginners [NEW]](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kubernetes Course - Full Beginners Tutorial (Containerize Your Apps!)](https://www.youtube.com/watch?v=d6WC5n9G_sM&t=1645s)
- [Learn Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
- [Deploying Redis Cluster on Kubernetes | Tutorial and Examples](https://www.containiq.com/post/deploy-redis-cluster-on-kubernetes)
