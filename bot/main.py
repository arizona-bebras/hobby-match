import os
import io
import logging
from typing import Any

from dotenv import load_dotenv
import aiohttp
from aiohttp import FormData
from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton, MenuButtonWebApp
from telegram.constants import ParseMode
from telegram.ext import (
    ApplicationBuilder,
    CommandHandler,
    ContextTypes,
    CallbackQueryHandler,
    ConversationHandler,
    MessageHandler,
    filters
)

# Загрузка переменных окружения
load_dotenv('.env')

BOT_TOKEN = os.getenv('BOT_TOKEN')
APP_URL = os.getenv("APP_URL")
API_BASE_URL = os.getenv("API_BASE_URL")
BOT_USERNAME= os.getenv("BOT_USERNAME")

# Логирование для отладки
logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    level=logging.INFO
)
logger = logging.getLogger(__name__)

MENU, DELETE = range(2)

GROUP = range(1)

# --- Класс для работы с API ---
class ApiClient:
    """Класс-обертка для запросов к твоему бэкенду"""

    @staticmethod
    async def _api_call(method: str, endpoint=None, form_data=False, **kwargs: Any) -> tuple[int, Any | None]:
        """
        Выполняет HTTP-запрос к API, обрабатывает ошибки и логирует не-200 ответы.
        Возвращает кортеж (status_code, response_data).
        """
        url = f"{API_BASE_URL}/api/tg"
        if endpoint != None:
            url += f"/{endpoint}"
        try:
            async with aiohttp.ClientSession() as session:
                headers={
                    "Authorization": f"Bearer {os.getenv("BOT_AUTH_TOKEN")}",
                    "Content-Type": "multipart/form-data"
                }
                async with session.request(method, url, headers=headers, **kwargs) as response:
                    if response.status not in [200, 201, 204]:
                        logger.warning(
                            f"API Call {method} {url} returned non-success status: {response.status}. "
                            f"Payload: {kwargs.get('json')}. Response: {await response.text()}"
                        )

                    data = None
                    if response.content_length and response.content_length > 0:
                        content_type = response.headers.get('Content-Type', '')
                        if 'application/json' in content_type:
                            data = await response.json()
                        elif 'text/plain' in content_type:
                            # Для get_user, который отдает json с content_type text/plain
                            data = await response.json(content_type='text/plain')

                    return response.status, data
        except aiohttp.ClientError as e:
            logger.error(f"Ошибка подключения к API ({method} {url}): {e}")
            return 500, None

    @staticmethod
    async def get_user(tg_id: int) -> dict | None:
        """GET запрос: проверяет наличие пользователя"""
        status, data = await ApiClient._api_call("GET", params={"id": tg_id})
        return data if status == 200 else None

    @staticmethod
    async def register_user(user_data: dict) -> int:
        """POST запрос: регистрация пользователя"""
        status, _ = await ApiClient._api_call("POST", json=user_data)
        return status

    @staticmethod
    async def update_hide_status(tg_id: int, hide: bool) -> bool:
        """PATCH запрос: обновить статус видимости"""
        status, _ = await ApiClient._api_call("PATCH", params={"id": tg_id}, json={"hide": hide})
        return status == 200

    @staticmethod
    async def delete_user(tg_id: int) -> bool:
        """DELETE запрос: удалить пользователя"""
        status, _ = await ApiClient._api_call("DELETE", params={"id": tg_id})
        return status in [200, 204]
    
    @staticmethod
    async def enter_namespace(enter_data: dict) -> int:
        status, _ = await ApiClient._api_call("POST", "namespace/add", json=enter_data)
        return status
    
    @staticmethod
    async def create_namespace(group_data: dict, photo_bytes: io.BytesIO = None) -> tuple[int, dict | None]:
        url = f"{API_BASE_URL}/api/tg/namespace/create"
        
        # Используем FormData для multipart/form-data
        data = FormData()
        
        # Добавляем текстовые поля
        for key, value in group_data.items():
            data.add_field(key, str(value))
        
        # Добавляем файл, если он есть
        if photo_bytes:
            photo_bytes.seek(0) # Сбрасываем указатель в начало
            data.add_field('picture', 
                           photo_bytes, 
                           filename='group_avatar.jpg', 
                           content_type='image/jpeg')

        try:
            async with aiohttp.ClientSession() as session:
                headers = {'Authorization': f"Bearer {os.getenv('BOT_AUTH_TOKEN')}"}
                async with session.post(url, data=data, headers=headers) as response:
                    res_data = await response.json() if response.content_length else None
                    return response.status, res_data
        except Exception as e:
            logger.error(f"API Error: {e}")
            return 500, None


