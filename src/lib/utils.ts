import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import client from '$lib/api/client';
import { onMount } from 'svelte';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function getCorrectForm(number: number, words_arr: string[] | string) {
  number = Math.abs(number);
  if (Number.isInteger(number)) {
    const options = [2, 0, 1, 1, 1, 2];
    return words_arr[
      number % 100 > 4 && number % 100 < 20
        ? 2
        : options[number % 10 < 5 ? number % 10 : 5]
    ];
  }
  return words_arr[1];
}

export async function getImage(
  object: 'users' | 'namespaces' | 'widgets',
  id: string,
) {
  const image = await client.GET('/api/files/{object}', {
    params: {
      path: {
        object: object,
      },
      ...(object === 'users' ? {} : { query: { id } }),
    },
  });
  if (image.data?.files) {
    return `data:image/png;base64,${image.data?.files}`;
  } else {
    return '';
  }
}
