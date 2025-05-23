import { getSocialMediaData } from '$lib/index';

// eslint-disable-next-line no-undef
onRecordEnrich((e) => {
  e.record.withCustomData(true); // for security custom props require to be enabled explicitly
  if (e.record.data.type == 'social_media')
    e.record.set('additionalData', {
      socialMediaData: getSocialMediaData(e.record.data.link),
    });

  e.next();
}, 'widgets');
