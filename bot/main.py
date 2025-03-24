from telegram import Update
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes


async def hello(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(f'Hello {update.effective_user.first_name}')


app = ApplicationBuilder().token("7580099584:AAF1g5v2UAnPcL5YZzvG4d804G3GtyighMo").build()

app.add_handler(CommandHandler("start", hello))

app.run_polling()