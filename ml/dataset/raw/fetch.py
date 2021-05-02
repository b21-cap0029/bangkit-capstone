import csv

import requests
import os
import json
import time

# To set your environment variables in your terminal run the following line:
# export 'BEARER_TOKEN'='<your_bearer_token>'

bearer = "BEARER_TOKEN"

def auth():
    return bearer
    # return os.environ.get("BEARER_TOKEN")


def create_url(next_token=None):
    query = "(hidup OR depresi OR cemas OR bunuh diri OR mati OR gantung diri OR " \
            "takut OR khawatir OR gelisah OR malas OR akhiri OR ga kuat OR ga " \
            "berarti OR selamat tinggal OR bye OR sedih OR hampa OR stress OR " \
            "kenapa hidup) -is:retweet -is:reply -has:mentions -has:links lang:id"
    # Tweet fields are adjustable.
    # Options include:
    # attachments, author_id, context_annotations,
    # conversation_id, created_at, entities, geo, id,
    # in_reply_to_user_id, lang, non_public_metrics, organic_metrics,
    # possibly_sensitive, promoted_metrics, public_metrics, referenced_tweets,
    # source, text, and withheld
    tweet_fields = "tweet.fields=id,author_id,text,public_metrics"
    if next_token is None:
        url = "https://api.twitter.com/2/tweets/search/recent?query={}&{}".format(
            query, tweet_fields
        )
    else:
        url = "https://api.twitter.com/2/tweets/search/recent?query={}&{}&next_token={}".format(
            query, tweet_fields, next_token
        )
    return url


def create_headers(bearer_token):
    headers = {"Authorization": "Bearer {}".format(bearer_token)}
    return headers


def connect_to_endpoint(url, headers):
    response = requests.request("GET", url, headers=headers)
    print(response.status_code)
    if response.status_code != 200:
        raise Exception(response.status_code, response.text)
    return response.json()


def main():
    bearer_token = auth()

    next_token = None

    with open('next_token.txt', 'r') as f:
        token = f.readlines()[0]

        if len(token) > 3:
            next_token = token

    count = 0
    fetch_max = 1000
    print('fetching {} data'.format(fetch_max))
    while count <= 1000:
        url = create_url(next_token)
        headers = create_headers(bearer_token)
        json_response = connect_to_endpoint(url, headers)

        with open('result.csv', 'a') as f:
            writer = csv.writer(f, delimiter=',')

            for entry in json_response["data"]:

                writer.writerow([entry["id"], entry["author_id"],
                                 entry["public_metrics"]["like_count"],
                                 entry["public_metrics"]["quote_count"],
                                 entry["public_metrics"]["reply_count"],
                                 entry["public_metrics"]["retweet_count"],
                                 entry["text"].strip().replace("\n", ".")])

        next_token = json_response["meta"]["next_token"]

        with open('next_token.txt', 'w+') as f:
            f.write(next_token)

        print("written {} rows".format(json_response["meta"]["result_count"]))
        count += json_response["meta"]["result_count"]
        print("count is now {}".format(count))
        print("next token is {}".format(next_token))
        print("sleeping for 25 seconds..")
        time.sleep(25.0)  # 5 seconds for leniency

if __name__ == "__main__":
    main()
