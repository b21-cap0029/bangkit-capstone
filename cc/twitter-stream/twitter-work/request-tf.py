import requests
import json
import csv
import os
import os.path
from itertools import chain
from urllib.parse import urljoin

TENSORFLOW_BASE_URL = os.getenv("TENSORFLOW_BASE_URL")
AAIDA_BACKEND_BASE_URL= os.getenv("AAIDA_BACKEND_BASE_URL")
print(os.getcwd())
fileresponse = 'responses-record.csv'# file response
input_fetch = 'twitter_fetch.csv'
treshold = 0.936
#workdir = os.chdir("..")#naik satu folder untuk ambil twitter-fetch
#print(workdir)


#Persiapan Requests API
url = urljoin(TENSORFLOW_BASE_URL, '/v1/models/model:predict')  #dummy load untuk setting request tensor
url_aaida = urljoin(AAIDA_BACKEND_BASE_URL, 'cases/submit')  #dummy load dengan HTTP-POST
#REQUEST BUILDER



with open(input_fetch,'r') as f:
    csvReader= csv.reader(f)
    line_counter=0
    payload = {}
    for row in csvReader:
        if line_counter == 0:#untuk skip row dan penanda
            line_counter =+ 1
            print("this is first entry")
        else:
            #CREATE REQUEST AND PAYLOAD FOR TENSOR HERE
            print(row[3])
            data_input = row[3]
            request_input = {'instances':[[data_input]]}
            print(request_input)
            json_dumps = json.dumps(request_input)
            json_payload = json.loads(json_dumps)
            print(json_payload)
            #create request and catch the responses prepare the requests
            resp = requests.post(url,json=json_payload)#atur parameter request disini(doc with requests)
            print(resp.text)
            result_dict = resp.json()  #ambil hasilnya dalam bentuk json. cari key response(hasil test dengan curl
            print(result_dict)
            #hasilnya result_dict["predictions"]
            result = result_dict["predictions"]  #untuk mempermudah akses nilai
            result_list = list(chain.from_iterable(result))
            result = result_list[0]
            print(result_list)
            print(result)
            if result >= treshold:#sementara aja ini kalau resultnya dipasang treshold 0.5 dari 1
                print("ada masukan ke aaida-backend")#verbose just test in text
                ## disini untuk aaida access
                payload["tweet_id"] = int(row[2])# index[2] menyimpan tweet_id
                if float(result) >= treshold:#index[4] menyimpan score prediction
                    payload["class"] = "Teridentifikasi"
                else:
                    payload["class"] = "tidak teridentifikasi"
                payload["score"] = float(result)#index[4]menimpan score prediction
                payload["twitter_user_id"] = int(row[0])#menyimpan id twitter user
                payload["is_claimed"] = False
                payload["is_closed"] = False        
                #print(dict(payload)) #just test
                resp_aaida = requests.post(url_aaida,json=payload)
                if resp_aaida.status_code != 200:
                    print(resp_aaida.status_code)
                with open(fileresponse,'a',newline='') as responses:
                    csvWriter = csv.writer(responses)
                    print(row[0],row[1],row[2],row[3],result)#verbose all record(id_str,username,id_tweet,tweet_text, dan prediction score)(print only)
                    csvWriter.writerow([row[0],row[1],row[2],row[3],result]) 
            else:
                print("tidak ada isinya")#verbose mode
    

