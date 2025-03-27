import type { Actions } from './$types';
import type { PageLoad } from './$types';
import {STEAM_API_KEY as key} from '$env/static/private';

export const actions = {
    Settings: async (event) => {
        const data = await event.request.formData();
        const link = data.get('link')
        console.log('Форма для поиска игры принята')
        const gameList = await someFunc(link)
        console.log(gameList);
        event.locals.user = JSON.stringify(gameList);
    },
    Aboba: async (event) => {
        console.log("Форма \"Aboba\" принята")
        const data = await event.request.formData();
    }

} satisfies Actions;

export const load: PageLoad = async (event) => {
    return {
        games: event.locals.user
    };
};

async function someFunc(url: string): Promise<void> {
    const userProfileURL = url
    let steamID;
    const name = userProfileURL.split('/')[4]

    if (isFinite(name)){
        steamID = name
    }
    else{
        steamID = await getSteamID(key, name)
        steamID = steamID.steamid
    }

    const games = await getGames(key, steamID)
    const gameList = [];
    for (const game of games) {
        game.img_icon_url = `https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`
        // console.log(game.appid, game.name, Math.trunc(game.playtime_forever/60));
        gameList.push(game)
    }
    return gameList
}


async function getSteamID(key: string, name: string) {
    const response = await fetch(`http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${name}`);
    const json = await response.json();
    return json.response;
}





async function getGames(key: string, steamID: string) {
    const response = await fetch(`https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&steamid=${steamID}&include_appinfo=true&include_played_free_games=true&format=json`);
    const json = await response.json();
    return json.response.games;
}



// console.log(/steamcommunity\.com\/id\/(.+)\//.exec('https://steamcommunity.com/id/xrystikonelove/dasdasda')[1])