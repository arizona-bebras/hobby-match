import PocketBase from 'pocketbase';
import { PUBLIC_PB_ENDPOINT } from '$env/static/public';

export const pb = new PocketBase(PUBLIC_PB_ENDPOINT ?? 'http://127.0.0.1:8090');