# --- Хендлеры ---

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    tg_user = update.message.from_user
    tg_id = tg_user.id
    
    # Формируем данные согласно твоему ТЗ
    user_data = {
        "tg_id": str(tg_id),
        "username": tg_user.username,
        "firstname": tg_user.first_name
    }
    # 2. Регистрируем, если нет (POST)
    result = await ApiClient.register_user(user_data)
    if result == 500:
        await update.message.reply_text("Произошла ошибка при регистрации. Попробуйте позже.")
        return
    # if result == 400:
    #     await update.message.reply_text("Пользователь уже зарегестрирован")
    #     return

    if context.args:
        print(context.args)
        namespace_id = context.args[0]
        result = await ApiClient.enter_namespace({
            "user_id": str(tg_id),
            "namespace_id": namespace_id
        })
        if result != 200:
            await update.message.reply_text("Не удалось зайти в неймспейс. Попробуйте позже.")
        text = "*Привет, я Shumi\\!* 👋\n\n" + "Я помогу найти тебе новые знакомства\\.\n" + "Заполняй анкету и вперед к поискам\\!\nТы уже приглашен в неймспейс\\!\n"
    else:
        text="*Привет, я Shumi\\!* 👋\n\n" + "Я помогу найти тебе новые знакомства\\.\n" + ("Заполняй анкету и вперед к поискам\\!\n" if result == 200 else "Заходи в приложение и находи себе друзей\\!\n")

    keyboard = InlineKeyboardMarkup.from_button(InlineKeyboardButton(
        text="Открыть Shumi",
        web_app=WebAppInfo(url=f"{APP_URL}/")))

    await context.bot.set_chat_menu_button(
        chat_id=update.effective_chat.id, 
        menu_button=MenuButtonWebApp('Shumi', web_app=WebAppInfo(url=f"{APP_URL}/"))
    )
    
    await context.bot.send_message(
        chat_id=update.effective_chat.id,
        text=text,
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
    
async def offer_namespace_creation(update: Update, context: ContextTypes.DEFAULT_TYPE):
    """Общая функция для отправки предложения о создании неймспейса"""
    markup = InlineKeyboardMarkup([
        [
            InlineKeyboardButton("Создать", callback_data="create_ns"),
            InlineKeyboardButton("Пока не надо", callback_data="cancel_ns")
        ]
    ])
    
    text = "Привет! Я вижу новую группу. Создать неймспейс (общую ленту) для участников этой группы в Shumi?"
    
    if update.callback_query:
        await update.callback_query.edit_message_text(text, reply_markup=markup)
    else:
        await context.bot.send_message(
            chat_id=update.effective_chat.id,
            text=text,
            reply_markup=markup
        )
    return GROUP

async def on_bot_added(update: Update, context: ContextTypes.DEFAULT_TYPE):
    for member in update.message.new_chat_members:
        if member.id == context.bot.id:
            return await offer_namespace_creation(update, context)

async def group_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    query = update.callback_query

    if not query:
        return GROUP
    await query.answer()
    
    if query.data == "cancel_ns":
        user_id = query.from_user.id
        admins = await update.effective_chat.get_administrators()
        if not any(admin.user.id == user_id for admin in admins):
            await query.answer("❌ Только администратор может это сделать", show_alert=True)
            return GROUP
        await query.edit_message_text("👌 Понял. Если передумаете — просто тегните меня в сообщении!")
        return ConversationHandler.END # Завершаем, чтобы не висел стейт

    if query.data == "create_ns":
        chat = await context.bot.get_chat(update.effective_chat.id)
        user_id = query.from_user.id
        
        # Правильная проверка на админа
        admins = await chat.get_administrators()
        if not any(admin.user.id == user_id for admin in admins):
            await query.answer("❌ Только администратор может это сделать", show_alert=True)
            return GROUP
        
        user_data = {
            "tg_id": str(query.from_user.id),
            "username": query.from_user.username,
            "firstname": query.from_user.first_name
        }
        # 2. Регистрируем админа
        result = await ApiClient.register_user(user_data)
        if result == 500:
            await update.message.reply_text("Произошла ошибка при регистрации пользователя. Попробуйте позже.")
            return GROUP

        photo_buffer = None
        
        # 2. Получаем и скачиваем фото в память
        if chat.photo:
            try:
                tg_file = await context.bot.get_file(chat.photo.big_file_id)
                photo_buffer = io.BytesIO()
                # Скачиваем файл напрямую в буфер в оперативной памяти
                await tg_file.download_to_memory(photo_buffer)
                photo_buffer.seek(0)
            except Exception as e:
                logger.error(f"Ошибка при загрузке фото: {e}")

        # 3. Отправляем данные на бекенд
        status, res_data = await ApiClient.create_namespace({
            "title": chat.title,
            "admin_id": str(user_id),
            "description": f"Неймспейс группы {chat.title}"
        }, photo_bytes=photo_buffer)

        if status in [200, 201]:
            # Предположим, API вернул созданный ID
            ns_id = res_data.get('namespace_id', 'unknown') if res_data else "123"
            await query.edit_message_text(
                f"✅ Неймспейс создан\!\nТеперь участники могут заходить: \n[Открыть Shumi](https://t.me/{context.bot.username}?start={ns_id})",
                disable_web_page_preview=True,
                parse_mode=ParseMode.MARKDOWN_V2
            )
        else:
            await query.edit_message_text("❌ Ошибка при создании неймспейса на сервере.")
        
        return ConversationHandler.END

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

    group_conv_handler = ConversationHandler(
        entry_points=[
            MessageHandler(filters.StatusUpdate.NEW_CHAT_MEMBERS, on_bot_added),
            MessageHandler(filters.Entity("mention") & filters.ChatType.GROUPS, offer_namespace_creation)
        ],
        states={
            GROUP: [CallbackQueryHandler(group_handler, pattern="^(create_ns|cancel_ns)$")]
        },
        fallbacks=[],
        allow_reentry=True
    )

    app.add_handler(conv_handler)
    app.add_handler(group_conv_handler)
    print("Бот запущен...")
    app.run_polling()
