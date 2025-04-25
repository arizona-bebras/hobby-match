<script lang="ts">
  let { video } = $props();
</script>

<div class="TextBox">
<!--  <img src="alyaVideo.jpg" alt="video" class="pl-[10px]">-->
  <video controls class="w-full max-h-screen pl-[10px] rounded-[12px]">
    <source src="{video}" type="video/mp4">
  </video>
</div>