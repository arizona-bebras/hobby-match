<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import * as Select from '$lib/components/ui/select';
  import { gameScheme } from '$lib/components/editor/steam-game/gameScheme';
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

  const form = superForm(defaults(zod4(gameScheme)), {
    SPA: true,
    validators: zod4(gameScheme),
    onSubmit: async () => {
      isLoading = true;
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
      isLoading = false;
      onClose();
    },
  });

  const { form: formData, enhance, validateForm, reset, validate } = form;

  $effect(() => {
    console.log('Сработало');
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => {
          console.log(result.data.accountLink);
          userSteamUrl = result.data.accountLink;
          $formData.gameId = result.data.gameId;
        });
    } else {
      console.log('Сработал RESET');
      reset();
    }
  });
  let isButtonActive = $state(false);
  let isSteamUrlCorrect = $state(false);
  let userSteamUrl = $state('');
  let isLoading = $state(false);

  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });

  $effect(() => {
    $formData.accountLink = userSteamUrl;
    validate('accountLink', { update: false }).then((responce) => {
      console.log('Проверка валидации!');
      isSteamUrlCorrect = responce === undefined;
      console.log(isSteamUrlCorrect);
    });
  });

  $effect(() => {
    console.log(
      'Попытка отправить запрос на получении игр пользователя',
      isSteamUrlCorrect,
    );
    if (isSteamUrlCorrect && open) {
      console.log('Запрос отправлен');
      getUserGames(userSteamUrl);
    }
  });

  let steamGames: steamGames[] = $state([]);
  async function getUserGames(accountLink: string) {
    console.log('ЗАПУСК ФУНКЦИИ getUserGames');
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
    steamGames = games.response.games.sort((a, b) =>
      a.name.localeCompare(b.name),
    );
  }
  $effect(() => {
    if (!open) {
      userSteamUrl = '';
    }
  });
</script>

<Sheet.Root
  bind:open
  onOpenChange={(state) => {
    console.log('ЗАПУСК ФУНКЦИИ!');
    if (!state) onClose();
  }}
>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Время игры"
        </p>
        <p class="pb-2 text-text-color">Введите ссылку на свой аккаунт Steam</p>
        <Form.Field {form} name="accountLink">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                {...props}
                bind:value={userSteamUrl}
                placeholder="https://steamcommunity.com/..."
                class="mb-1"
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
            <Select.Trigger
              class="w-full mt-2 text-ellipsis overflow-clip whitespace-nowrap"
              >{steamGames.length >= 1
                ? steamGames.find(
                    (game) => game.appid === parseInt($formData.gameId),
                  )?.name || 'Выберите игру'
                : 'Загрузка...'}</Select.Trigger
            >
            <Select.Content class="h-[50vh]">
              {#each steamGames as game}
                <div class="flex">
                  <img
                    src={`https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`}
                    alt="Game"
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
        {#if !!widgetId}
          <DeleteButton {widgetId} />
        {/if}
        <SaveButton
          {isLoading}
          onClick={() => {
            form.submit();
          }}
          {isButtonActive}
        />
        <!--        <SuperDebug data={$formData} />-->
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>
