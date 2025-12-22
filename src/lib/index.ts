// import PocketBase from 'pocketbase';
import { PUBLIC_API_ENDPOINT } from '$env/static/public';
//
// export const pb = new PocketBase(PUBLIC_PB_ENDPOINT ?? 'http://localhost:8090');
export const db = PUBLIC_API_ENDPOINT ?? 'http://localhost:8090';
