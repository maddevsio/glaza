import unittest
from parser import * 

client = MongoClient('localhost', 27017)
db = client.glaza

test_instagram_users_data = db.test_instagram_users_data

class TestParser(unittest.TestCase):

    def test_parser_with_my_user(self):
        result = {'username': 'helloworld7743', 'media_urls': ['https://www.instagram.com/p/BoLhd7HBtI4', 'https://www.instagram.com/p/BoLhcYHhnLg'], 'follows': 4, 'followed_by': 0, 'media_count': 5}
        self.assertEqual(get_user_data("helloworld7743"), result)

    def test_parser_with_another_user(self):
        result = {'username': '_yanita', 'media_urls': [], 'follows': 2, 'followed_by': 5, 'media_count': 0}
        self.assertEqual(get_user_data("_yanita"), result)

    def test_inserting_to_db(self):
        test_instagram_users_data.drop()
        user_info('helloworld7743', db=test_instagram_users_data)
        self.assertEqual(test_instagram_users_data.count(), 1)

if __name__ == '__main__':
    unittest.main()