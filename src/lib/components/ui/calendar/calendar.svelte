<script lang="ts">
  import {
    Calendar as CalendarPrimitive,
    type WithoutChildrenOrChild,
  } from 'bits-ui';
  import {
    DateFormatter,
    getLocalTimeZone,
    today,
    type DateValue,
  } from '@internationalized/date';
  import * as Calendar from './index.js';
  import * as Select from '$lib/components/ui/select/index.js';
  import { cn } from '$lib/utils.js';

  let {
    value = $bindable<DateValue>(),
    placeholder = today(getLocalTimeZone()),
    class: className = '',
    appendYears = 0,
    prependYears = 100,
    ...restProps
  } = $props();

  const currentDate = today(getLocalTimeZone());

  const monthFmt = new DateFormatter('ru-RU', {
    month: 'long',
  });

  const monthOptions = Array.from({ length: 12 }, (_, i) => {
    const month = currentDate.set({ month: i + 1 });
    return {
      value: month.month,
      label: monthFmt.format(month.toDate(getLocalTimeZone())),
    };
  });

  const yearOptions = Array.from({ length: appendYears }, (_, i) => {
    const currentYear = new Date().getFullYear();
    return {
      label: String(currentYear + (appendYears - i)),
      value: currentYear + (appendYears - i),
    };
  })
    // + 1 to include current year
    .concat(
      Array.from({ length: prependYears + 1 }, (_, i) => {
        const currentYear = new Date().getFullYear();
        return {
          label: String(currentYear - i),
          value: currentYear - i,
        };
      }),
    );

  const defaultYear = $derived(
    placeholder
      ? { value: placeholder.year, label: String(placeholder.year) }
      : undefined,
  );

  const defaultMonth = $derived(
    placeholder
      ? {
          value: placeholder.month,
          label: monthFmt.format(placeholder.toDate(getLocalTimeZone())),
        }
      : undefined,
  );

  const monthLabel = $derived(
    monthOptions.find((m) => m.value === defaultMonth?.value)?.label ??
      'Выберите месяц',
  );
  function capitalize(str: string) {
    return str[0]?.toUpperCase() + str.slice(1);
  }
</script>

<!--
Discriminated Unions + Destructing (required for bindable) do not
get along, so we shut typescript up by casting `value` to `never`.
-->
<CalendarPrimitive.Root
  {...restProps}
  type="single"
  weekdayFormat="short"
  locale="ru"
  class={cn('bg-background w-[325px] rounded-md p-3', className)}
  bind:value
  bind:placeholder
>
  {#snippet children({ months, weekdays })}
    <Calendar.Header class="flex w-full items-center justify-between gap-2">
      <Select.Root
        type="single"
        value={`${defaultMonth?.value}`}
        onValueChange={(v) => {
          if (!placeholder) return;
          if (v === `${placeholder.month}`) return;
          placeholder = placeholder.set({ month: Number.parseInt(v) });
        }}
      >
        <Select.Trigger class="w-[60%] border-0">
          {capitalize(monthLabel)}
        </Select.Trigger>
        <Select.Content
          class="border-border max-h-[200px] overflow-y-auto border shadow-xl"
        >
          {#each monthOptions as { value, label } (value)}
            <Select.Item value={`${value}`} label={capitalize(label)} />
          {/each}
        </Select.Content>
      </Select.Root>
      <Select.Root
        type="single"
        value={`${defaultYear?.value}`}
        onValueChange={(v) => {
          if (!v || !placeholder) return;
          if (v === `${placeholder?.year}`) return;
          placeholder = placeholder.set({ year: Number.parseInt(v) });
        }}
      >
        <Select.Trigger class="w-[40%] border-0">
          {defaultYear?.label ?? 'Select year'}
        </Select.Trigger>
        <Select.Content
          class="border-border max-h-[200px] overflow-y-auto border shadow-xl"
        >
          {#each yearOptions as { value, label } (value)}
            <Select.Item value={`${value}`} {label} />
          {/each}
        </Select.Content>
      </Select.Root>
      <Calendar.PrevButton />
      <Calendar.NextButton />
    </Calendar.Header>
    <Calendar.Months>
      {#each months as month (month)}
        <Calendar.Grid>
          <Calendar.GridHead>
            <Calendar.GridRow class="flex">
              {#each weekdays as weekday (weekday)}
                <Calendar.HeadCell>
                  {weekday.slice(0, 2)}
                </Calendar.HeadCell>
              {/each}
            </Calendar.GridRow>
          </Calendar.GridHead>
          <Calendar.GridBody>
            {#each month.weeks as weekDates (weekDates)}
              <Calendar.GridRow class="mt-2 w-full">
                {#each weekDates as date (date)}
                  <Calendar.Cell class="select-none" {date} month={month.value}>
                    <Calendar.Day />
                  </Calendar.Cell>
                {/each}
              </Calendar.GridRow>
            {/each}
          </Calendar.GridBody>
        </Calendar.Grid>
      {/each}
    </Calendar.Months>
  {/snippet}
</CalendarPrimitive.Root>
