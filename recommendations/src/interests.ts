import { AutoRouter, error, IRequest } from 'itty-router';

export const router = AutoRouter<IRequest, [Env, ExecutionContext]>({ base: '/interests' });

router.delete('/delete', async (request, env) => {
	const body = await request.json<{ ids: string[] }>();
	return await env.TAGS_VEC.deleteByIds(body.ids);
});

router.post('/upsert', async (request, env) => {
	const body = await request.json<{ id: string; tag: string }>();

	// @ts-expect-error idk types are wrong
	const embeddings: AiTextEmbeddingsOutput = await env.AI.run('@cf/baai/bge-m3', {
		text: [body.tag],
	});
	return await env.TAGS_VEC.upsert([
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

	return await env.TAGS_VEC.query(queryVector.data[0], {
		topK: body.take ?? 5,
		returnMetadata: 'all',
	});
});
