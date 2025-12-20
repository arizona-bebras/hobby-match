<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import Similarity from '$lib/components/profile/Similarity.svelte';
  import RadarChart from '$lib/components/profile/RadarChart.svelte';
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import type { InterestType } from '$lib/questionnaireTypes/questionnaireTypes';
  let {
    userSimilarity,
    interests,
  }: { userSimilarity: number; interests: InterestType[] } = $props();
</script>

<Sheet.Root>
  <Sheet.Trigger class="w-full"><Similarity {userSimilarity} /></Sheet.Trigger>
  <Sheet.Content side="bottom" class="max-h-[calc(100vh-75px)]">
    <Sheet.Header>
      <Sheet.Title>Результаты личностного теста</Sheet.Title>
      <!--      <Sheet.Description>-->
      <!--        This action cannot be undone. This will permanently delete your account-->
      <!--        and remove your data from our servers.-->
      <!--      </Sheet.Description>-->
    </Sheet.Header>
    <div class="overflow-auto h-[calc(100vh-125px)]">
      <RadarChart />
      <div>
        <p class="text-accent-foreground text-xl font-semibold">
          Вы похожи на {userSimilarity}%
        </p>
        <InterestsWidget {interests} />
      </div>
    </div>
  </Sheet.Content>
</Sheet.Root>
