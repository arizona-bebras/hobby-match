from typing import Any, Coroutine

from dotenv import load_dotenv
from pocketbase.errors import ClientResponseError
from telegram.constants import ParseMode

load_dotenv('.env')

from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton, MenuButtonWebApp
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, CallbackQueryHandler, ConversationHandler
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

MENU, DELETE = range(2)

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

    await context.bot.set_chat_menu_button(chat_id=update.effective_chat.id, menu_button=MenuButtonWebApp('Shumi', web_app=WebAppInfo(url=f"{app_url}/")))
    
    await context.bot.send_message(chat_id=update.effective_chat.id,
                                   text="*Привет, я Shumi\\!* 👋\n\n" +
                                        "Я помогу найти тебе новые знакомства\\.\n" +
                                        ("Заполняй анкету и вперед к поискам\\!\n" if first_time else
                                         "Заходи в приложение и находи себе друзей\\!\n"),
                                   parse_mode=ParseMode.MARKDOWN_V2,
                                   reply_markup=keyboard)
    
async def menu(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    try:
        user = pb.collection('users').get_first_list_item(f"telegram_id = '{update.effective_chat.id}'")
        keyboard = InlineKeyboardMarkup(inline_keyboard = [
            [
                InlineKeyboardButton(
                    text = "👀 Скрыть анкету" if not user.hide else "👀 Показывать анкету",
                    callback_data="hide" if not user.hide else "reveal"),
                InlineKeyboardButton(
                    text="🗑 Удалить анкету",
                    callback_data="delete")
            ]
        ])
        text = ("⚙️ *Параметры*\n\n" +
            "Ты можешь _временно_ скрыть свою анкету из поиска, " +
            "она не будет отображаться у других\\. " +
            "Анкета активируется автоматически, когда ты снова зайдешь в ленту\\.\n\n" +
            "Удаление анкеты уничтожает все твои данные *безвозвратно*\\. " +
            "После этого нужно будет заново заполнять анкету\\.")
        if update.callback_query and update.callback_query.message:
            await update.callback_query.edit_message_text(text,
                                          parse_mode=ParseMode.MARKDOWN_V2,
                                          reply_markup=keyboard)
        else:
            await context.bot.send_message(chat_id=update.effective_chat.id,
                                           text=text,
                                           parse_mode=ParseMode.MARKDOWN_V2,
                                           reply_markup=keyboard)
    except ClientResponseError:
        await start(update, context)
    return MENU
    
async def menu_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query
    if query.data == "delete":
        delete_keyboard = InlineKeyboardMarkup(inline_keyboard=[
            [
                InlineKeyboardButton(
                    text=f"Да",
                    callback_data="Yes"),
                InlineKeyboardButton(
                    text="Нет",
                    callback_data="No")
            ]
        ])

        await query.edit_message_text('🚫 *Точно удалить?*\n\n' +
                                      'Это действие никак не отменить\\! Ты потеряешь все данные анкеты и виджеты\\.',
                                      parse_mode=ParseMode.MARKDOWN_V2,
                                      reply_markup=delete_keyboard)
        await query.answer()
        return DELETE
    elif query.data == "hide" or query.data == "reveal":
        user = pb.collection('users').get_first_list_item(f"telegram_id = '{query.from_user.id}'")
        pb.collection('users').update(user.id,{
            "hide": query.data == "hide"
        })
        await query.edit_message_text("✅ *Анкета скрыта*\n\nЖдем тебя ещё\\!" if query.data == "hide" else "✅ Твою анкету снова видно",
                                      parse_mode=ParseMode.MARKDOWN_V2,
                                      reply_markup=InlineKeyboardMarkup(inline_keyboard=[
                                          [
                                              InlineKeyboardButton(
                                                  text = "👀 Показывать" if query.data == "hide" else "👀 Скрыть",
                                                  callback_data="reveal" if query.data == "hide" else "hide"),
                                              InlineKeyboardButton(
                                                  text = "⬅️ К меню",
                                                  callback_data="menu"),
                                          ]
                                      ]))
        await query.answer()
        return MENU
    elif query.data == "menu":
        await menu(update, context)
        await query.answer()
        return MENU

# async def delete(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:

#     keyboard = InlineKeyboardMarkup(inline_keyboard = [
#             [
#                 InlineKeyboardButton(
#                     text=f"Да",
#                     callback_data="Yes"),
#                 InlineKeyboardButton(
#                     text="Нет",
#                     callback_data="No")
#             ]
#         ])
    
#     await context.bot.send_message(chat_id=update.effective_chat.id,
#                                    text="Вы точно хотите удалить Вашу анкету?",
#                                    parse_mode=ParseMode.MARKDOWN_V2,
#                                    reply_markup=keyboard)
    
async def delete_button_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query
    telegram_id = query.from_user.id

    if query.data == "Yes":
        user = pb.collection('users').get_first_list_item(f"telegram_id = '{telegram_id}'")
        pb.collection('users').delete(user.id)
        await query.edit_message_text('✅ *Анкета удалена*\n\nНадеюсь, ещё увидимся\\!',
                                      parse_mode=ParseMode.MARKDOWN_V2,
                                      reply_markup=InlineKeyboardMarkup(inline_keyboard=[]))
        await query.answer()
        return MENU
    else:
        await query.edit_message_text('👌 Не удаляем',
                                      reply_markup=InlineKeyboardMarkup(inline_keyboard=[
                                            [
                                                InlineKeyboardButton(
                                                    text = "⬅️ К меню",
                                                    callback_data="menu"),
                                            ]
                                        ]))
        await query.answer()
        return MENU

    
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

conv_handler = ConversationHandler(
    entry_points=[CommandHandler('start', start), CommandHandler('menu', menu)],
    states={
        MENU: [CallbackQueryHandler(menu_handler)],
        DELETE: [CallbackQueryHandler(delete_button_handler)]
    },
    fallbacks=[CommandHandler('start', start), CommandHandler('menu', menu)]
)

# app.add_handler(CommandHandler("start", start))

# #app.add_handler(CommandHandler("delete", delete))
# app.add_handler(CommandHandler("menu", menu))

# app.add_handler(CallbackQueryHandler(menu_handler))

# app.add_handler(CallbackQueryHandler(delete_button_handler))

app.add_handler(conv_handler)

app.run_polling()
