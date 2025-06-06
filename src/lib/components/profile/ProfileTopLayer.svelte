<script lang="ts">
    import { pb } from '$lib/index'
    let { data } = $props();
</script>

<div class="w-full h-16 bg-background fixed flex z-2">
    <img src={pb.files.getURL(pb.authStore.record!, data.user_photo)} class="w-16 h-16 rounded-full p-2" alt="UserPhoto" />
    <div class="container font-[Inter] p-[4px] w-fit">
      <p class="font-extrabold text-[16px] flex items-center">
        {data.miniapp_name}, {data.age}
      </p>
      <p class="font-semibold text-[16px]"><span>{data.location}</span></p>
    </div>
    <div
      class="w-24.5 h-8.5 bg-accent/25 rounded-[28px] font-[Inter] my-auto ml-auto flex items-center justify-center"
    >
      <a href="/registration" class="text-accent-foreground font-medium"
        >Изменить</a
      >
    </div>
</div>