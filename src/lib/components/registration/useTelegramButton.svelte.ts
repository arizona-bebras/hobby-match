export function useTelegramButton(callback: () => void) {
  $effect(() => {
    window.Telegram.WebApp.MainButton.onClick(callback);
    return () => {
      window.Telegram.WebApp.MainButton.offClick(callback);
    };
  });
}
