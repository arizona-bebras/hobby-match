<script lang="ts">
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import { pb } from '$lib';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import ReportButton from '$lib/components/profile/ReportButton.svelte';
  import Zodiac from '$lib/components/profile/Zodiac.svelte';

  let {
    data,
    changeMode = false,
    isNamespaceProfile = false,
  }: {
    data: PageData;
    changeMode?: boolean;
    isNamespaceProfile?: boolean;
  } = $props();

  // let userDescription: Text = {
  //   type: 'text',
  //   text: pb.authStore.record!.user_info,
  // };
</script>

{#if !changeMode}
  {#if !isNamespaceProfile}
    <div class="font-extrabold text-[32px] mt-2 wrap-anywhere flex gap-1">
      <span>{data.miniapp_name}, {data.age}</span>
      <!--    <Zodiac />-->
      {#if pb.authStore.record?.id !== data.id}
        <ReportButton offender={data.id} />
      {/if}
    </div>
    <p class="font-semibold text-[20px] break-words">
      <span>{data.location}</span>
    </p>
  {:else}
    <div class="flex justify-between mb-2">
      <div class="flex items-center gap-3">
        <img
          src="https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg"
          alt="Изображение пользователя"
          class="size-16 rounded-full object-cover"
        />
        <div class="flex flex-col">
          <p class="font-semibold">{data.miniapp_name}, {data.age}</p>
          <p class="text-inactive">{data.location}</p>
        </div>
      </div>
      {#if pb.authStore.record?.id !== data.id}
        <ReportButton offender={data.id} />
      {/if}
    </div>
    <p>{data.user_info}</p>
  {/if}
{/if}

<InterestsWidget interests={data.interests} />
{#if !isNamespaceProfile}
  <div class="TextBox text-container my-2">
    <b>О себе: </b>
    <p class="break-words italic">{data.user_info}</p>
  </div>
{/if}
