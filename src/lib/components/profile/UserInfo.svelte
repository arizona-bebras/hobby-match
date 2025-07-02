<script lang="ts">
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import { pb } from '$lib';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { Flag } from '@lucide/svelte';
  import { toast } from 'svelte-sonner';

  let { data, changeMode = false }: { data: PageData; changeMode?: boolean } =
    $props();
  // let userDescription: Text = {
  //   type: 'text',
  //   text: pb.authStore.record!.user_info,
  // };
</script>

{#if !changeMode}
  <p class="font-extrabold text-[32px] mt-2 wrap-anywhere">
    <span>{data.miniapp_name}, {data.age}</span>
    {#if pb.authStore.record?.id !== data.id}
      <button
        class="text-destructive px-3 opacity-60"
        onclick={async () => {
          try {
            await pb.collection('reports').create({
              reporter: pb.authStore.record?.id,
              offender: data.id,
            });
            toast.success('Жалоба отправлена. Спасибо!');
          } catch {
            toast.warning('Жалоба уже отправлена');
          }
        }}
      >
        <Flag />
      </button>
    {/if}
  </p>
  <p class="font-semibold text-[20px] break-words">
    <span>{data.location}</span>
  </p>
{/if}
<!--<InterestsWidget interests={data.interests.expand?.interests} />-->
<InterestsWidget interests={data.interests} />
<div class="TextBox text-container my-2">
  <b>О себе: </b>
  <p class="break-words italic">{data.user_info}</p>
</div>
