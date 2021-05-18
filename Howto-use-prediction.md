giovanism@akashicazure05: run-to-gke-plus-tensor-serving ~/R/B/b/cc> curl -v http://<tensorflow-service IP>:8501/v1/models/model:predict -d '{"instances":[["Ded"], ["Meninggol"]]}'       
*   Trying 35.247.137.111:8501...
* Connected to 35.247.137.111 (35.247.137.111) port 8501 (#0)
> POST /v1/models/model:predict HTTP/1.1
> Host: 35.247.137.111:8501
> User-Agent: curl/7.76.1
> Accept: */*
> Content-Length: 38
> Content-Type: application/x-www-form-urlencoded
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Tue, 18 May 2021 03:14:50 GMT
< Content-Length: 58
< 
{
    "predictions": [[0.223400027], [0.223400027]
    ]
* Connection #0 to host 35.247.137.111 left intact
}