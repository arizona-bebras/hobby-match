<script lang="ts">
  let { link, platform } = $props();
</script>

<div class="TextBox">
  <!--  <img src="alyaVideo.jpg" alt="video" class="pl-[10px]">-->
  <video controls class="w-full max-h-screen pl-[10px] rounded-[12px]">
    <source src={link} type="video/mp4" />
  </video>
</div>
