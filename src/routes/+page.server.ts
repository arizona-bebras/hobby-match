export const actions = {
    default: async (event: any) => {
        const data = await event.request.formData();
        const name = data.get('name');
        if (typeof name === 'string') {
            event.locals.user = name;
        }
        console.log(name);

    }
};


export function load(event: any) {
    return {
        name: event.locals.user,
    };
}
