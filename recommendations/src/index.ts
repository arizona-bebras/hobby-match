import { AutoRouter, error, IRequest, RequestHandler } from 'itty-router';
import { router as interests } from './interests';
import { router as feed } from './feed';

type CFArgs = [Env, ExecutionContext];

const router = AutoRouter<IRequest, CFArgs>();
const withToken: RequestHandler<IRequest, CFArgs> = (request, env) => {
	const token = request.headers.get('Authorization');
	if (token !== `Bearer ${env.WORKER_SECRET}`) return error(401, 'Invalid token.');
};

router.get('/health', () => 'OK');
router.all('*', withToken).get('/', () => 'Hi!');
router.all('/interests/*', interests.fetch);
router.all('/feed/*', feed.fetch);
export default router;
