from dotenv import load_dotenv
from pocketbase.errors import ClientResponseError
from telegram.constants import ParseMode
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
        "filter": "sent=false",
        "sort": "-created"
    })
    for like in likes:
        gender = like.expand['user'].gender
        name = (like.expand['user'].miniapp_name
                .replace('_', '\\_')
                .replace('*', '\\*')
                .replace('[', '\\[')
                .replace(']', '\\]')
                .replace('(', '\\(')
                .replace(')', '\\)')
                .replace('~', '\\~')
                .replace('`', '\\`')
                .replace('>', '\\>')
                .replace('#', '\\#')
                .replace('+', '\\+')
                .replace('-', '\\-')
                .replace('=', '\\=')
                .replace('|', '\\|')
                .replace('{', '\\{')
                .replace('}', '\\}')
                .replace('.', '\\.')
                .replace('!', '\\!')
                .replace('\\', '\\\\'))
        is_mutual = False
        try:
            pb.collection('likes').get_first_list_item(f'user="{like.liked_user}" && liked_user="{like.user}"')
            is_mutual = True
        except ClientResponseError:
            pass

        gender_postfix = 'а' if gender == 'female' else ''

        profile_url = f"tg://user?id={like.expand['user'].telegram_id}"
        name_link = f"_*[{name}]({profile_url})*_"
        if is_mutual:
            text = ("*Взаимный лайк\\!* 💖\n\n" +
                    f"{name_link} поставил{gender_postfix} тебе лайк в ответ\\.\n" +
                    "Приятного общения\\!")
        else:
            text = ("*Твоя анкета кому\\-то понравилась\\!* ❤️\n\n" +
                    f"{name_link} поставил{gender_postfix} тебе лайк\\.")

        keyboard = InlineKeyboardMarkup(inline_keyboard = [
            [
                InlineKeyboardButton(
                    text=f"Написать",
                    url=profile_url)
            ],
            [
                InlineKeyboardButton(
                    text="Посмотреть анкету",
                    web_app=WebAppInfo(url=f"{app_url}/Search?opened={like.user}"))
            ]
        ])
        try:
            photo_url = pb.collection('users').get_file_url(like.expand['user'],
                                                            like.expand['user'].user_photo,
                                                            {"thumb": "350x350"})

            await bot.send_photo(chat_id=like.expand['liked_user'].telegram_id,
                                 photo=photo_url,
                                 caption=text,
                                 reply_markup=keyboard,
                                 parse_mode=ParseMode.MARKDOWN_V2)
            pb.collection('likes').update(like.id, {
                "sent": True,
            })
        except TelegramError as e:
            print(e)
    
asyncio.run(send_likes())
