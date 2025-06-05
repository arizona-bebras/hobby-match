<script lang="ts">
    let { data } = $props();
    import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
</script>

<p class="font-extrabold text-[32px] flex items-center">
    {data.miniapp_name}, {data.age}
</p>
<p class="font-semibold text-[20px]"><span>{data.location}</span></p>
<InterestsWidget
    interests={ data.interests }
/>