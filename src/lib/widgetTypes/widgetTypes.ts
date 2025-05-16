export type Widget = {
  id: string;
  telegram_id: string;
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
    | Empty;
};

export type Audio = {
  type: 'audio';
  link: string;
};

export type Video = {
  type: 'video';
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
  type: 'photo';
  photos: Photo[];
};

export type Task = {
  order: number;
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
  platform: 'YouTube' | 'Twitch' | 'VK' | 'Steam' | 'Twitter' | 'Undefined';
  username: string;
  link: string;
};

export type SteamGame = {
  type: 'steam_game';
  game: string;
  steam_user_id: string;
  hours_played: number;
  game_icon: string;
};

export type Sticker = {
  type: 'sticker';
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
  type: 'survey';
  question: string;
  options: Option[];
  summaryVotes: number;
};

export type Text = {
  type: 'text';
  text: string;
};

export type Empty = {
  type: 'empty';
};
