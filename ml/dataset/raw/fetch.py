import csv
import os

import requests
import time

BEARER_TOKEN = "BEARER_TOKEN"


def read_keywords():
    f = open('keywords.txt', 'r')
    keywords = f.readlines()
    f.close()
    return [word.strip() for word in keywords]


def negate_keywords_query(keywords):
    return ' '.join(['-' + x for x in keywords])


def include_keywords_query(keywords):
    return '({})'.format(' OR '.join(keywords))


def create_url(next_token=None):
    keywords = read_keywords()
    keywords_query = include_keywords_query(keywords)

    query = '{} -is:retweet -is:reply -has:mentions -has:links ' \
            'lang:id'.format(keywords_query)
    tweet_fields = "tweet.fields=id,author_id,text,public_metrics"

    if next_token is None:
        url = "https://api.twitter.com/2/tweets/search/recent?query={}&" \
              "{}".format(query, tweet_fields)
    else:
        url = "https://api.twitter.com/2/tweets/search/recent?query={}&" \
              "{}&next_token={}".format(query, tweet_fields, next_token)
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


def get_token_if_exists():
    token = None
    try:
        token_file = open('next_token.txt', 'r')
        token = token_file.readlines()[0]

        if len(token) < 3:
            token = None

        token_file.close()
    except IOError:
        print("token file not found. assuming new state")

    return token


CSV_COLUMNS = 'id,author_id,like_count,quote_count,reply_count,retweet_count,text'


def append_response_to_csv(json_response):
    if not os.path.exists('result.csv'):
        with open('result.csv', 'w') as f:
            f.write(CSV_COLUMNS + '\n')

    with open('result.csv', 'a') as f:
        writer = csv.writer(f, delimiter=',')

        for entry in json_response["data"]:
            writer.writerow([entry["id"], entry["author_id"],
                             entry["public_metrics"]["like_count"],
                             entry["public_metrics"]["quote_count"],
                             entry["public_metrics"]["reply_count"],
                             entry["public_metrics"]["retweet_count"],
                             entry["text"].strip().replace("\n", ".")])

    print("written {} rows".format(json_response["meta"]["result_count"]))


def update_token(new_token):
    with open('next_token.txt', 'w+') as f:
        f.write(new_token)


FETCH_MAX = 10  # get FETCH_MAX statuses in one run
SLEEP_TIME = 25.0  # +5 seconds for leniency


def fetch_status(headers, next_token):
    count = 0
    print('fetching {} data'.format(FETCH_MAX))
    while count < FETCH_MAX:
        url = create_url(next_token)
        json_response = connect_to_endpoint(url, headers)
        append_response_to_csv(json_response)

        count += json_response["meta"]["result_count"]
        print("count is now {}".format(count))

        next_token = json_response["meta"]["next_token"]
        print("next token is {}".format(next_token))
        update_token(next_token)

        if count < FETCH_MAX:
            print("sleeping for {} seconds..".format(SLEEP_TIME))
            time.sleep(SLEEP_TIME)
        else:
            os.remove('next_token.txt')
            print("finished. total new rows: {}".format(count))


def main():
    headers = create_headers(BEARER_TOKEN)
    next_token = get_token_if_exists()
    fetch_status(headers, next_token)


if __name__ == "__main__":
    main()
