#!/usr/bin/python3
# thanks to https://github.com/ArtemyMagarin/tripadvisor_parser/blob/master/tripparse.py
import os
import sys
import re
import pickle
import requests
import pprint
import datetime
from pymongo import MongoClient
from bs4 import BeautifulSoup 

s = requests.session()

def getIdFromSoup(elem):
    return elem['data-reviewid']

def getIndexedUrl(url, index):
    return url.replace('Reviews-', 'Reviews-or{}-'.format(index))

def soupToDate(elem, mask="YYYY-MM-DD"):
    # YYYY-MM-DD
    text = elem['title']
    if not text:
        text = elem.text
    if not text:
        return ''

    year = re.findall(r'\d{4}', text)
    if year:
        year = year[0]
    else:
        return ''

    day = '0' + str(re.findall(r'\d{1,2}', text)[0])
    
    if "January" in text:
        month = '01'
    if "February" in text:
        month = '02'
    if "March" in text:
        month = '03'
    if "April" in text:
        month = '04'
    if "May" in text:
        month = '05'
    if "June" in text:
        month = '06'
    if "July" in text:
        month = '07'
    if "August" in text:
        month = '08'
    if "September" in text:
        month = '09'
    if "October" in text:
        month = '10'
    if "November" in text:
        month = '11'
    if "December" in text:
        month = '12'

    return mask.replace("YYYY", year).replace("MM", month).replace("DD", day[-2:])

def soupToRank(elem):
    return int(int(re.findall(r'\d{2}', elem['class'][1])[0])/10)

def getReviewIds(url):
    
    ids = []

    r = s.get(url)
    PUID = r.text[-38:-12].strip()

    soup = BeautifulSoup(r.text, "html.parser")

    ids += [getIdFromSoup(elem) for elem in soup.findAll('div', {'class': 'review-container'})] 

    try:
        lastIndex = int(soup.find('a', {'class': 'last'})['data-offset'])
    except Exception:
        return ids


    s.headers["user-agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"
    s.headers["authority"] = "www.tripadvisor.ru"
    s.headers["method"] = "POST"
    s.headers["scheme"] = "https"
    s.headers["x-puid"] = PUID
    s.headers['x-requested-with'] = 'XMLHttpRequest'

    for index in range(10, lastIndex+1, 10):
        s.headers["path"] = getIndexedUrl(url, index).replace('https://www.tripadvisor.com', '')

        r = s.post(getIndexedUrl(url, index), data={
                'reqNum': '1',
                'isLastPoll': 'false',
                'paramSeqId': '0',
                'changeSet': 'REVIEW_LIST',
                'puid': PUID,
            })
        soup = BeautifulSoup(r.text, "html.parser")
        ids += [getIdFromSoup(elem) for elem in soup.findAll('div', {'class': 'review-container'})]

    return ids


def getReviewsByIds(ids):
    data = {
        'reviews': ','.join(ids),
        'widgetChoice': 'EXPANDED_REVIEW_RESPONSIVE'
    }

    r = s.post("https://www.tripadvisor.com/OverlayWidgetAjax?Mode=EXPANDED_HOTEL_REVIEWS&metaReferer=Attraction_Review", data=data)
    
    # print(r.text)
    # sys.exit()

    soup = BeautifulSoup(r.text, "html.parser")

    texts, dates, rates = [], [], []
    texts += [elem.text for elem in soup.findAll('p', {'class': 'partial_entry'})]
    dates += [soupToDate(elem) for elem in soup.findAll('span', {'class': 'ratingDate'})]
    rates += [soupToRank(elem) for elem in soup.findAll('span', {'class': 'ui_bubble_rating'})]

    # IDK why but dates and rates duplicated
    del dates[1::2]
    del rates[1::2]

    result = []

    for d, r, t in zip(dates, rates, texts):
        result.append({
                'date': d,
                'rating': r,
                'text': t
            })

    return result

URL = sys.argv[2]
trip_advisor_reviews = getReviewsByIds(getReviewIds(URL))
pp = pprint.PrettyPrinter(indent=4)
#pp.pprint(trip_advisor_reviews)

db = MongoClient('mongodb', 27017)['glaza']

trip_advisor_reviews = {
    "attraction": sys.argv[1],
    "reviews": trip_advisor_reviews,
    "date": datetime.datetime.utcnow()
}

post_id = db.tripadvisor.insert_one(trip_advisor_reviews).inserted_id
print("mongo saving status is: " + str(post_id))