import type { PageLoad } from '../../../../.svelte-kit/types/src/routes/(MainApp)/Search/$types';

export async function load({ parent, data }) {
  await parent();
  const { b } = data;
  console.log(b);
  return {
    a: 1,
    b,
  };
}
