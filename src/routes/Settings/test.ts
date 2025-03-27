const key = 'CD5E6E3330889C0216BF7DD98CC6FA85'
const userProfileURL = 'https://steamcommunity.com/profiles/76561199001879011'
export let steamID: string;
const name = userProfileURL.split('/')[4]

async function getSteamID(key, name) {
    const response = await fetch(`http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${name}`);
    const json = await response.json();
    return json.response;
}

if (isFinite(name)){
    steamID = name
}
else{
    steamID = await getSteamID(key, name)
    steamID = steamID.steamid
}
console.log(steamID)



async function getGames(key, steamID) {
    const response = await fetch(`https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&steamid=${steamID}&include_appinfo=true&include_played_free_games=true&format=json`);
    const json = await response.json();
    return json.response.games;
}

const games = await getGames(key, steamID)
for (const game of games) {
    // console.log(game.appid, game.name, Math.trunc(game.playtime_forever/60));
    console.log(game);
}
