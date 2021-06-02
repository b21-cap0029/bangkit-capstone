#prepare for dependencies from tweepy

import sys
sys.path.append('/.../twitter-work')
import os
import csv
import datetime
import json
import time
import tweepy
from API_KEY import (ACCESS_TOKEN, ACCESS_TOKEN_SECRET, CONSUMER_KEY,
                     CONSUMER_SECRET)
from tweepy import OAuthHandler, Stream
from tweepy.streaming import StreamListener
from dotenv import load_dotenv
load_dotenv()

##set boundary entries
language = ['id']
#coordinates = []
#filesave ='twitter_fetch'+(datetime.datetime.now().strftime("%Y-%m-%d %H:%M"))+'.csv'
filesave = 'twitter_fetch.csv'
#authentication to twitter
auth = OAuthHandler(CONSUMER_KEY,CONSUMER_SECRET)
auth.set_access_token(ACCESS_TOKEN,ACCESS_TOKEN_SECRET)
api = tweepy.API(auth)

#check API credentials

try:   
    api.verify_credentials()
    print("Authentication Success")
except:
    print("Authentication error")

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
        print(json_tweet)
        id_user = json_tweet['user']['id']
        id_tweet = json_tweet['id']
        name = json_tweet['user']['screen_name']
        text_tweet = json_tweet['text']
        print(text_tweet)
        with open(filesave,'a', newline='') as f:
            csvWriter = csv.writer(f)
            csvWriter.writerow([id_user,name,id_tweet,text_tweet])




#creating the stream
customListener = Listener()
customStream = Stream(auth, listener = customListener)

#membuat filesave
#BEWARE karena disini pakai mode W potensi ilang, ubah modenya untuk siap PROD
with open (filesave, 'w', newline='') as csvFile:
    csvWriter = csv.writer(csvFile)
    csvWriter.writerow(['ID_USER', 'USERNAME','ID_TWEET','TWEET_TEXT'])

#start streaming
#customStream.sample(languages=["id"])
customStream.filter(track=['election','loser'],languages=["en","fr","es","id"])#batas bahasa indonesia wajib pakai track(APIKEYnya tidak dapat akses feed semua tweet)
#filter track dalam list
