<script lang="ts">
  import { goto } from '$app/navigation';
  import { Eye, MessagesSquare, User } from '@lucide/svelte';
  import { page } from '$app/state';

  let footerButtons = [
    {
      icon: Eye,
      title: 'Просмотр',
      redirectTo: 'view',
    },
    {
      icon: MessagesSquare,
      title: 'Сообщества',
      redirectTo: 'community',
    },
    {
      icon: User,
      title: 'Моя анкета',
      redirectTo: 'profile',
    },
  ];
  let currentPage: 'view' | 'community' | 'profile' = $state('profile');

  $effect(() => {
    let path = page.url.pathname.toLowerCase();
    if (
      path.includes('view') ||
      (path.includes('/community') && path.split('/').length - 1 === 3)
    ) {
      currentPage = 'view';
    } else if (path.includes('community')) {
      currentPage = 'community';
    } else {
      currentPage = 'profile';
    }
  });
  $inspect(page.url.pathname.toLowerCase());
</script>

<footer
  class="flex justify-around text-text-color font-[Inter] text-[14px] bg-background"
>
  {#each footerButtons as button (button.title)}
    {@const Icon = button.icon}
    {@const isActive = currentPage === button.redirectTo}
    <button
      class="flex flex-col items-center py-2"
      onclick={() => goto(`/${button.redirectTo}`)}
    >
      <Icon class={isActive ? 'stroke-accent' : ''} />
      <span
        class:text-accent-foreground={isActive}
        class:font-semibold={isActive}>{button.title}</span
      >
    </button>
  {/each}
</footer>
