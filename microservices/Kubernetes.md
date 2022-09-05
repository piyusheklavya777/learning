# Kubernetes
- Folder : infra/k8s/<servicename>.yml
- start a kubernetes pod from a config file :
    - kubectl apply -f ..pathtoyml
- stop a pod started from a config :
    - kubectl delete -f infra/k8s/
- get the list of pods running
    - kubectl get pods
- to run a command inside a pod :
    - kubectl podname exec -t <sh or whatever command>
- to get all the logs from a pod :
    - kubectl logs podname
- describe pods :
    - kubectl describe pod <podname>

# Kubernetes deploymnents (alias kubectl as k)
- see the deployment script :
- k get deployments
- k describe deployments
- k apply -f fileWithKindDeployment.yml
- k delete deployment

# Updating a deployment
- rebuild the docker image with the new changes and push to docker hub.
- In the deployment plan, while specifying the name of the image for the container in the pod, append with :latest
- k rollout restart deployments <deployment-name>

# Services
- four types of communications
    - clusterIP
    - node port
    - load balancer
    - external name
- Created by the same yml file with kind: Service. Launched by the same k apply -f file.yml command
- get by k get services
- the external port is of the range 3xxxx. The port we specify as input mapping is not the one we can send requests on.
- get this input port from the report on terminal when service is started or by k get services
    - posts-service   NodePort    10.108.60.92   <none>        4000:31629/TCP   9s ::: input port here is 31629

# Side notes
- Try to get kubernetes dashboard

- Creating a secret in kubernetes :
    - kubectl create secret generic jwt-secret --from-literal=JWT_KEY=asdf
    - kubectl get secrets
    - If we try to get the reference to 