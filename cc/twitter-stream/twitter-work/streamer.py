#prepare for dependencies from tweepy
#import credentials as cred
import tweepy
from tweepy import Stream
from tweepy import OAuthHandler
from tweepy.streaming import StreamListener
import time
import json
import datetime
import csv


#environment variables


#from secret-key import key(remove the # to provide the key)
#CONSUMER_KEY="oAGvuZiLB56WqgZPdb9ogZJ8y"
#CONSUMER_SECRET=""
#ACCESS_TOKEN_KEY=""
#ACCESS_TOKEN_SECRET=""
#how to separate key from code?

##set boundary entries
language = ['id']
#coordinates = []

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
    reset_block_counter = 100
    #to avoid loop of unlimited retries connection
    #set to reconnect trial to n


    def on_status(self, status):
        print(status)

    def on_error(self, status_code):
        if status_code !=420:
            print(status_code)
            return False #reconnect stream with 3 tries
        else:
            return True

    
    def on_data(self, data):
        self.data_payload =+ 1
        tweet = json.loads(data)
        username = tweet.get("user",{}).encode('utf-8').strip()
        text = tweet.get("text","").encode('utf-8').strip()
        _id = tweet.get("user",{}).get("id","")
        #print username



#creating the stream
customListener = Listener()
customStream = Stream(auth= api.auth, listener = customListener)

#start streaming
try:
    customStream.filter(languages=['id'])#batas bahasa indonesia
    #adding file csv
    #filesave ='twitter_fetch'+(datetime.datetime.now().strftime("%Y-%m-%d %H:%M"))+'.csv'
    filesave = 'twitter_fetch.csv'
    #storing the tweet fetched
    with open (filesave, 'a+', newline='') as csvFile:
        csvWriter = csv.writer(csvFile)
        for tweet in tweepy.Cursor(customStream).items():
            tweets_encoded = tweet.text.encode('utf-8')
            tweet_decoded = tweets_encoded.decode('utf-8')
            csvWriter.writerow([datetime.datetime.now().strftime("%Y-%m-%d %H:%M"), tweet.id, tweet.author.name , tweet_decoded,tweet._json["user"]["location"]])
except KeyboardInterrupt as e:
    print("stopped")
finally:
    print('done')
    customStream.disconnect()
    
#catch the prediction responses


#process the prediction