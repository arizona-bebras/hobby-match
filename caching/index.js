import { getPlayerUsername, getOwnedGames, getSteamLevel, resolveVanityUrl } from './steamApiRequests.js';
import fastify from 'fastify'
import NodeCache from 'node-cache';

const server = fastify({logger:true});

const cache = new NodeCache({stdTTL: 86400, checkperiod: 120})

server.get('/getPlayerUsername', 
    async (request, reply) => {
        const key = request.query.key;
        const id = request.query.id;
        if (!cache.has(`username-${id}`)) {
            cache.set(`username-${id}`, await getPlayerUsername(key, id));
        }
        return { username: cache.get(`username-${id}`) }
    }
)


server.get('/getOwnedGames',
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        if (!cache.has(`games-${id}`)) {
            const games = await getOwnedGames(key, id)
                cache.set(`games-${id}`, games);
                for (let game of games) {
                    cache.set(`game-${id}-${game.appid}`, {
                        game_name: game.name,
                        hours: Math.floor(game.playtime_forever / 60)
                    })
                }
        }
        return { games: cache.get(`games-${id}`) }
    }
)

server.get('/getGameHours', 
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        const appid = request.query.appid
        if (!cache.has(`games-${id}`)) {
                const games = await getOwnedGames(key, id)
                cache.set(`games-${id}`, games);
                for (let game of games) {
                    cache.set(`game-${id}-${game.appid}`, {
                        game_name: game.name,
                        hours: Math.floor(game.playtime_forever / 60)
                    })
                }
            }
        console.log(cache)
        return { game: cache.get(`game-${id}-${appid}`)}
    }
)

server.get('/getSteamLevel', 
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        if (!cache.has(`level-${id}`)) {
            cache.set(`level-${id}`, await getSteamLevel(key, id));
        }
        return { level: cache.get(`level-${id}`) }
    }
)

server.get('/resolveVanityUrl', 
    async (request, reply) => {
        const key = request.query.key
        const vanityurl = request.query.vanityurl
        if (!cache.has(`id-${vanityurl}`)) {
            cache.set(`id-${vanityurl}`, await resolveVanityUrl(key, vanityurl));
        }
        return { id: cache.get(`id-${vanityurl}`) }
    }
)

server.listen({port: 1488});