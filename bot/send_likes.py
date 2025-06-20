from dotenv import load_dotenv
from telegram.error import TelegramError

load_dotenv('.env')

from telegram import Bot, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from pocketbase import PocketBase
import os
import asyncio
pb = PocketBase(os.getenv("DB_ADDRESS"))
DB_ADMIN_LOGIN = os.getenv("DB_ADMIN_LOGIN")
DB_ADMIN_PASSWORD = os.getenv("DB_ADMIN_PASSWORD")
bot = Bot(token=os.getenv('BOT_TOKEN'))
app_url = os.getenv("APP_URL")

async def send_likes() -> None:
    pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
    likes = pb.collection('likes').get_full_list(query_params={
            "expand": "user,liked_user",
            "sort": "-created"
        })
    for like in likes:
        if not like.sent:
            tg_username = f"@{like.expand['user'].telegram_username}"
            liked_user_tg_id = like.expand['liked_user'].telegram_id
            liked_user_id = like.user
            print(liked_user_tg_id)

            keyboard = InlineKeyboardMarkup.from_button(InlineKeyboardButton(
                text="Посмотреть анкету",
                web_app=WebAppInfo(url=f"{app_url}/Search?opened={liked_user_id}")))
            try:
                await bot.send_message(chat_id=liked_user_tg_id,
                                       text=f"Вы кому-то понравились!\n{tg_username}",
                                       reply_markup=keyboard)
            except TelegramError as e:
                print(e)
            pb.collection('likes').update(like.id, {
                "sent": True,
            })
    
asyncio.run(send_likes())
