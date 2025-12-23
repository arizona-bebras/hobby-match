<script lang="ts">
  import { LineChart } from 'layerchart';
  import TrendingUpIcon from '@lucide/svelte/icons/trending-up';
  import { curveLinearClosed } from 'd3-shape';
  import { scaleBand } from 'd3-scale';
  import * as Chart from '$lib/components/ui/chart/index.js';
  import * as Card from '$lib/components/ui/card/index.js';
  import { Monitor } from '@lucide/svelte';
  import { userData } from '$lib/storage/userData.svelte';

  const chartData = [
    {
      category: '🥳',
      me: userData.current?.personality_test![0],
      anotherUser: 2,
    },
    {
      category: '🛠️',
      me: userData.current?.personality_test![1],
      anotherUser: 1,
    },
    {
      category: '🎨',
      me: userData.current?.personality_test![2],
      anotherUser: 3,
    },
    {
      category: '🤝',
      me: userData.current?.personality_test![4],
      anotherUser: 5,
    },
    {
      category: '🎯',
      me: userData.current?.personality_test![0],
      anotherUser: 3,
    },
  ];

  const chartConfig = {
    me: { label: 'Я' },
    anotherUser: {
      label: 'Пользователь',
    },
  } satisfies Chart.ChartConfig;
</script>

<Card.Root class="bg-transparent border-0">
  <Card.Header class="items-center">
    <!--    <Card.Title>Radar Chart - Dots</Card.Title>-->
    <!--    <Card.Description-->
    <!--      >Showing total visitors for the last 6 months</Card.Description-->
    <!--    >-->
  </Card.Header>
  <Card.Content class="px-2 flex-1">
    <Chart.Container
      config={chartConfig}
      class="mx-auto aspect-square max-h-[250px]"
    >
      <LineChart
        data={chartData}
        series={[
          {
            key: 'me',
            label: 'Desktop',
            color: 'var(--color-accent-foreground)',
            props: {
              fill: 'var(--color-accent-foreground)',
              fillOpacity: 0.5,
            },
          },
          {
            key: 'anotherUser',
            label: 'Mobile',
            color: 'var(--color-inactive)',
            props: {
              fill: 'var(--color-inactive)',
              fillOpacity: 0.5,
            },
          },
        ]}
        yDomain={[0, 10]}
        radial
        x="category"
        xScale={scaleBand()}
        points={{ r: 4 }}
        padding={12}
        props={{
          spline: {
            curve: curveLinearClosed,
            fillOpacity: 0.6,
            stroke: '0',
            motion: 'tween',
          },
          xAxis: {
            tickLength: 0,
          },
          yAxis: {
            format: () => '',
          },
          grid: {
            radialY: 'linear',
          },
          tooltip: {
            context: {
              mode: 'voronoi',
            },
          },
          highlight: {
            lines: false,
            points: false,
          },
        }}
      >
        {#snippet tooltip()}
          <Chart.Tooltip />
        {/snippet}
        {#snippet legend({ visibleSeries })}
          <div
            class="absolute -bottom-14 flex justify-center gap-x-6 pb-4 flex-wrap w-[300px] left-1/2 transform -translate-x-1/2"
          >
            <span>🥳 - Общение</span>
            <span>🛠️ - Реализация</span>
            <span>🎨 - Структура</span>
            <span>🤝 - Взаимодействие</span>
            <span>🎯 - Концентрация</span>
          </div>
        {/snippet}
      </LineChart>
    </Chart.Container>
  </Card.Content>
  <Card.Footer class="flex-col gap-2 text-sm">
    <!--    <div class="flex items-center gap-2 leading-none font-medium">-->
    <!--      Trending up by 5.2% this month <TrendingUpIcon class="size-4" />-->
    <!--    </div>-->
    <!--    <div class="text-muted-foreground flex items-center gap-2 leading-none">-->
    <!--      January - June 2024-->
    <!--    </div>-->
  </Card.Footer>
</Card.Root>
