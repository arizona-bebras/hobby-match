import type {
  Audio,
  Geo,
  PhotoData,
  Photos,
  Post,
  ProgressBar,
  SocialMediaData,
  SocialMediaLink,
  SteamGame,
  SteamGameData,
  Sticker,
  Survey,
  SurveyData,
  Text,
  Todo,
  Video,
} from '$lib/widgetTypes/widgetTypes';
import type { components } from '$lib/api/types';

// export type PageData = {
//   user_photo: string;
//   miniapp_name: string;
//   age: number;
//   location: string;
//   interests: { expand: { interests: object[] } };
//   widgets: WidgetWithService[];
//   textForm: any;
// };

export type InterestType = {
  id: string;
  tag: string;
  similarity?: number;
};

type WidgetType = {
  additionalData?: SocialMediaData | PhotoData | SurveyData | SteamGameData;
  collectionId: string;
  collectionName: string;
  created: string;
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
  files: string[];
  id: string;
  order: number;
  updated: string;
  user: string;
};

export type PageData = Required<components['schemas']['handlers.PageData']>;
