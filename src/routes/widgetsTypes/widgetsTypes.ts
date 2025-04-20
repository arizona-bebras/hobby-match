export type BasicWidget = {
    telegram_id: string,
    order: number,
    data : Audio | Video | Photo | Todo | ProgressBar | Geo | SocialMediaLink | SteamGame | Sticker | Survey | Text
}

type Audio = {
        description: string,
        name: string,
        artist: string,
        duration: number,
        play_count: number,
        link: string
};

type Video = {
    link: string,
    platform: "YouTube" | "VK"
};

type Photo = {
    file_id: string
    width: number,
    height: number,
    size: number
};

type Task = {
    description: string,
    isCompleted: boolean
}

type Todo =  {
    title: string,
    tasks: Task[]
};

type ProgressBar =  { 
    description: string,
    currentProgress: number,
    maxProgress: number
};

type Geo =   {
    id: string,
    mapLink: string
};

type SocialMediaLink = {
    platform: "YouTube" | "Twitch" | "VK" | "Steam" | "Twitter",
    username: string,
    link: string
};

type SteamGame =    {
    game: string,
    steam_user_id: string
};

type Sticker =   {
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

type Survey =  {
    options: Option[],
    SummuryVotes: number
};

type Text =  {
    text: string
};