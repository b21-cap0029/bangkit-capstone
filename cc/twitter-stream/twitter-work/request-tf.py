import requests
import json
import csv
import os
import os.path

print(os.getcwd())
#workdir = os.chdir("..")#naik satu folder untuk ambil twitter-fetch
#print(workdir)
#url="http://35.247.137.111:8501/v1/models/model:predict"#dummy load untuk setting request tensor 
test_payload = '{"instances":[["Ded"], ["Meninggol"]]}'

test = requests.post(url,data=test_payload)
print(test.status_code)


#path = "E:\BANGKIT\Project\bangkit-capstone\cc\twitter-stream" ##belum bisa akses relative path sama belum bisa buat workpathnya(aulah)
with open('twitter_fetch.csv','r') as f:
    csvReader= csv.reader(f)
    line_counter=0
    for row in csvReader:
        if line_counter == 0:#untuk skip row dan penanda
            line_counter =+ 1
            print("this is first entry")
        else:
            result = 1#test workload !!!Jangan lupa buang line ini kalau sudah bekerja API Requestnya
            #result = resp.responses
            if result >= 0.5:#sementara aja ini kalau resultnya dipasang treshold 0.5 dari 1
                with open('responses-record.csv','a',newline='') as responses:
                    csvWriter = csv.writer(responses)
                    #create request and catch the responses prepare the requests
                    payload = {}
                    payload["instances"] = [row[2]]
                    json_payload = json.dumps(payload)
                    #resp = requests.post(url,json=json_payload)#atur parameter request disini(doc with requests)
                    #result_dict = resp.json()  #ambil hasilnya dalam bentuk json. cari key response(hasil test dengan curl
                    #hasilnya result_dict["predictions"]
                    #result = result_dict["predictions"]  #untuk mempermudah akses nilai
                    print(row[0],row[1],row[2],result)#verbose all record
                    #csvWriter.writerow([row[0],row[1],result]) #opsi dengan line ini hanya menyimpan id_str,username, dan prediction score
    

