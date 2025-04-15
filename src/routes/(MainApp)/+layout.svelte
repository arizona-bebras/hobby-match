<script>
  let { children } = $props();
  import "./app.css";

</script>

<div class="w-screen h-screen flex flex-col">
{@render children()}

<footer class="flex justify-between">
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Profile">Profile</a></button>
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Search">Search</a></button>
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Settings">Settings</a></button>
</footer>
</div>
