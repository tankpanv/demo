env=$1
service_name=federation-demo-recommend
s=`kubectl get pods|grep -v Terminating|grep ${service_name}-$env|awk '{print$1}'`
first_pod=`echo $s |awk '{print $1}'`
echo $first_pod
kubectl exec -it $first_pod -- /bin/bash
