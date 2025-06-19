from dotenv import load_dotenv
load_dotenv('.env')

from telegram import Update, Bot
from telegram.ext import ContextTypes
from pocketbase import PocketBase
import os
import asyncio
pb = PocketBase("DB_ADDRESS")
DB_ADMIN_LOGIN = os.getenv("DB_ADMIN_LOGIN")
DB_ADMIN_PASSWORD = os.getenv("DB_ADMIN_PASSWORD")
bot = Bot(token=os.getenv('BOT_TOKEN'))

async def send_likes() -> None:
    pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
    likes = pb.collection('likes').get_full_list(query_params={
            "expand": "user,liked_user",
            "sort": "-created"
        })
    for like in likes:
        if (not like.sent):
            tg_username = f"@{like.expand['user'].telegram_username}"
            liked_user_tg_id = f"{like.expand['liked_user'].telegram_id}"
            print(liked_user_tg_id)
            await bot.send_message(chat_id=liked_user_tg_id,
                                        text=f"Вы кому-то понравились! \n" + tg_username)
            pb.collection('likes').update(like.id, {
                "sent": True,
            })
    
asyncio.run(send_likes())