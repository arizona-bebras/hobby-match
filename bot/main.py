from telegram import Update, WebAppInfo
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, MessageHandler
from tg_bot_users import BotUser
from pocketbase import PocketBase
from dotenv import load_dotenv
import os

load_dotenv('.env')

BOT_TOKEN = os.getenv('BOT_TOKEN')
pb = PocketBase("http://localhost:8090")

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(f'Start use app!')
    print(update.message.date)
    user_data =  {
        "telegram_id": update.message.from_user.id,
        "first_name": update.message.from_user.first_name,
        "first_message_date": update.message.date.isoformat(),
        "telegram_username": update.message.from_user.username,
        "password": "123",
        "passwordConfirm" : "123"
    }
    user = BotUser(user_data, pb)
    if not user.is_user_in_db():
        user.add_user_to_db()


app = ApplicationBuilder().token(BOT_TOKEN).build()

app.add_handler(CommandHandler("start", start))

app.run_polling()