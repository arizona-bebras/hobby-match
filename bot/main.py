import os
import logging
from typing import Any

from dotenv import load_dotenv
import aiohttp
from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton, MenuButtonWebApp
from telegram.constants import ParseMode
from telegram.ext import (
    ApplicationBuilder,
    CommandHandler,
    ContextTypes,
    CallbackQueryHandler,
    ConversationHandler
)

# Загрузка переменных окружения
load_dotenv('.env')

BOT_TOKEN = os.getenv('BOT_TOKEN')
APP_URL = os.getenv("APP_URL")
API_BASE_URL = os.getenv("API_BASE_URL")

# Логирование для отладки
logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    level=logging.INFO
)
logger = logging.getLogger(__name__)

MENU, DELETE = range(2)

# --- Класс для работы с API ---
class ApiClient:
    """Класс-обертка для запросов к твоему бэкенду"""

    @staticmethod
    async def get_user(tg_id: int) -> dict | None:
        """GET запрос: проверяет наличие пользователя"""
        url = f"{API_BASE_URL}/api/tg"
        async with aiohttp.ClientSession() as session:
            try:
                # Передаем tg_id как query parameter
                async with session.get(url, params={ "id": tg_id }) as response:
                    if response.status == 200:
                        return await response.json(content_type='text/plain')
                    return None
            except Exception as e:
                logger.error(f"Ошибка подключения к API (check_user): {e}")
                return None

    @staticmethod
    async def register_user(user_data: dict) -> int:
        """POST запрос: регистрация пользователя"""
        url = f"{API_BASE_URL}/api/tg"
        async with aiohttp.ClientSession() as session:
            try:
                async with session.post(url, json=user_data) as response:
                    return response.status
            except Exception as e:
                logger.error(f"Ошибка подключения к API (register_user): {e}")
                return 500

    @staticmethod
    async def update_hide_status(tg_id: int, hide: bool) -> bool:
        """PATCH запрос: обновить статус видимости"""
        # Предполагаем, что есть эндпоинт для обновления по ID
        url = f"{API_BASE_URL}/api/tg"
        async with aiohttp.ClientSession() as session:
            try:
                async with session.patch(url, json={"hide": hide}, params={ "id": tg_id }) as response:
                    return response.status == 200
            except Exception as e:
                logger.error(f"Ошибка API (update_hide_status): {e}")
                return False

    @staticmethod
    async def delete_user(tg_id: int) -> bool:
        """DELETE запрос: удалить пользователя"""
        url = f"{API_BASE_URL}/api/tg"
        async with aiohttp.ClientSession() as session:
            try:
                async with session.delete(url, params={ "id": tg_id }) as response:
                    return response.status in [200, 204]
            except Exception as e:
                logger.error(f"Ошибка API (delete_user): {e}")
                return False

# --- Хендлеры ---

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    tg_user = update.message.from_user
    tg_id = tg_user.id
    
    # Формируем данные согласно твоему ТЗ
    register_data = {
        "tg_id": str(tg_id),
        "username": tg_user.username,
        "firstname": tg_user.first_name
    }
    # 2. Регистрируем, если нет (POST)
    result = await ApiClient.register_user(register_data)
    if result == 500:
        await update.message.reply_text("Произошла ошибка при регистрации. Попробуйте позже.")
        return
    if result == 400:
        await update.message.reply_text("Пользователь уже зарегестрирован")
        return

    keyboard = InlineKeyboardMarkup.from_button(InlineKeyboardButton(
        text="Открыть Shumi",
        web_app=WebAppInfo(url=f"{APP_URL}/")))

    await context.bot.set_chat_menu_button(
        chat_id=update.effective_chat.id, 
        menu_button=MenuButtonWebApp('Shumi', web_app=WebAppInfo(url=f"{APP_URL}/"))
    )
    
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text="*Привет, я Shumi\\!* 👋\n\n" +
             "Я помогу найти тебе новые знакомства\\.\n" +
             ("Заполняй анкету и вперед к поискам\\!\n" if result == 200 else
              "Заходи в приложение и находи себе друзей\\!\n"),
        parse_mode=ParseMode.MARKDOWN_V2,
        reply_markup=keyboard
    )

