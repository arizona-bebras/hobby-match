export async function authMiddleware(request, reply) {
  const authHeader = request.headers['authorization'];

  if (!authHeader) {
    return reply
      .status(401)
      .send({ error: 'Missing or invalid Authorization header' });
  }

  if (authHeader !== `Bearer SECRET`) {
    return reply.status(401).send({ error: 'Invalid token' });
  }
}
