<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import Similarity from '$lib/components/profile/Similarity.svelte';
  import RadarChart from '$lib/components/profile/RadarChart.svelte';
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import type { InterestType } from '$lib/questionnaireTypes/questionnaireTypes';
  import { Progress } from '$lib/components/ui/progress/index.js';
  let {
    userSimilarity,
    interests,
  }: { userSimilarity: number; interests: InterestType[] } = $props();
  let test = $state(25);
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
      <div class="text-accent-foreground">
        <p class="text-xl font-semibold">
          Вы похожи на {userSimilarity}%
        </p>
        <InterestsWidget {interests} />
        <p class="text-xl font-semibold">Сходство по интересам</p>
        {#each interests as interest (interest.id)}
          {@const similarity = Math.round(interest.similarity * 100)}
          <!--          <p>1234567890</p>-->
          <div class="mb-2.5">
            <div class="flex flex-row justify-between mb-1 font-semibold">
              <p>{interest.tag}</p>
              <p>
                {similarity}<span class="text-inactive text-[12px]">%</span>
              </p>
            </div>
            <Progress value={similarity} max={100} class="bg-text-color" />
          </div>
        {/each}
      </div>
    </div>
  </Sheet.Content>
</Sheet.Root>
