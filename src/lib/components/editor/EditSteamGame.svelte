<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import * as Select from '$lib/components/ui/select/index.js';
  import { gameScheme } from '$lib/components/editor/schemes/gameScheme';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { SteamGame } from '$lib/widgetTypes/widgetTypes';
  interface steamGames {
    appid: number;
    img_icon_url: string;
    name: string;
    playtime_forever: number;
    [key: string]: any;
  }
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
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
    onSubmit: async () => {
      const widget: SteamGame = {
        type: 'steam_game',
        accountLink: $formData.accountLink,
        gameId: $formData.gameId,
      };
      if (widgetId) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
    },
  });

  const { form: formData, enhance, validateForm, reset, validate } = form;

  $effect(() => {
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => {
          $formData.accountLink = result.data.accountLink;
          $formData.gameId = result.data.gameId;
        });
    } else {
      reset();
    }
  });
  let isButtonActive = $state(false);
  let isSteamUrlCorrect = $state(false);
  let userSteamUrl = $state('');

  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  $inspect($formData.accountLink);

  $effect(() => {
    $formData.accountLink = userSteamUrl;
    validate('accountLink', { update: false }).then((responce) => {
      isSteamUrlCorrect = responce === undefined;
    });
  });

  $effect(() => {
    if (isSteamUrlCorrect && !widgetId) {
      console.log('Запрос отправлен');
      getUserGames(userSteamUrl);
    }
  });

  let steamGames: steamGames[] = $state([]);
  async function getUserGames(accountLink: string) {
    const steamRegex =
      /^(?:https:\/\/)?steamcommunity\.com\/((?:id)|(?:profiles))\/(\w+)/gm;
    const match = steamRegex.exec(accountLink);
    let steamID = match![2];

    if (match![1] === 'profiles') {
      console.log('Выполнение поиска по profile');
      steamID = match![2];
    } else if (match![1] === 'id') {
      console.log('Выполнение поиска по id');
      let result = await pb.send(`/steam/vanityurl?vanityurl=${steamID}`, {});
      steamID = result.response.id;
      console.log(steamID);
    }

    console.log('Полученный steamID:', steamID);
    const games = await pb.send(`/steam/games?id=${steamID}`, {});
    steamGames = games.response.games;
    console.log(steamGames);
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

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
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
                bind:value={userSteamUrl}
                placeholder="https://steamcommunity.com/..."
                class="mb-6.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#if isSteamUrlCorrect}
          <Select.Root
            type="single"
            bind:value={$formData.gameId}
            name="gameId"
          >
            <Select.Trigger class="w-fit mb-9"
              >{steamGames.length >= 1
                ? steamGames.find(
                    (game) => game.appid === parseInt($formData.gameId),
                  )?.name || 'Выберите игру'
                : '2'}</Select.Trigger
            >
            <Select.Content class="h-[50vh]">
              {#each steamGames as game}
                <div class="flex">
                  <img
                    src={`https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`}
                  />
                  <Select.Item value={game.appid.toString()}
                    >{game.name}</Select.Item
                  >
                </div>
              {/each}
              <!--              <Select.Item value="light">Light</Select.Item>-->
              <!--              <Select.Item value="dark">Dark</Select.Item>-->
              <!--              <Select.Item value="system">System</Select.Item>-->
            </Select.Content>
          </Select.Root>
        {/if}
        {#if widgetId !== undefined}
          <DeleteButton {widgetId} />
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
