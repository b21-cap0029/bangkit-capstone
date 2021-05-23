#prepare for dependencies from tweepy
#import credentials.py
import tweepy
from tweepy import Stream
from tweepy import OAuthHandler
from tweepy.streaming import StreamListener
import time
import json

#from secret-key import key(remove the # to provide the key)
#CONSUMER_KEY=""
#CONSUMER_SECRET=""
#ACCESS_TOKEN_KEY=""
#ACCESS_TOKEN_SECRET=""
#how to separate key from code?


class listener(StreamListener):
    def on_data(self,data):
        #print(data) for verbosing only
        all_data = json.loads(data)
        tweet = all_data["text"]
        username = all_data["user"]["screen_name"]
        #print((username,tweet))
        return True

    ##def on_error(self, status):
    ##    print (status.id)


#authentication to twitter
auth = OAuthHandler(CONSUMER_KEY,CONSUMER_SECRET)
auth.set_access_token(ACCESS_TOKEN,ACCESS_TOKEN_SECRET)

#stream<?> the tweets and and save the result<?>
twitterStream = Stream(auth, listener())



#catch the prediction responses


#process the prediction