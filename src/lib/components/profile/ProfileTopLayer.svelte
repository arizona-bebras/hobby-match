<script lang="ts">
  import { pb } from '$lib/index';
  import { cn } from '$lib/utils';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import Pencil from '@lucide/svelte/icons/pencil';
  import { Plus } from '@lucide/svelte';
  let {
    data,
    shadow = true,
    container = $bindable(),
  }: { data: PageData; shadow: true } = $props();
</script>

<div
  class={cn(
    'w-full bg-background flex z-2 items-center gap-2 p-2 fixed',
    shadow ? 'shadow-2xl' : '',
  )}
  bind:this={container}
>
  <!--    src={pb.files.getURL(pb.authStore.record ?? {}, data.user_photo)}-->
  <img
    src={pb.files.getURL(pb.authStore.record ?? {}, data.user_photo)}
    class="w-14 h-14 rounded-full object-cover"
    alt="UserPhoto"
  />
  <div class="container font-[Inter] flex-1 flex-col">
    <p class="font-extrabold text-[16px]">
      {data.miniapp_name}, {data.age}
    </p>
    <p class="font-semibold text-[16px]"><span>{data.location}</span></p>
  </div>
  <div
    class="bg-accent/25 rounded-[28px] font-[Inter] flex items-center justify-center px-3 py-1"
  >
    <a
      href="/registration"
      class="text-accent-foreground font-medium flex gap-2 justify-center items-center"
    >
      <Pencil class="size-4" />
      Профиль
    </a>
  </div>
</div>
