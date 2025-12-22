<script lang="ts">
  import Aries from '../signs/Aries.svelte';
  import Taurus from '../signs/Taurus.svelte';
  import Gemini from '../signs/Gemini.svelte';
  import Cancer from '../signs/Cancer.svelte';
  import Leo from '../signs/Leo.svelte';
  import Virgo from '../signs/Virgo.svelte';
  import Libra from '../signs/Libra.svelte';
  import Scorpio from '../signs/Leo.svelte';
  import Sagittarius from '../signs/Sagittarius.svelte';
  import Capricorn from '../signs/Capricorn.svelte';
  import Aquarius from '../signs/Aquarius.svelte';
  import Pisces from '../signs/Pisces.svelte';

  // let { birthDate } = $props();
  //
  // console.log(birthDate);
  let birthDate = new Date();
  // let birthDate = pb.authStore.record?.birth_date;

  type ZodiacSign =
    | 'aquarius'
    | 'pisces'
    | 'aries'
    | 'taurus'
    | 'gemini'
    | 'cancer'
    | 'leo'
    | 'virgo'
    | 'libra'
    | 'scorpio'
    | 'sagittarius'
    | 'capricorn';

  const signs = {
    2: {
      day: 19,
      signs: ['aquarius', 'pisces'],
    },
    3: {
      day: 20,
      signs: ['pisces', 'aries'],
    },
    4: {
      day: 20,
      signs: ['aries', 'taurus'],
    },
    5: {
      day: 20,
      signs: ['taurus', 'gemini'],
    },
    6: {
      day: 20,
      signs: ['gemini', 'cancer'],
    },
    7: {
      day: 22,
      signs: ['cancer', 'leo'],
    },
    8: {
      day: 22,
      signs: ['leo', 'virgo'],
    },
    9: {
      day: 22,
      signs: ['vigro', 'libra'],
    },
    10: {
      day: 22,
      signs: ['libra', 'scorpio'],
    },
    11: {
      day: 22,
      sings: ['scorpio', 'sagittarius'],
    },
    12: {
      day: 21,
      signs: ['sagittarius', 'capricorn'],
    },
    1: {
      day: 19,
      signs: ['capricorn', 'aquarius'],
    },
  };

  function defineSign(birthDate: string): ZodiacSign {
    const date = new Date(birthDate);
    console.log(typeof birthDate);
    console.log(date);
    const month = date.getMonth() + 1;
    const day = date.getDate();
    console.log(month);
    console.log(day);
    if (day <= signs[month].day) return signs[month].signs[0];
    else return signs[month].signs[1];
  }

  let zodiac = defineSign(birthDate);
</script>

<div class="">
  {#if zodiac == 'aquarius'}
    <Aquarius />
  {:else if zodiac == 'pisces'}
    <Pisces />
  {:else if zodiac == 'aries'}
    <Aries />
  {:else if zodiac == 'taurus'}
    <div class="flex items-center font-semibold text-[20px] break-words">
      <Taurus />
    </div>
  {:else if zodiac == 'gemini'}
    <Gemini />
  {:else if zodiac == 'cancer'}
    <Cancer />
  {:else if zodiac == 'leo'}
    <Leo />
  {:else if zodiac == 'virgo'}
    <Virgo />
  {:else if zodiac == 'libra'}
    <Libra />
  {:else if zodiac == 'scorpio'}
    <Scorpio />
  {:else if zodiac == 'sagittarius'}
    <Sagittarius />
  {:else if zodiac == 'capricorn'}
    <Capricorn class="" />
  {/if}
</div>
