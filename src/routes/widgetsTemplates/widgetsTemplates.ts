export type BasicWidget = {
    telegram_id: string,
    widget_position: number,
    data : Audio | Video | Todo | ProgressBar | Geo | SocialMediaLink | SteamGame | Sticker | Survey | Text
}

export type Audio = {
        description: string,
        name: string,
        artist: string,
        duration: number,
        play_count: number,
        link: string
};

export type Video = {
    link: string,
    platform: string
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
    cords: string
};

export type SocialMediaLink = {
    platform: string,
    link: string
};

export type SteamGame =    {
    game: string,
    steam_user_id: string
};

export type Sticker =   {
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
    options: Option[],
    SummuryVotes: number
};

export type Text =  {
    text: string
};