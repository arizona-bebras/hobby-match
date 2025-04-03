<script>
    let {progress, max} = $props();
</script>

<div class="GameBox flex-col">
  <p class="font-[Inter] text-[20px] font-medium">&#9876; Эйсы в Valorant</p>
  <div class="w-full bg-gray-300 rounded-lg overflow-hidden relative">
    <div class="h-8.25 bg-gradient-to-r from-green-300 to-blue-500"
         style="
         width: calc(100% * {progress} / {max});
         clip-path: polygon(0 0, 100% 0, 90% 100%, 0 100%);"
    >
    </div>
    <div class="absolute inset-0 flex items-center justify-center font-bold text-black">
      {progress} / {max}
    </div>
  </div>
</div>