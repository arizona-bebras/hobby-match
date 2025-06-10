import { AutoRouter, IRequest } from 'itty-router';

const INTERESTS_WEIGHT = 0.6;
const TEXT_WEIGHT = 0.4;

export const router = AutoRouter<IRequest, [Env, ExecutionContext]>({ base: '/feed' });

router.delete('/delete', async (request, env) => {
	const body = await request.json<{ id: string }>();
	return await env.PROFILES_VEC.deleteByIds([body.id]);
});

router.post('/upsert', async (request, env) => {
	const body = await request.json<{
		id: string;
		interest_ids: string[];
		text: string;
	}>();

	// @ts-expect-error idk types are wrong
	const embeddings: AiTextEmbeddingsOutput = await env.AI.run('@cf/baai/bge-m3', {
		text: [body.text],
	});

	const vector = embeddings.data[0];
	for (let i = 0; i < vector.length; i++) {
		vector[i] = vector[i] * TEXT_WEIGHT;
	}

	const interestsVectors = await env.TAGS_VEC.getByIds(body.interest_ids);
	for (const interestsVector of interestsVectors) {
		for (let i = 0; i < vector.length; i++) {
			vector[i] += (interestsVector.values[i] * INTERESTS_WEIGHT) / interestsVectors.length;
		}
	}

	return await env.PROFILES_VEC.upsert([
		{
			id: body.id,
			values: vector,
			metadata: {
				userid: body.id,
			},
		},
	]);
});

router.post('/query', async (request, env) => {
	const body = await request.json<{
		id: string;
		from: string[];
	}>();

	const myVec = await env.PROFILES_VEC.getByIds([body.id]);
	return await env.PROFILES_VEC.query(myVec[0].values, {
		topK: 3,
		filter: {
			userid: {
				// @ts-expect-error it works though
				$in: body.from,
			},
		},
	});
});
