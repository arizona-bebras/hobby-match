<script>
    let { children } = $props();
    import "../app.css";

</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
  <div class="flex flex-col bg-white flex-1 items-center text-black">
    {@render children()}
  </div>

  <footer class="flex justify-between w-full">
    <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Profile">Profile</a></button>
    <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Search">Search</a></button>
    <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Settings">Settings</a></button>
  </footer>
</div>