async def menu(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    tg_id = update.effective_chat.id
    
    # Получаем актуальные данные пользователя, чтобы узнать статус hide
    user_data = await ApiClient.get_user(tg_id)
    
    # Если база недоступна или юзера нет, отправляем на старт
    if not user_data:
        await update.message.reply_text("Сначала нужно зарегестрироваться👀. Напиши /start")
        return ConversationHandler.END

    # Предполагаем, что API возвращает поле 'hide' или 'is_hidden'
    is_hidden = user_data.get('hide', False) 

    keyboard = InlineKeyboardMarkup(inline_keyboard=[
        [
            InlineKeyboardButton(
                text="👀 Скрыть анкету" if not is_hidden else "👀 Показывать анкету",
                callback_data="hide" if not is_hidden else "reveal"),
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
        await update.callback_query.edit_message_text(
            text,
            parse_mode=ParseMode.MARKDOWN_V2,
            reply_markup=keyboard
        )
    else:
        await context.bot.send_message(
            chat_id=tg_id,
            text=text,
            parse_mode=ParseMode.MARKDOWN_V2,
            reply_markup=keyboard
        )
    return MENU

async def menu_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query
    tg_id = query.from_user.id
    
    if query.data == "delete":
        delete_keyboard = InlineKeyboardMarkup(inline_keyboard=[
            [
                InlineKeyboardButton(text="Да", callback_data="Yes"),
                InlineKeyboardButton(text="Нет", callback_data="No")
            ]
        ])

        await query.edit_message_text(
            '🚫 *Точно удалить?*\n\n' +
            'Это действие никак не отменить\\! Ты потеряешь все данные анкеты и виджеты\\.',
            parse_mode=ParseMode.MARKDOWN_V2,
            reply_markup=delete_keyboard
        )
        await query.answer()
        return DELETE
        
    elif query.data == "hide" or query.data == "reveal":
        should_hide = (query.data == "hide")
        
        # Обновляем статус через API
        success = await ApiClient.update_hide_status(tg_id, should_hide)
        
        if success:
            msg_text = ("✅ *Анкета скрыта*\n" +
                        "Анкета активируется автоматически, когда ты снова зайдешь в ленту\\.\n\n" +
                        "Ждем тебя ещё\\!" if should_hide else "✅ Твою анкету снова видно")
            
            # Клавиатура обновляется в зависимости от нового состояния
            reply_markup = InlineKeyboardMarkup(inline_keyboard=[
                [
                    InlineKeyboardButton(
                        text="👀 Показывать" if should_hide else "👀 Скрыть",
                        callback_data="reveal" if should_hide else "hide"),
                    InlineKeyboardButton(
                        text="⬅️ К меню",
                        callback_data="menu"),
                ]
            ])
            await query.edit_message_text(msg_text, parse_mode=ParseMode.MARKDOWN_V2, reply_markup=reply_markup)
        else:
            await query.answer("Ошибка связи с сервером", show_alert=True)
            
        await query.answer()
        return MENU
        
    elif query.data == "menu":
        await menu(update, context)
        await query.answer()
        return MENU

async def delete_button_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query
    tg_id = query.from_user.id

    if query.data == "Yes":
        # Удаляем через API
        success = await ApiClient.delete_user(tg_id)
        
        if success:
            await query.edit_message_text(
                '✅ *Анкета удалена*\n\nНадеюсь, ещё увидимся\\!',
                parse_mode=ParseMode.MARKDOWN_V2,
                reply_markup=InlineKeyboardMarkup(inline_keyboard=[])
            )
        else:
             await query.edit_message_text(
                '❌ *Ошибка удаления*\n\nПопробуйте позже\\.',
                parse_mode=ParseMode.MARKDOWN_V2,
                reply_markup=InlineKeyboardMarkup(inline_keyboard=[])
            )
            
        await query.answer()
        return MENU
    else:
        await query.edit_message_text(
            '👌 Не удаляем',
            reply_markup=InlineKeyboardMarkup(inline_keyboard=[
                [
                    InlineKeyboardButton(text="⬅️ К меню", callback_data="menu")
                ]
            ])
        )
        await query.answer()
        return MENU

if __name__ == '__main__':
    if not BOT_TOKEN:
        print("Ошибка: BOT_TOKEN не найден в .env")
        exit()

    app = ApplicationBuilder().token(BOT_TOKEN).build()

    conv_handler = ConversationHandler(
        entry_points=[CommandHandler('start', start), CommandHandler('menu', menu)],
        states={
            MENU: [CallbackQueryHandler(menu_handler)],
            DELETE: [CallbackQueryHandler(delete_button_handler)]
        },
        fallbacks=[CommandHandler('start', start), CommandHandler('menu', menu)]
    )

    app.add_handler(conv_handler)
    print("Бот запущен...")
    app.run_polling()
