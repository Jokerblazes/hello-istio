# hello-istio

## How To Test
1. Deploy client in same namespace(in client package)

2. run command(my namespace is hello-istio)

   1. get pod name

   `export SLEEP_POD=$(kubectl get pod -n hello-istio -l app=sleep -o jsonpath={.items..metadata.name})`

   2. curl

   `kubectl -n hello-istio exec -it $SLEEP_POD -c sleep -- sh -c 'curl  http://helloworld:6000/hello' | python -m json.tool`
   Or
   `kubectl -n hello-istio exec -it $SLEEP_POD -c sleep -- sh -c 'curl -H "end-user: dev1" http://helloworld:6000/hello' | python -m json.tool`

