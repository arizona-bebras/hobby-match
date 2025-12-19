import PocketBase from 'pocketbase';
import { PUBLIC_PB_ENDPOINT } from '$env/static/public';

export const pb = new PocketBase(PUBLIC_PB_ENDPOINT ?? 'http://localhost:8090');
export const db = 'http://localhost:8090';
