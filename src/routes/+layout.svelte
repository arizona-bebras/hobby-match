<script>
    import { pb } from '$lib/index'
    import { goto } from '$app/navigation'
    import { tick } from 'svelte';
    console.log(pb.authStore.isValid)
    if (!pb.authStore.isValid) {
        goto("/authError")
    }
</script>
<slot></slot>