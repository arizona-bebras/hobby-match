<script>
  import { pb } from '$lib/index';
  import { goto } from '$app/navigation';
  if (pb.authStore.isValid) {
    console.log(pb.authStore.model)
    if (pb.authStore.model?.mini_app_name == "" ||
        pb.authStore.model?.gender == "" ||
        pb.authStore.model?.birth_date == "" ||
        pb.authStore.model?.location == "" ||
        pb.authStore.model?.user_info == "" ||
        pb.authStore.model?.photo == "" ||
        pb.authStore.model?.interests == null) 
        { goto('/registration'); }
    else { goto('/Profile')}
  }
  else {
    goto('AuthError')
  }
</script>

<h1>Auth...</h1>
