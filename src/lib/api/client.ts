import createClient from 'openapi-fetch';
import type { paths } from './types';
import { authMiddleware } from '$lib/api/authMiddleware';

const client = createClient<paths>({ baseUrl: 'http://localhost:8080/' });
client.use(authMiddleware);
export default client;
