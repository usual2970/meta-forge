/**
 * This middleware is used to dynamically update the Layouts system.
 *
 * As soon as the route changes, it tries to pull the layout that we want to display from the laptop. Then it loads the layout component, and assigns the loaded component to the meta in the layout Component variable. And layoutComponent is used in the basic layout AppLayout.vue, there is a dynamic update of the layout component
 *
 * If the layout we want to display is not found, loads the default layout App Layout Default.vue
 * */



export async function loadLayoutMiddleware(route) {
    try {
        let layout = route.meta.layout
        let layoutComponent = await import(`../../layouts/${layout}.vue`)
        route.meta.layoutComponent = layoutComponent.default
    } catch (e) {
        console.error('Error occurred in processing of layouts: ', e)
        console.log('Mounted default layout AppLayoutDefault')
        let layout = 'Blank'
        let layoutComponent = await import(`../../layouts/${layout}.vue`)
        route.meta.layoutComponent = layoutComponent.default
    }
}