# kubectl apply -f chain-deployment.yaml
# kubectl apply -f chain-service.yaml 
kubectl apply -f coordinator.yaml 
kubectl apply -f worker.yaml 
kubectl apply -f prometheus.yaml 
kubectl apply -f grafana.yaml 
