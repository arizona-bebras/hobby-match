export type Widget = {
  id: string;
  order: number;
  data:
    | Audio
    | Video
    | Photos
    | Todo
    | ProgressBar
    | Geo
    | SocialMediaLink
    | SteamGame
    | Sticker
    | Survey
    | Text
    | Post;
  additionalData?: SocialMediaData | PhotoData | SurveyData;
};

export type WidgetType = Widget['data']['type'];

// SoundCloud embed
export type Audio = {
  type: 'audio';
  link: string;
};

export type Video = {
  type: 'video';
  link: string;
  platform: 'YouTube' | 'Rutube' | 'TikTok';
};

export type Photos = {
  type: 'photo';
};

export type PhotoData = {
  type: 'photo';
  urls: string[];
};

export type Task = {
  description: string;
  isCompleted: boolean;
};

export type Todo = {
  type: 'todo';
  title: string;
  tasks: Task[];
};

export type ProgressBar = {
  type: 'progress_bar';
  description: string;
  currentProgress: number;
  maxProgress: number;
};

export type Geo = {
  type: 'geo';
  id: string;
  mapLink: string;
};

export type SocialMediaLink = {
  type: 'social_media';
  platform: 'YouTube' | 'Twitch' | 'VK' | 'Steam' | 'X' | 'Telegram';
  link: string;
};

export type SocialMediaData =
  | {
      type: 'YouTube';
      title?: string;
      subscribers?: number;
    }
  | {
      type: 'Twitch';
      title?: string;
      followers?: number;
    }
  | {
      type: 'VK';
      followers?: number;
    }
  | {
      type: 'Steam';
      level?: number;
      username?: string;
    }
  | {
      type: 'X';
      followers?: number;
      tweets?: number;
    };

export type SteamGame = {
  type: 'steam_game';
  gameId: string;
  accountLink: string;
};

export type Sticker = {
  type: 'sticker';
  stickerId: string;
  cords: {
    x: number;
    y: number;
  };
};

export type Option = {
  description: string;
};

export type Survey = {
  type: 'survey';
  question: string;
  options: Option[];
};

export type SurveyData = {
  type: 'survey';
  stats: number[];
  myVote?: number;
};

export type Text = {
  type: 'text';
  text: string;
};

export type Post = {
  type: 'post';
  link: string;
};
