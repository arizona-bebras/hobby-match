from dotenv import load_dotenv
load_dotenv('.env')

from telegram import Update
from telegram.ext import ContextTypes
from pocketbase import PocketBase
import os

async def send_likes(update: Update, 
                     context: ContextTypes.DEFAULT_TYPE, 
                     pb: PocketBase = PocketBase("http://127.0.0.1:8090"),
                     DB_ADMIN_LOGIN = os.getenv("DB_ADMIN_LOGIN"),
                     DB_ADMIN_PASSWORD = os.getenv("DB_ADMIN_PASSWORD")) -> None:
    pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
    telegram_id = update.message.from_user.id,
    user_db_id = await pb.collection('users').get_first_list_item(f"telegram_id = '{telegram_id}'").id
    likes = pb.collection('likes').get_full_list(batch=100, query_params={
            "filter": f"liked_user = '{user_db_id}'",
            "expand": "user",
            "sort": "-created"
        })
    tg_ids=""
    for like in likes:
        tg_ids += f"@{like.expand['user'].telegram_username} " + "\n"
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text=f"У Вас {len(likes)} лайков! \n" + tg_ids)