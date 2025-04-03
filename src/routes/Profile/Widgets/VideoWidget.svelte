<script lang="ts">
  let { video } = $props();
</script>

<div class="TextBox">
<!--  <img src="alyaVideo.jpg" alt="video" class="pl-[10px]">-->
  <video controls class="w-full h-49 pl-[10px] object-fill">
    <source src="{video}" type="video/mp4">
  </video>
</div>