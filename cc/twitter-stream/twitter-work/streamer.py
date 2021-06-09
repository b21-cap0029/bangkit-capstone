#prepare for dependencies from tweepy

import sys
import os
import csv
import datetime
import json
import time
import flask
import tweepy
from tweepy import OAuthHandler, Stream
from tweepy.models import Status
from tweepy.streaming import StreamListener
import re
from itertools import chain
from urllib.parse import urljoin
import requests
import os.path
from flask import Flask, redirect,url_for
from flask import request
#flask starter
app = Flask(__name__)

from API_KEY import (ACCESS_TOKEN, ACCESS_TOKEN_SECRET,
                    CONSUMER_KEY, CONSUMER_SECRET)

#ACCESS_TOKEN = os.getenv("ACCESS_TOKEN")
#ACCESS_TOKEN_SECRET = os.getenv("ACCESS_TOKEN_SECRET")
#CONSUMER_KEY = os.getenv("CONSUMER_KEY")
#CONSUMER_SECRET = os.getenv("CONSUMER_SECRET")


#GLOBAL PARAMS & VARIABLES
#builder tensor requests
TENSORFLOW_BASE_URL = os.getenv("TENSORFLOW_BASE_URL")
#tensor_url = urljoin(TENSORFLOW_BASE_URL, '/v1/models/model:predict')#prod 
tensor_url = urljoin("https://tensorflow-serving-4tl56tjpnq-as.a.run.app",'/v1/models/model:predict')#test purpose

#builder aaida-access
AAIDA_BACKEND_BASE_URL= os.getenv("AAIDA_BACKEND_BASE_URL")
#url_aaida = urljoin(AAIDA_BACKEND_BASE_URL, 'cases/submit')  #prod
url_aaida = urljoin("https://aaida-backend-4tl56tjpnq-as.a.run.app",'/cases/submit')#test purpose

#parameter extra
treshold = 0.936#Comply dengan model prediction
##set boundary entries
language_setup = ['id']
keyword_track = ['kecemasan','lelah']

def prediction(text_tweet,tensor_url):
    request_input = {'instances':[[text_tweet]]}
    json_dumps = json.dumps(request_input)
    json_payload = json.loads(json_dumps)
    resp = requests.post(tensor_url,json=json_payload)#atur parameter request disini(doc with requests)
    result_dict = resp.json()  #ambil hasilnya dalam bentuk json. cari key response(hasil test dengan curl
    #hasilnya result_dict["predictions"]
    result = result_dict["predictions"]  #untuk mempermudah akses nilai
    result_list = list(chain.from_iterable(result))
    result = result_list[0] #hasil akhirnya ada di result
    print(result)
    return result

def case_submit(id_user,id_tweet,pred_score,url_aaida,treshold):
    payload = {}
    payload["tweet_id"] = int(id_tweet)
    if float(pred_score) >= float(treshold):#untuk menandai 
        payload["class"] = "Teridentifikasi"
    else:
        payload["class"] = "tidak teridentifikasi"
    payload["score"] = float(pred_score)#index[4]menimpan score prediction
    payload["twitter_user_id"] = int(id_user)#menyimpan id twitter user
    payload["is_claimed"] = False
    payload["is_closed"] = False        
    #print(dict(payload)) #just test
    resp = requests.post(url_aaida,json=payload)
    print(f'{resp.status_code=} {resp.text=}')
    return resp.status_code


def deEmojify(text):
    regrex_pattern = re.compile(pattern = "["
        u"\U0001F600-\U0001F64F"  # emoticons
        u"\U0001F300-\U0001F5FF"  # symbols & pictographs
        u"\U0001F680-\U0001F6FF"  # transport & map symbols
        u"\U0001F1E0-\U0001F1FF"
        u"\U0001F700-\U0001F77F"  # alchemical symbols
        u"\U0001F780-\U0001F7FF"  # Geometric Shapes Extended
        u"\U0001F800-\U0001F8FF"  # Supplemental Arrows-C
        u"\U0001F900-\U0001F9FF"  # Supplemental Symbols and Pictographs
        u"\U0001FA00-\U0001FA6F"  # Chess Symbols
        u"\U0001FA70-\U0001FAFF"  # Symbols and Pictographs Extended-A
        u"\U00002702-\U000027B0"  # Dingbats
        u"\U000024C2-\U0001F251"# flags (iOS)
                           "]+", flags = re.UNICODE)
    return regrex_pattern.sub(r'',text)

def pre_process(text):
    # remove links
    text = re.sub("https://t.co/\S*","",text) 

    #remove @<blabla>
    text = re.sub("\@\S+","",text)

    #remove RT marker
    text = re.sub("RT","",text)

    # remove newline
    text = re.sub("\n","",text)

    #  remove digits
    text = re.sub("[0-9]","",text)

    # remove emojis
    text = deEmojify(text)

    return text

#setup listener
class Listener(StreamListener):

    #def on_status(self, status):
     #   print(status.text)


    def on_error(self, status_code):
        if status_code !=420:
            print(status_code)
            return False
        else:
            return True

    
    def on_data(self, data):
        json_tweet = json.loads(data)
        #print(json_tweet)
        id_user = json_tweet['user']['id']
        id_tweet = json_tweet['id']
        name = json_tweet['user']['screen_name']
        text_tweet = json_tweet['text']
        text_tweet = pre_process(text_tweet)
        print(text_tweet)
        pred_score = prediction(text_tweet=text_tweet,tensor_url=tensor_url)
        case_submit(id_user=id_user,id_tweet=id_tweet,pred_score=pred_score,url_aaida=url_aaida,treshold=treshold)
        """ with open(filesave,'a', newline='') as f:
            csvWriter = csv.writer(f)
            csvWriter.writerow([id_user,name,id_tweet,text_tweet])"""

@app.route('/main', methods = ['GET'])
def main():
    #authentication to twitter
    auth = OAuthHandler(CONSUMER_KEY,CONSUMER_SECRET)
    auth.set_access_token(ACCESS_TOKEN,ACCESS_TOKEN_SECRET)
    api = tweepy.API(auth)

    #flask work
    #content = request.data()
    #key_track = content["key_track"]

    #check API credentials
    try:   
        api.verify_credentials()
        print("Authentication Success")
    except:
        print("Authentication error")


    #creating the stream
    customListener = Listener()
    customStream = Stream(auth, listener = customListener)

    #start streaming
    try:
        customStream.filter(track=keyword_track,is_async=True)
        runtime = 120 #menandai lamanya 2mnt
        time.sleep(runtime)
        customStream.disconnect()
    except tweepy.TweepError as e:
        print("error occured")
        print(e.api_code,e.args,e.reason,e.response)
        return("error occured, "+ e.__class__,"worker stop now")
    else:
        return ("work done for "+ str(runtime))
    #filter track dalam list


@app.route('/', methods=['GET'])
def welcomer():
    return "welcome in"

if __name__ == "__main__":
    app.run(port=9000)