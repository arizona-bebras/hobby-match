<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import * as Select from '$lib/components/ui/select/index.js';
  import { gameScheme } from '$lib/components/editor/schemes/gameScheme';
  import { Trash2 } from '@lucide/svelte';
  import { pb } from '$lib';
  import { invalidate } from '$app/navigation';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  let testGames = [
    {
      appid: 10,
      name: 'Counter-Strike',
      playtime_forever: 21563,
      img_icon_url:
        'https://media.steampowered.com/steamcommunity/public/images/apps/10/6b0312cda02f5f777efa2f3318c307ff9acafbb5.jpg',
    },
    {
      appid: 80,
      name: 'Counter-Strike: Condition Zero',
      playtime_forever: 0,
      img_icon_url:
        'https://media.steampowered.com/steamcommunity/public/images/apps/80/077b050ef3e89cd84e2c5a575d78d53b54058236.jpg',
    },
  ];
  type Game = {
    appid: string;
    name: string;
    playtime_forever: string;
    img_icon_url: string;
  };
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  const form = superForm(defaults(zod(gameScheme)), {
    SPA: true,
    validators: zodClient(gameScheme),
    onSubmit: async ({ formData }) => {
      formData.set('type', 'steam_game');
      const formValues = Object.fromEntries(formData);
      console.log(formValues);
      if (widgetId != undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });

  const { form: formData, enhance, validateForm } = form;

  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId)
      .then((result) => ($formData.text = result.data.text));
  }
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  let steamGames: Game[] = $state([]);
  async function getUserGames() {
    const games = await fetch(
      '/api/steam?link=https%3A%2F%2Fsteamcommunity.com%2Fid%2Fxrystikonelove%2F',
    );
    return await games.json();
  }

  // $effect(() => {
  //   if (isButtonActive) {
  //     getUserGames().then((result) => {
  //       console.log(result);
  //       steamGames = result;
  //     });
  //   }
  // });
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Время игры"
        </p>
        <p class="pb-2">Введите ссылку на свой аккаунт Steam</p>
        <Form.Field {form} name="accountLink">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                {...props}
                bind:value={$formData.accountLink}
                placeholder="https://steamcommunity.com/id/..."
                class="mb-6.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#if isButtonActive}
          <Select.Root
            type="single"
            bind:value={$formData.gameId}
            name="gameId"
          >
            <Select.Trigger class="w-fit mb-9"
              >{testGames.length >= 1
                ? testGames.find(
                    (game) => game.appid === parseInt($formData.gameId),
                  )?.name || 'Выберите игру'
                : '2'}</Select.Trigger
            >
            <Select.Content class="h-[50vh]">
              {#each testGames as game}
                <div class="flex">
                  <img src={game.img_icon_url} />
                  <Select.Item value={game.appid.toString()}
                    >{game.name}</Select.Item
                  >
                </div>
              {/each}
              <Select.Item value="light">Light</Select.Item>
              <Select.Item value="dark">Dark</Select.Item>
              <Select.Item value="system">System</Select.Item>
            </Select.Content>
          </Select.Root>
        {/if}

        <SaveButton
          onClick={() => {
            form.submit();
            open = false;
            onOpenChange();
          }}
          {isButtonActive}
        />
        <SuperDebug data={$formData} />
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>
