export type Widget = {
  id: string;
  telegram_id: string;
  order: number;
  type:
    | 'audio'
    | 'video'
    | 'photo'
    | 'todo'
    | 'progress_bar'
    | 'geo'
    | 'social_media'
    | 'steam_game'
    | 'sticker'
    | 'survey'
    | 'text';
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
    | Text;
};

export type Audio = {
  link: string;
};

export type Video = {
  link: string;
  platform: 'YouTube' | 'Rutube' | 'TikTok' | 'Undefined';
};

export type Photo = {
  name: string;
  width: number;
  height: number;
  size: number;
};

export type Photos = {
  photos: Photo[];
};

export type Task = {
  order: number;
  description: string;
  isCompleted: boolean;
};

export type Todo = {
  title: string;
  tasks: Task[];
};

export type ProgressBar = {
  description: string;
  currentProgress: number;
  maxProgress: number;
};

export type Geo = {
  id: string;
  mapLink: string;
};

export type SocialMediaLink = {
  platform: 'YouTube' | 'Twitch' | 'VK' | 'Steam' | 'Twitter' | 'Undefined';
  username: string;
  link: string;
  subscribers: number;
};

export type SteamGame = {
  game: string;
  steam_user_id: string;
  hours_played: number;
  game_icon: string;
};

export type Sticker = {
  sticker_id: string;
  cords: {
    x: number;
    y: number;
  };
};

export type Option = {
  description: string;
  votes: number;
};

export type Survey = {
  question: string;
  options: Option[];
  summuryVotes: number;
};

export type Text = {
  text: string;
};
