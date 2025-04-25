export type Widget = {
    telegram_id: string,
    order: number,
    type: "audio" | "video" | "photo" | "todo" | "progress_bar" | "geo" | "socail_media" | "steam_game" | "sticker" | "survey" | "text"
    data : Audio | Video | Photo | Todo | ProgressBar | Geo | SocialMediaLink | SteamGame | Sticker | Survey | Text | {}
}

export type Audio = {
    link: string
};

export type Video = {
    link: string,
    platform: "YouTube" | "VK" | "Undefined";
};

export type Photo = {
    file_id: string
    width: number,
    height: number,
    size: number
};

type Task = {
    description: string,
    isCompleted: boolean
}

export type Todo =  {
    title: string,
    tasks: Task[]
};

export type ProgressBar =  { 
    description: string,
    currentProgress: number,
    maxProgress: number
};

export type Geo =   {
    id: string,
    mapLink: string
};

export type SocialMediaLink = {
    platform: "YouTube" | "Twitch" | "VK" | "Steam" | "Twitter" | "Undefined",
    username: string,
    link: string
};

export type SteamGame = {
    game: string,
    steam_user_id: string
};

export type Sticker = {
    sticker_id: string,
    cords: {
        x: number,
        y: number
    }
};

type Option = {
    description: string,
    votes: number
}

export type Survey =  {
    question: string,
    options: Option[],
    SummuryVotes: number
};

export type Text =  {
    text: string
};

