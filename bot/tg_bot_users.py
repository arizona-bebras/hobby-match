from pocketbase import PocketBase
from pocketbase.utils import ClientResponseError
import os

DB_ADMIN_PASSWORD = os.getenv('DB_ADMIN_PASSWORD')
DB_ADMIN_LOGIN = os.getenv('DB_ADMIN_LOGIN')

class BotUser:

    def __init__(self, data: dict, pb: PocketBase):
        self.data = data
        self.pb = pb

    def add_user_to_db(self):
        try:
            result = self.pb.collection("users").create(self.data)
            print(result)
        except ClientResponseError as e:
            print(e.data, "err1")
        except Exception as e:
            print(e, "err2")

        
    def is_user_in_db(self):
        try:
          result = self.pb.collection('users').get_first_list_item(f"telegram_id = '{self.data['telegram_id']}'")
          print("Запись найдена:", result)
          return True
        except ClientResponseError:
            print("Запись не найдена")
            return False
