from dotenv import load_dotenv
load_dotenv('.env')

from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, MessageHandler, filters
from tg_bot_users import BotUser
from send_likes import send_likes
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
    if not user.is_user_in_db():
        user.add_user_to_db()

    keyboard = InlineKeyboardMarkup.from_button(InlineKeyboardButton(
        text="Shumi",
        web_app=WebAppInfo(url=f"{app_url}/")))

    
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="Приветствуем Вас в Shumi!",
                                   reply_markup=keyboard)

app = ApplicationBuilder().token(BOT_TOKEN).build()

app.add_handler(CommandHandler("start", start))

app.add_handler(CommandHandler("likes", send_likes))

app.run_polling()
