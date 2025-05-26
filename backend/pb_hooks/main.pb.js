console.log('Hello World');
onAfterBootstrap((e) => {
  $app.logger().debug('Hello World');
  e.next();
});

export function getSocialMediaData(widget) {
  const platform = widget.data.platform;
  const link = widget.data.link;
  if (platform == 'Youtube') {
    return getYoutubeChannelStats(getUsernameFromUrl(link));
  }
  if (platform == 'Steam') {
    return getSteamData(link);
  }
}

export function definePlatform(url) {
  const patterns = {
    Youtube: /(youtube\.com|youtu\.be)/i,
    Tiktok: /tiktok\.com/i,
    Twitter: /(twitter\.com|x\.com)/i,
    Rutube: /rutube\.ru/i,
    VK: /(vk\.com|vkontakte\.ru)/i,
    Steam: /(steamcommunity\.com|store\.steampowered\.com)/i,
    SoundCloud: /soundcloud\.com/i,
  };

  for (const [platform, regex] of Object.entries(patterns)) {
    if (regex.test(url)) {
      return platform;
    }
  }

  return 'unknown';
}

export async function getYoutubeChannelStats(username) {
  const res = await fetch('/api/youtube', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username: username }),
  });
  const data = await res.json();
  const youtubeStats = {
    socialMediaData: data.subscribers,
  };
  console.log(youtubeStats);
  return youtubeStats;
}

export async function getSteamData(link) {
  const res = await fetch('/api/steam', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ link: link }),
  });
  const data = await res.json();
  const steamStats = {
    socialMediaData: data.player_level,
    steamUsername: data.player_name,
  };
  console.log(steamStats);
  return steamStats;
}

export function getUsernameFromUrl(url) {
  const platform = definePlatform(url);
  if (platform == 'Youtube') return url.slice(25);
  return '';
}

// eslint-disable-next-line no-undef
onRecordEnrich(
  (e) => {
    console.log('hook executed');
    e.record.withCustomData(true); // for security custom props require to be enabled explicitly
    if (e.record.data.type == 'social_media')
      e.record.set('additional_data', {
        socialMediaData: getSocialMediaData(e.record.data.link),
      });

    e.next();
  },
  ['widgets'],
);
