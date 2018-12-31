#!/usr/bin/python3

import json
import xmltodict
import datetime
from pymongo import MongoClient

f = open("nmap_output.xml")
xml_content = f.read()
f.close()

nmap_dict = xmltodict.parse(xml_content)
json = json.dumps(nmap_dict, sort_keys=True)

print("json size is: ",  len(json))

db = MongoClient('mongodb', 27017)['glaza']

data = {
    "host": "maddevs.io",
    "nmap_output": nmap_dict,
    "date": datetime.datetime.utcnow()
}

post_id = db.nmap.insert_one(data).inserted_id
print("mongo saving status is: " + str(post_id))