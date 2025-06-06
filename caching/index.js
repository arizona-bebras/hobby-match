import { getPlayerUsername, getOwnedGames, getSteamLevel, resolveVanityUrl } from './steamApiRequests.js';
import fastify from 'fastify'
import NodeCache from 'node-cache';

const server = fastify({logger:true});

const cache = new NodeCache({stdTTL: 86400, checkperiod: 120})

server.get('/getPlayerUsername', 
    async (request, reply) => {
        const key = request.query.key;
        const id = request.query.id;
        if (!cache.has('username')) {
            cache.set('username', await getPlayerUsername(key, id));
        }
        return { username: cache.get('username') }
    }
)


server.get('/getOwnedGames',
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        if (!cache.has('games')) {
            cache.set('games', await getOwnedGames(key, id));
        }
        return { games: cache.get('games') }
    }
)

server.get('/getGameHours', 
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        const appid = request.query.appid
        if (!cache.has('games')) {
                cache.set('games', await getOwnedGames(key, id));
            }
        let app = cache.get('games').find(game => game.appid == parseInt(appid));
        if (!cache.has(`game-${app.appid}`)) {
            cache.set(`game-${app.appid}`, {
                app_name: app.name,
                hours: Math.floor(app.playtime_forever / 60)
            })
        }
        return { game: cache.get(`game-${app.appid}`)}
    }
)

server.get('/getSteamLevel', 
    async (request, reply) => {
        const key = request.query.key
        const id = request.query.id
        if (!cache.has('level')) {
            cache.set('level', await getSteamLevel(key, id));
        }
        return { level: cache.get('level') }
    }
)

server.get('/resolveVanityUrl', 
    async (request, reply) => {
        const key = request.query.key
        const vanityurl = request.query.vanityurl
        if (!cache.has('id')) {
            cache.set('id', await resolveVanityUrl(key, vanityurl));
        }
        return { 'id': cache.get('id') }
    }
)

server.listen({port: 1488});