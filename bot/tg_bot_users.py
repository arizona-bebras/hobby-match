from pocketbase import PocketBase
from pocketbase.utils import ClientResponseError

class BotUser:

    def __init__(self, data: dict, pb: PocketBase):
        self.data = data
        self.pb = pb

    def add_user_to_db(self):
        
        try:
            self.pb.admins.auth_with_password("maxi.solts@gmail.com", 'DB_ADMIN_PASSWORD')

            result = self.pb.collection("tg_bot_users").create(self.data)
            print(result)
        except ClientResponseError as e:
            print(e.data, "err1")
        except Exception as e:
            print(e, "err2")

        
    def is_user_in_db(self):
        try:
            self.pb.admins.auth_with_password("maxi.solts@gmail.com", 'DB_ADMIN_PASSWORD')
            
            result = self.pb.collection('tg_bot_users').get_list(1, 1, {
                "filter": f"tg_id = '{self.data["tg_id"]}'"
            })
            
            if result.total_items > 0:
                print("Запись найдена:", result.items[0].__dict__)
                return True
            else:
                print("Запись не найдена")
                return False
                
        except ClientResponseError as e:
            print(f"API ошибка: {e.status} - {e.message}")
        except Exception as e:
            print(f"Общая ошибка: {e}")