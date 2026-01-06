<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form';
  import * as Select from '$lib/components/ui/select/index.js';
  import { scheme } from '$lib/components/editor/activity/scheme';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Activity, Text, Widget } from '$lib/widgetTypes/widgetTypes';
  import client from '$lib/api/client';
  import { Input } from '$lib/components/ui/input';
  import { Separator } from '$lib/components/ui/separator/index.js';

  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: Activity;
    onClose: CallableFunction;
  } = $props();

  const INITIAL_EDUCATION_OBJ = {
    type: 'university',
    educationStage: 'bachelor',
    course: 2,
  };

  const INITIAL_WORK_OBJ = {
    company: '',
    job_title: '',
    experience: null,
  };
  let isLoading = $state(false);
  const form = superForm(defaults(zod4(scheme)), {
    SPA: true,
    validators: zod4(scheme),
    onSubmit: async () => {
      isLoading = true;
      const widget: Activity = {
        type: 'activity',
        status: {
          isStudent: $formData.isStudent,
          isWorker: $formData.isWorker,
        },
        education: $formData.education,
        work: $formData.work,
      };
      console.log(widget);
      if (widgetId != undefined) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
      isLoading = false;
      onClose();
    },
  });
  const { form: formData, enhance, validateForm, reset } = form;

  $effect(() => {
    if (widgetData) {
      $formData.isStudent = widgetData.status.isStudent;
      $formData.isWorker = widgetData.status.isWorker;
      $formData.education = widgetData.education;
      $formData.work = widgetData.work;
      // $formData.text = widgetData.text;
    } else {
      reset();
    }
  });

  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  $effect(() => {
    if (
      $formData.education?.type === 'school' &&
      !('class' in $formData.education)
    ) {
      $formData.education = {
        type: 'school',
        class: 10,
      };
    } else if (
      $formData.education?.type === 'university' &&
      !('course' in $formData.education)
    ) {
      $formData.education = INITIAL_EDUCATION_OBJ;
    }
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form
        method="POST"
        use:enhance
        class="text-text-color overflow-y-auto h-[calc(100vh-100px)]"
      >
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Деятельность"
        </p>
        <div class="mb-4">
          <p class="text-xl font-semibold">Ваш статус</p>
          <p class="">Выберите один или несколько вариантов вашей занятости</p>
        </div>
        <!--        <p>Hello world</p>-->
        <div class="flex gap-x-2 mb-3.5">
          <button
            type="button"
            class="flex flex-col flex-1/2 items-center justify-center p-6 border border-inactive rounded-[8px] {$formData.isStudent
              ? 'bg-accent/25 !border-accent text-accent'
              : ''}"
            onclick={() => {
              $formData.isStudent = !$formData.isStudent;
              if ($formData.isStudent) {
                $formData.education = INITIAL_EDUCATION_OBJ;
              } else $formData.education = undefined;
            }}
          >
            <span class="scale-175 mb-1.5">🎓</span>
            <span>Учусь</span>
          </button>
          <button
            type="button"
            class="flex flex-col flex-1/2 items-center justify-center p-6 border border-inactive rounded-[8px] {$formData.isWorker
              ? 'bg-accent/25 !border-accent text-accent'
              : ''}"
            onclick={() => {
              $formData.isWorker = !$formData.isWorker;
              if ($formData.isWorker) {
                $formData.work = INITIAL_WORK_OBJ;
              } else $formData.work = undefined;
            }}
          >
            <span class="scale-175 mb-1.5">💼</span>
            <span>Работаю</span>
          </button>
        </div>
        {#if $formData.isStudent}
          <div class="flex w-full items-center overflow-hidden">
            <span class="ml-1.5 text-nowrap text-accent-foreground mr-2.5"
              >🎓 ОБУЧЕНИЕ
            </span>
            <!--            <Separator orientation="horizontal" class="bg-inactive w-full" />-->
          </div>
          <span class="text-[14px] font-semibold">ТИП ЗАВЕДЕНИЯ</span>
          <Select.Root type="single" bind:value={$formData.education.type}>
            <Select.Trigger class="w-full">
              {$formData.education.type === 'school' ? 'Школа' : 'Университет'}
            </Select.Trigger>
            <Select.Content>
              <Select.Group>
                {#each [{ value: 'school', label: 'Школа' }, { value: 'university', label: 'Университет' }] as type (type)}
                  <Select.Item value={type.value} label={type.label}>
                    {type.label}
                  </Select.Item>
                {/each}
              </Select.Group>
            </Select.Content>
          </Select.Root>
          {#if $formData.education.type === 'university'}
            <div class="flex gap-x-2">
              <div class="flex-1/2">
                <span class="text-[13px] font-semibold">СТУПЕНЬ</span>
                <Select.Root
                  type="single"
                  bind:value={$formData.education.educationStage}
                >
                  <Select.Trigger class="w-full">
                    {$formData.education.educationStage === 'bachelor'
                      ? 'Бакалавриат'
                      : $formData.education.educationStage === 'master'
                        ? 'Магистратура'
                        : 'Аспирантура'}
                  </Select.Trigger>
                  <Select.Content>
                    <Select.Group>
                      {#each [{ value: 'bachelor', label: 'Бакалавриат' }, { value: 'master', label: 'Магистратура' }, { value: 'postgraduate', label: 'Аспирантура' }] as type (type)}
                        <Select.Item value={type.value} label={type.label}>
                          {type.label}
                        </Select.Item>
                      {/each}
                    </Select.Group>
                  </Select.Content>
                </Select.Root>
              </div>
              <div class="flex-1/2">
                <Form.Field {form} name="education.course">
                  <Form.Control>
                    {#snippet children({ props })}
                      <Form.Label class="text-[14px] font-semibold"
                        >КУРС</Form.Label
                      >
                      <Input
                        {...props}
                        bind:value={$formData.education.course}
                        type="number"
                      />
                    {/snippet}
                  </Form.Control>
                  <Form.Description />
                  <Form.FieldErrors />
                </Form.Field>
              </div>
            </div>
          {:else if $formData.education.type === 'school'}
            <Form.Field {form} name="education.class">
              <Form.Control>
                {#snippet children({ props })}
                  <Form.Label class="text-[14px] font-semibold"
                    >КЛАСС</Form.Label
                  >
                  <Input
                    {...props}
                    bind:value={$formData.education.class}
                    type="number"
                  />
                {/snippet}
              </Form.Control>
              <Form.Description />
              <Form.FieldErrors />
            </Form.Field>
          {/if}
        {/if}
        {#if $formData.isWorker}
          <div class="flex w-full items-center overflow-hidden">
            <span class="ml-1.5 text-nowrap text-accent-foreground mr-2.5"
              >💼 РАБОТА
            </span>
            <!--            <Separator orientation="horizontal" class="bg-inactive w-full" />-->
          </div>
          <Form.Field {form} name="work.company">
            <Form.Control>
              {#snippet children({ props })}
                <Form.Label class="">МЕСТО РАБОТЫ</Form.Label>
                <Input
                  {...props}
                  bind:value={$formData.work.company}
                  placeholder="Назвиние компании"
                />
              {/snippet}
            </Form.Control>
            <Form.Description />
            <Form.FieldErrors />
          </Form.Field>
          <div class="flex gap-x-2">
            <Form.Field {form} name="work.job_title">
              <Form.Control>
                {#snippet children({ props })}
                  <Form.Label class="">ДОЛЖНОСТЬ</Form.Label>
                  <Input
                    {...props}
                    bind:value={$formData.work.job_title}
                    placeholder="Ваша роль"
                  />
                {/snippet}
              </Form.Control>
              <Form.Description />
              <Form.FieldErrors />
            </Form.Field>
            <Form.Field {form} name="work.experience">
              <Form.Control>
                {#snippet children({ props })}
                  <Form.Label class="">СТАЖ</Form.Label>
                  <Input
                    {...props}
                    bind:value={$formData.work.experience}
                    type="number"
                  />
                {/snippet}
              </Form.Control>
              <Form.Description />
              <Form.FieldErrors />
            </Form.Field>
          </div>
        {/if}
        {#if widgetId !== undefined}
          <DeleteButton {widgetId} />
        {/if}

        <SaveButton
          {isLoading}
          onClick={() => {
            form.submit();
          }}
          {isButtonActive}
        />
        {#if import.meta.env.DEV}
          <SuperDebug data={$formData} />
        {/if}
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>
