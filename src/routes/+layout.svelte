<script>
    import { pb } from '$lib/index';
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    onMount ( async () => {
            if (!pb.authStore.isValid) {
                //await goto("/authError")
            }
        }
    )
</script>
<slot></slot>
