<script lang="ts">
  import { pb } from '$lib/index';
  import { Input } from '$lib/components/ui/input/index.js';
  import { Plus } from '@lucide/svelte';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import SuperDebug, {
    type Infer,
    superForm,
    type SuperValidated,
  } from 'sveltekit-superforms';
  import {
    type FormSchema,
    interestsScheme,
  } from '$lib/components/registration/InterestsFormShema';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import { onDestroy } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  type Word = {
    id: string;
    tag: string;
  };
  let selectedInterests: Word[] = $state([]);
  let userInterest: string = $state('');
  let { form: interests }: { form: SuperValidated<Infer<FormSchema>> } =
    $props();

  const form = superForm(interests, {
    validators: zodClient(interestsScheme),
    dataType: 'json',
  });

  const { form: formData, enhance, validateForm } = form;

  export async function save() {
    window.Telegram.WebApp.MainButton.showProgress();
    await pb
      .collection('users')
      .update(pb.authStore.record!.id, $formData)
      .finally(window.Telegram.WebApp.MainButton.hideProgress);
  }

  async function handleTelegramButtonClick() {
    await save();
    form.submit();
  }

  useTelegramButton(handleTelegramButtonClick);
  $effect(() => {
    validateForm().then((response) => {
      if (response.valid) {
        window.Telegram.WebApp.MainButton.setParams({
          color: window.Telegram.WebApp.themeParams.button_color,
          is_active: true,
          is_visible: true,
        });
      } else {
        window.Telegram.WebApp.MainButton.setParams({
          color: '#808080',
          is_active: false,
          is_visible: true,
        });
      }
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  $effect(() => {
    pb.collection('users')
      .getOne(pb.authStore.record!.id, {
        fields: 'expand',
        expand: 'interests',
        requestKey: null,
      })
      .then((user) => {
        selectedInterests = user.expand?.interests ?? [];
        $formData.interests = selectedInterests.map((i) => i.id);
      });
  });
  // $effect(() => {
  //   $formData.interests = selectedInterests;
  // });
  onDestroy(() => {
    window.Telegram.WebApp.MainButton.hide();
  });
  let suggestedWords: Word[] = $state([]);
  $inspect(suggestedWords);
  async function getWords(userInterest: string) {
    const result = await pb.send('/worker/autocomplete', {
      method: 'GET',
      query: {
        query: userInterest,
      },
    });
    suggestedWords = result.response.matches.map(
      (element: { id: string; metadata: { tag: string } }) => ({
        id: element.id,
        tag: element.metadata.tag,
      }),
    );
    // response['response']['matches'].forEach((element) => {
    //   if (!suggestedWords.includes(element['metadata']['tag'])) {
    //     suggestedWords.push({
    //       id: element.id,
    //       tag: element['metadata']['tag'],
    //     });
    //   }
    // }),
    // console.log(result);
  }
  $effect(() => {
    const query = userInterest;
    const timeout = setTimeout(() => {
      if (query.length >= 1) {
        getWords(query);
      }
    }, 300);

    return () => clearTimeout(timeout);
  });
  $inspect(suggestedWords);
  let showingInterests = $derived([
    ...selectedInterests.map((element) => ({ ...element, selected: true })),
    ...suggestedWords
      .filter(
        (element) =>
          !selectedInterests.find((selected) => selected.id === element.id),
      )
      .map((element) => ({ ...element, selected: false })),
  ]);
</script>

<form method="POST" action="?/interests" use:enhance>
  <div class="flex flex-col gap-y-2">
    <div class="flex w-full mb-2 items-center">
      <p class="font-medium text-2xl h-15.25">
        <Emoji symbol="💗" class="size-6" />Давай узнаем друг-друга поближе?
      </p>
    </div>
    <p>Выберите наиболее интересующие вас темы</p>
    <p class="pb-1 text-gray-400">Отметьте от 3-x до 5-и тем</p>
    <div class="flex flex-col gap-y-4">
      <Input
        type="text"
        placeholder="Начните вводить"
        bind:value={userInterest}
        onkeydown={(e) => {
          if (e.key === 'Enter') {
            e.preventDefault();
          }
        }}
      />

      {#if userInterest.length >= 1 || selectedInterests.length >= 1}
        <div class="flex flex-row gap-2 w-full flex-wrap font-medium">
          {#each showingInterests as element, i}
            <button
              onclick={(e) => {
                if (!element.selected) {
                  selectedInterests.push(element);
                } else {
                  selectedInterests.splice(i, 1);
                }
                $formData.interests = selectedInterests.map(
                  (element) => element.id,
                );
                // if (selectedInterests.includes (element['id'])) {
                //   selectedInterests.splice(
                //     selectedInterests.indexOf(element['id']),
                //
                //     1,
                //   );
                // } else {
                //   selectedInterests.push(element['id']);
                // }
                // $formData.interests = selectedInterests;
                e.preventDefault();
              }}
              class="{element.selected
                ? 'bg-accent'
                : 'bg-accent/25'} rounded-3xl flex flex-row items-center justify-center px-3 py-2 gap-1.5"
            >
              <Plus
                class="{element.selected
                  ? 'text-white rotate-45'
                  : 'text-accent-foreground'} w-4 h-5 stroke-3"
              />
              <p
                class="text-accent-foreground {element.selected
                  ? 'text-white'
                  : 'text-accent-foreground'}"
              >
                {element['tag']}
              </p>
            </button>
          {/each}
        </div>
      {/if}

      <!--{#if userInterest.length >= 1}-->
      <!--  {#each suggestedWords as word}-->
      <!--    <button-->
      <!--      type="button"-->
      <!--      class="bg-green-400 text-black"-->
      <!--      onclick={() => {-->
      <!--        listOfInterests.push(word);-->
      <!--      }}>{word}</button-->
      <!--    >-->
      <!--  {/each}-->
      <!--  <p>Предложенные интересы</p>-->
      <!--{/if}-->
      <!--      <div class="flex flex-row gap-2 w-full flex-wrap font-medium">-->
      <!--        {#each listOfInterests as element}-->
      <!--          <button-->
      <!--            onclick={(e) => {-->
      <!--              if (selectedInterests.includes(element)) {-->
      <!--                selectedInterests.splice(selectedInterests.indexOf(element), 1);-->
      <!--              } else {-->
      <!--                selectedInterests.push(element);-->
      <!--              }-->
      <!--              $formData.interests = selectedInterests;-->
      <!--              e.preventDefault();-->
      <!--            }}-->
      <!--            class="{selectedInterests.includes(element)-->
      <!--              ? 'bg-accent'-->
      <!--              : 'bg-accent/25'} rounded-3xl flex flex-row items-center justify-center px-3 py-2 gap-1.5"-->
      <!--          >-->
      <!--            <Plus-->
      <!--              class="{selectedInterests.includes(element)-->
      <!--                ? 'text-white rotate-45'-->
      <!--                : 'text-accent-foreground'} w-4 h-5 stroke-3"-->
      <!--            />-->
      <!--            <p-->
      <!--              class="text-accent-foreground {selectedInterests.includes(element)-->
      <!--                ? 'text-white'-->
      <!--                : 'text-accent-foreground'}"-->
      <!--            >-->
      <!--              {element}-->
      <!--            </p>-->
      <!--          </button>-->
      <!--        {/each}-->
      <!--      </div>-->
    </div>
  </div>
  <!--  <SuperDebug data={$formData} />-->
</form>
