
##  执行命令

```
kubectl apply -f tekton-cicd/pipeline/build-and-deploy-pipeline.yaml -f tekton-cicd/resources/picalc-git.yaml -f tekton-cicd/tasks/deploy-using-kubectl
kubectl apply -f egress.yaml -f clusterrole.yaml -f clusterrolebinding.yaml -f serviceaccount.yaml -f configmap.yaml -f service.yaml
```
