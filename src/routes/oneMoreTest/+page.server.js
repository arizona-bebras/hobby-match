export const actions = {
    default: async (event) => {
        const data = await event.request.formData();
        const name = data.get('name');
        if(typeof name === 'string') {
            event.locals.user = name;
        }
    }
}



export function load(event) {
    return {
        name: event.locals.user,
    };
}
