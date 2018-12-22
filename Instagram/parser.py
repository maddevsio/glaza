#!/usr/bin/python3.6

import sys
import datetime
from instaparser.agents import Agent
from instaparser.entities import Account, Media
from pymongo import MongoClient

def get_user_data(username):
    agent = Agent()
    account = Account(username)

    medias, _ = agent.get_media(account, count=1)
    follows = account.follows_count
    followers = account.followers_count
    media_count = account.media_count
    
    media_urls = []

    if medias:
        for media in medias:
            media_urls.append(media.display_url)

    user_data = {
            'username':username,
            'media_urls': media_urls,
            'follows': follows,
            'followed_by': followers,
            'media_count': media_count,
            'date': datetime.datetime.utcnow()
        }

    return user_data

def save_user_data(username, collection):
    collection.insert_one(get_user_data(username))

client = MongoClient('mongodb', 27017)
db = client['glaza']
instagram_collection = db.instagram
save_user_data(sys.argv[1], instagram_collection)
print("saved " + sys.argv[1])