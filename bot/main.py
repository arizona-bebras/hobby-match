from telegram import Update, WebAppInfo
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes
from tg_bot_users import BotUser
from pocketbase import PocketBase

pb = PocketBase("http://localhost:8090")

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(f'Start use app!')
    print(update.message.date)
    user_data =  {
        "tg_id": update.message.from_user.id,
        "name": update.message.from_user.first_name,
        "last_message_date": update.message.date.isoformat(),
        "first_message_date": update.message.date.isoformat()
    }
    user = BotUser(user_data, pb)
    if not user.is_user_in_db():
        user.add_user_to_db()


app = ApplicationBuilder().token('BOT_TOKEN').build()

app.add_handler(CommandHandler("start", start))

app.run_polling()