<script lang="ts">
  import { Heart, HeartOff } from '@lucide/svelte';
  let { likedUserId } = $props();
  let liked = $state(false);
  $effect(() => {
    if (!liked) {
      // TODO: pocketbase was here
      // pb.collection('likes').create({
      //   user: pb.authStore.record!.id,
      //   liked_user: likedUserId
      // })
    }
    console.log('LIKEDLIKEDLIKEDLIKED');
  });
</script>

<button
  onclick={() => (liked = !liked)}
  class="bg-accent size-12.5 flex items-center justify-center rounded-xl"
>
  {#if !liked}
    <Heart class="size-6 text-text-color" />
  {:else}
    <HeartOff class="size-6 text-text-color" />
  {/if}
</button>
