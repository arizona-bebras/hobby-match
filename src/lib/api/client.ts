import createClient from 'openapi-fetch';
import type { paths } from './types';
import { authMiddleware } from '$lib/api/authMiddleware';
import { db } from '$lib';

const client = createClient<paths>({ baseUrl: db });
client.use(authMiddleware);
export default client;
