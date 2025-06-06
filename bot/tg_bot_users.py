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
            self.pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)

            result = self.pb.collection("users").create(self.data)
            print(result)
        except ClientResponseError as e:
            print(e.data, "err1")
        except Exception as e:
            print(e, "err2")

        
    def is_user_in_db(self):
        self.pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
        
        result = self.pb.collection('users').get_list(1, 1, {
            "filter": f"telegram_id = '{self.data['telegram_id']}'"
        })
        
        if result.total_items > 0:
            print("Запись найдена:", result.items[0].__dict__)
            return True
        else:
            print("Запись не найдена")
            return False
