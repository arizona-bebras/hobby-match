<script lang="ts">
  import CalendarIcon from '@lucide/svelte/icons/calendar';
  import {
    type DateValue,
    DateFormatter,
    getLocalTimeZone,
    parseAbsolute,
    today,
    CalendarDate,
  } from '@internationalized/date';
  import { cn } from '$lib/utils.js';
  import { Button } from '$lib/components/ui/button/index.js';
  import { Calendar } from '$lib/components/ui/calendar/index.js';
  import * as Popover from '$lib/components/ui/popover/index.js';

  const df = new DateFormatter('ru', {
    dateStyle: 'short',
  });
  let {
    value = $bindable(),
    minValue,
    maxValue,
  }: {
    value: string;
    minValue?: DateValue | CalendarDate;
    maxValue?: DateValue | CalendarDate;
  } = $props();
  let dateValue = {
    get current(): DateValue {
      return value ? parseAbsolute(value, 'UTC') : today('UTC');
    },
    set current(val: DateValue) {
      if (!val) return;
      value = val.toDate('UTC').toISOString();
    },
  };
</script>

<Popover.Root>
  <Popover.Trigger>
    {#snippet child({ props })}
      <Button
        variant="outline"
        class={cn(
          'w-full justify-start text-left font-normal',
          !value && 'text-muted-foreground',
        )}
        {...props}
      >
        <CalendarIcon class="mr-2 size-4" />
        {value ? df.format(new Date(value)) : 'Выбери дату'}
      </Button>
    {/snippet}
  </Popover.Trigger>
  <Popover.Content class="w-auto p-0">
    <Calendar
      bind:value={dateValue.current}
      type="single"
      initialFocus
      {minValue}
      {maxValue}
    />
  </Popover.Content>
</Popover.Root>
