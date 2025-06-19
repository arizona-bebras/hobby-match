import { z } from 'zod';

const platformRegex = {
  youtube:
    /^https?:\/\/(?:www\.)?youtube\.com\/(channel|user|c|@)\/?([a-zA-Z0-9_-]+)/,
  twitch: /^https:\/\/(www\.)?twitch\.tv\/[a-zA-Z0-9_]{4,25}$/,
  vk: /^https:\/\/vk\.com\/[a-zA-Z0-9_.]+$/,
  steam:
    /^https:\/\/steamcommunity\.com\/(profiles\/[0-9]{17}|id\/[a-zA-Z0-9_-]+)\/?$/,
  twitter: /^https:\/\/(x\.com|twitter\.com)\/[a-zA-Z0-9_]{1,15}$/,
  telegram: /^https:\/\/t\.me\/[a-zA-Z0-9_]{5,32}$/,
};

const socialLinkValidation = (val: string) => {
  try {
    const url = new URL(val);

    if (url.hostname.includes('youtube.com'))
      return platformRegex.youtube.test(val);
    if (url.hostname.includes('twitch.tv'))
      return platformRegex.twitch.test(val);
    if (url.hostname.includes('vk.com')) return platformRegex.vk.test(val);
    if (url.hostname.includes('steamcommunity.com'))
      return platformRegex.steam.test(val);
    if (
      url.hostname.includes('x.com') ||
      url.hostname.includes('twitter.com')
    ) {
      return platformRegex.twitter.test(val);
    }
    if (url.hostname.includes('t.me')) return platformRegex.telegram.test(val);

    return false;
  } catch {
    return false;
  }
};

export const socialScheme = z.object({
  link: z.string().refine(socialLinkValidation, {
    message: 'Неверный формат ссылки',
  }),
});

export type FormSchema = typeof socialScheme;
