### Minikube Setup

- #### Run the chain:
  - Build the chain image using following command:
  ```console
      - docker build -f Dockerfile-chain  -t noblechain .
  ```
  - Deploy the chain using following cmd (from _/minikube-deployment_ dir):
  ```console
      - kubectl apply -f chain.yaml
  ```
- #### Run the load tests:
  - Build the docker images for coordinator, worker, prometheus, and grafana using following cmds:
  ```console
      - docker build -f Dockerfile-coordinator  -t noble-load-coordinator .
      - docker build -f Dockerfile-worker  -t noble-load-worker .
      - docker build -f Dockerfile-prometheus  -t load-test-prometheus .
      - docker build -f Dockerfile-grafana  -t load-test-grafana .
  ```
  - Deploy the coordinator, worker, prometheus server and grafana server by running the _load-test.sh_ script (from _/minikube-deployment_ dir)
  ```console
      - ./load-test.sh
  ```

  The prometheus server can be accessed at _<minikube_ip>:30000_
  
  The grafana server can be accessed at _<minikube_ip>:32000_ 
