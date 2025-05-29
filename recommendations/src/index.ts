import { AutoRouter, error, IRequest, RequestHandler } from 'itty-router';

type CFArgs = [Env, ExecutionContext];

const router = AutoRouter<IRequest, CFArgs>();
const withToken: RequestHandler<IRequest, CFArgs> = (request, env) => {
	const token = request.headers.get('Authorization');
	if (token !== `Bearer ${env.WORKER_SECRET}`) return error(401, 'Invalid token.');
};

router.get('/health', () => 'OK');
router.all('*', withToken).get('/', () => 'Hi!');

router.delete('/delete', async (request, env) => {
	const body = await request.json<{ ids: string[] }>();
	return await env.VECTORIZE.deleteByIds(body.ids);
});

router.post('/upsert', async (request, env) => {
	const body = await request.json<{ id: string; tag: string }>();

	// @ts-expect-error idk types are wrong
	const embeddings: AiTextEmbeddingsOutput = await env.AI.run('@cf/baai/bge-m3', {
		text: [body.tag],
	});
	return await env.VECTORIZE.upsert([
		{
			id: body.id,
			metadata: { tag: body.tag },
			values: embeddings.data[0],
		},
	]);
});

router.post('/query', async (request, env) => {
	const body = await request.json<{ query: string; take?: number }>();

	// @ts-expect-error idk types are wrong
	const queryVector: AiTextEmbeddingsOutput = await env.AI.run('@cf/baai/bge-m3', {
		text: [body.query],
	});

	if (!('data' in queryVector)) return error(500, "Model didn't return embeddings.");

	return await env.VECTORIZE.query(queryVector.data[0], {
		topK: body.take ?? 5,
		returnMetadata: 'all',
	});
});

export default router;
