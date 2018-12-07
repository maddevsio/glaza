from instaparser.agents import Agent
from instaparser.entities import Account, Media
from pymongo import MongoClient

client = MongoClient('localhost', 27017)
db = client.glaza

instagram_users_data = db.instagram_users_data

def get_user_data(username):
    agent = Agent()
    account = Account(username)

    media, pointer = agent.get_media(account, count=3)
    follows = account.follows_count
    followers = account.followers_count
    media_count = account.media_count
    media_urls = []
    if media:
        for i in range(2):
            media_urls.append("https://www.instagram.com/p/" + str(media[i]))

    user_data = {
            'username':username,
            'media_urls': media_urls,
            'follows': follows,
            'followed_by': followers,
            'media_count': media_count
        }

    return user_data

def user_info(username, db=instagram_users_data):
    db.insert_one(get_user_data(username))