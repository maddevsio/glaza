#!/usr/bin/python3.6

import sys
import datetime
from pymongo import MongoClient
from tgintegration import BotIntegrationClient

if len(sys.argv) !=2: 
    raise Exception("Need to pass bot name via cli args. Example: ./test.py @my_awesome_bot")

c = BotIntegrationClient(
        session_name='my_account',
        bot_under_test=sys.argv[1],
        max_wait_response=4,  # Wait a max of 4 seconds for responses, ...
        raise_no_response=False,  # ... then check for response.empty instead of raising
        min_wait_consecutive=2.0,  # Wait at least 2 seconds to collect more than one message
        global_action_delay=1.8,  # Space out all messages by 1.8 seconds
    )

c.start(False)

print("Telegram sesion is started") 

status = None
res = c.send_command_await("/start", num_expected=1)
if res.num_messages == 1:
    status = "OK"
else:
    status = "FAIL"
c.stop()

print("Bot " + sys.argv[1] + "status is " + status)

client = MongoClient('mongodb', 27017)
db = client['glaza']

data = {
    "bot_name": sys.argv[1], 
    "status": status, 
    "date": datetime.datetime.utcnow()
}

bots = db.bots
post_id = bots.insert_one(data).inserted_id
print("mongo saving status is: " + str(post_id))