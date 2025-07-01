from dotenv import load_dotenv
from telegram.constants import ParseMode

load_dotenv('.env')

from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, MessageHandler, filters
from tg_bot_users import BotUser
#from send_likes import send_likes
from pocketbase import PocketBase
import os

DB_ADDRESS = os.getenv("DB_ADDRESS")
BOT_TOKEN = os.getenv('BOT_TOKEN')
DB_ADMIN_PASSWORD = os.getenv('DB_ADMIN_PASSWORD')
DB_ADMIN_LOGIN = os.getenv('DB_ADMIN_LOGIN')

pb = PocketBase(DB_ADDRESS)
pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
app_url = os.getenv("APP_URL")

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    user_data =  {
        "telegram_id": update.message.from_user.id,
        "telegram_first_name": update.message.from_user.first_name,
        "first_message_date": update.message.date.isoformat(),
        "telegram_username": update.message.from_user.username,
        "password": "useless_password",
        "passwordConfirm" : "useless_password"
    }
    user = BotUser(user_data, pb)
    first_time = False
    if not user.is_user_in_db():
        first_time = True
        user.add_user_to_db()

    keyboard = InlineKeyboardMarkup.from_button(InlineKeyboardButton(
        text="Открыть Shumi",
        web_app=WebAppInfo(url=f"{app_url}/")))

    
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="*Привет, я Shumi\\!* 👋\n\n" +
                                        "Я помогу найти тебе новые знакомства\\.\n" +
                                        ("Заполняй анкету и вперед к поискам\\!\n" if first_time else
                                         "Заходи в приложение и находи себе друзей\\!\n"),
                                   parse_mode=ParseMode.MARKDOWN_V2,
                                   reply_markup=keyboard)
    
# async def likes(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
#     pb.admins.auth_with_password(DB_ADMIN_LOGIN, DB_ADMIN_PASSWORD)
#     telegram_id = update.message.from_user.id
#     user_db_id = pb.collection('users').get_first_list_item(f"telegram_id = '{telegram_id}'").id
#     likes = pb.collection('likes').get_full_list(batch=100, query_params={
#             "filter": f"liked_user = '{user_db_id}'",
#             "expand": "user",
#             "sort": "-created"
#         })
#     tg_ids=""
#     for like in likes:
#         tg_ids += f"@{like.expand['user'].telegram_username} " + "\n"
#     await context.bot.send_message(chat_id=update.effective_chat.id,
#                                    text=f"У тебя {len(likes)} лайков! \n" + tg_ids)

app = ApplicationBuilder().token(BOT_TOKEN).build()

app.add_handler(CommandHandler("start", start))

# app.add_handler(CommandHandler("likes", likes))

app.run_polling()
